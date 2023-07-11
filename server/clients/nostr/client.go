package nostr

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
	"unsafe"

	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip04"
	"github.com/nbd-wtf/go-nostr/nip19"
	"github.com/nbd-wtf/go-nostr/nip42"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type (
	NostrEvent = nostr.Event

	RelayEvent struct {
		Relay string     `json:"relay"`
		Event NostrEvent `json:"event"`
	}

	// Client for nostr protocol
	Client struct {
		// Reference to the server we are using
		server *Server
		// Secret key
		sk string
		// Public key
		pk string
	}

	// Subscription for events on a relay
	Subscription struct {
		id     string
		buffer *eventBuffer
		subs   []*nostr.Subscription
	}

	// Metadata used when setting metadata, see [nip01](https://github.com/nostr-protocol/nips/blob/master/01.md)
	Metadata struct {
		Name    string `json:"name"`
		About   string `json:"about"`
		Picture string `json:"picture"`
	}
)

const (
	// size of a subscription id
	SUB_ID_LENGTH = 10

	DEFAULT_LIMIT = 100

	kindSetMetadata     = 0
	kindTextNote        = 1
	kindRecommendServer = 2
	kindDirectMessage   = 4
)

var (
	// The default duration to try and connect to a relay
	relayConnectTimeout = time.Second * 5
	// The default duration to try and authenticate to a relay
	relayAuthTimeout = time.Second * 5

	// ErrRelayAuthFailed indicates the authentication on a relay completed, but failed
	ErrRelayAuthFailed = errors.New("failed to authenticate to the relay")
	// ErrRelayAuthTimeout indicates the authentication on a relay did not complete in time
	ErrRelayAuthTimeout = errors.New("timeout authenticating to the relay")
	// ErrFailedToPublishEvent indicates the event could not be published to the relay
	ErrFailedToPublishEvent = errors.New("failed to publish event to relay")
	/// ErrNoRelayConnected inidcates that we try to perform an action on a realay, but we aren't connected to any.
	ErrNoRelayConnected = errors.New("no relay connected currently")
)

func GenerateKeyPair() string {
	return nostr.GeneratePrivateKey()
}

// Id of the client, this is the enocded public key in NIP19 format
func (c *Client) Id() string {
	id, err := nip19.EncodePublicKey(c.pk)
	if err != nil {
		panic(fmt.Sprintf("can't encode public key, although this was previously validated. this should not happen (%s)", err))
	}

	return id
}

func (c *Client) PublicKey() string {
	return c.pk
}

func (c *Client) ConnectRelay(ctx context.Context, relayURL string) error {
	// ctxConnect, cancelFuncConnect := context.WithTimeout(ctx, relayConnectTimeout)
	// defer cancelFuncConnect()
	ctxConnect := context.Background()

	relay, err := nostr.RelayConnect(ctxConnect, relayURL)
	if err != nil {
		return errors.Wrap(err, "failed to connect to the provided relay")
	}

	// Add relay to the list of managed relays
	c.server.manageRelay(c.Id(), relay)

	return nil
}

// ConnectAuthRelay connect and authenticates to a NIP42 authenticated relay
func (c *Client) ConnectAuthRelay(ctx context.Context, relayURL string) error {
	ctxConnect, cancelFuncConnect := context.WithTimeout(ctx, relayConnectTimeout)
	defer cancelFuncConnect()

	relay, err := nostr.RelayConnect(ctxConnect, relayURL)
	if err != nil {
		return errors.Wrap(err, "failed to connect to the provided relay")
	}

	// Wait for the challenge send by the relay
	challenge := <-relay.Challenges

	event := nip42.CreateUnsignedAuthEvent(challenge, c.pk, relayURL)
	event.Sign(c.sk)

	ctxAuth, cancelFuncAuth := context.WithTimeout(ctx, relayAuthTimeout)
	defer cancelFuncAuth()

	auth_status, err := relay.Auth(ctxAuth, event)
	if err != nil {
		return errors.Wrap(err, "could not authenticate to relay")
	}

	if auth_status != nostr.PublishStatusSucceeded {
		return ErrRelayAuthFailed
	}

	// Add relay to the list of managed relays
	c.server.manageRelay(c.Id(), relay)

	return nil
}

// Add function to publish events to a set of relays, and returns the published event ID if successful
func (c *Client) publishEventToRelays(ctx context.Context, kind int, tags [][]string, content string) (string, error) {
	c.server.mutex.RLock()
	defer c.server.mutex.RUnlock()

	relays := c.server.connectedRelays[c.Id()]
	if len(relays) == 0 {
		return "", errors.New("no relays connected")
	}

	parsedTags := make(nostr.Tags, 0, len(tags))
	for _, rawTag := range tags {
		parsedTags = append(parsedTags, nostr.Tag(rawTag))
	}

	ev := nostr.Event{
		PubKey:    c.pk,
		CreatedAt: time.Now(),
		Kind:      kind,
		Tags:      parsedTags,
		Content:   content,
	}

	// calling Sign sets the event ID field and the event Sig field
	ev.Sign(c.sk)

	log.Debug().Str("component", "nostr").Msgf("Publish event to connected relays")
	for _, relay := range c.server.connectedRelays[c.Id()] {
		log.Debug().Str("component", "nostr").Msgf("publising event to relay: %+v", ev)
		status, err := relay.Publish(ctx, ev)
		if err != nil {
			return "", errors.Wrap(err, fmt.Sprintf("could not publish event to relay %s", relay.URL))
		}

		log.Debug().Str("component", "nostr").Msgf("published event to relay: %+v with status:%s", ev, status)

		if status == nostr.PublishStatusFailed {
			return "", ErrFailedToPublishEvent
		}
	}

	return ev.ID, nil
}

// PublishMetadata to connected relays. If metadata was published previously, the old metadata should be overwritten conforming relays
func (c *Client) PublishMetadata(ctx context.Context, tags []string, content Metadata) error {
	marshalledContent, err := json.Marshal(content)
	if err != nil {
		return errors.Wrap(err, "could not encode metadata")
	}

	if _, err := c.publishEventToRelays(ctx, kindSetMetadata, [][]string{tags}, string(marshalledContent)); err != nil {
		return err
	}

	return nil
}

// PublishTextNote to connected relays
func (c *Client) PublishTextNote(ctx context.Context, tags []string, content string) error {
	if _, err := c.publishEventToRelays(ctx, kindTextNote, [][]string{tags}, content); err != nil {
		return err
	}

	return nil
}

// PublishRecommendServer to connected relays. The content is supposed to be the URL of the relay being recommended
func (c *Client) PublishRecommendServer(ctx context.Context, tags []string, content string) error {
	if _, err := c.publishEventToRelays(ctx, kindRecommendServer, [][]string{tags}, content); err != nil {
		return err
	}

	return nil
}

// / PublishDirectMessage publishes a direct message for a given peer identified by the given pubkey on the connected relays
func (c *Client) PublishDirectMessage(ctx context.Context, receiver string, tags []string, content string) error {
	log.Debug().Str("Receiver", receiver).Msg("Sending direct message")
	ss, err := nip04.ComputeSharedSecret(receiver, c.sk)
	if err != nil {
		return errors.Wrap(err, "could not compute shared secret for receiver")
	}
	msg, err := nip04.Encrypt(content, ss)
	if err != nil {
		return errors.Wrap(err, "could not encrypt message")
	}

	if _, err := c.publishEventToRelays(ctx, kindDirectMessage, [][]string{{"p", receiver}, tags}, msg); err != nil {
		return err
	}

	return nil
}

// SubscribeTextNotes to textnote events on a relay
func (c *Client) SubscribeTextNotes() (string, error) {
	var filters nostr.Filters
	if _, _, err := nip19.Decode(c.Id()); err == nil {
		filters = []nostr.Filter{{
			Kinds: []int{kindTextNote},
			Limit: DEFAULT_LIMIT,
		}}
	} else {
		return "", errors.New("could not create client filters")
	}

	return c.subscribeWithFiler(filters)

}

// SubscribeMessages subscribes to direct messages (Kind 4) on all relays and decrypts them if they are addressed to the client
func (c *Client) SubscribeMessages() (string, error) {
	var filters nostr.Filters
	if _, v, err := nip19.Decode(c.Id()); err == nil {
		t := make(map[string][]string)
		t["p"] = []string{v.(string)}
		filters = []nostr.Filter{{
			Kinds: []int{kindDirectMessage},
			Limit: DEFAULT_LIMIT,
			Tags:  t,
		}}
	} else {
		return "", errors.New("could not create client filters")
	}

	return c.subscribeWithFiler(filters)
}

// SubscribeStallCreation subscribes to stall creation events (Kind 30017) on all relays
func (c *Client) SubscribeStallCreation(tag string) (string, error) {
	var filters nostr.Filters
	if _, _, err := nip19.Decode(c.Id()); err == nil {
		filters = []nostr.Filter{{
			Kinds: []int{kindSetStall},
			Limit: DEFAULT_LIMIT,
		}}
		if tag != "" {
			t := make(map[string][]string)
			t["t"] = []string{tag}
			filters[0].Tags = t
		}
	} else {
		return "", errors.New("could not create client filters")
	}

	return c.subscribeWithFiler(filters)
}

// Subscribe ProductCreation subscribes to product creation events (Kind 30018) on all relays
func (c *Client) SubscribeProductCreation(tag string) (string, error) {
	var filters nostr.Filters
	if _, _, err := nip19.Decode(c.Id()); err == nil {
		filters = []nostr.Filter{{
			Kinds: []int{kindSetProduct},
			Limit: DEFAULT_LIMIT,
		}}
		if tag != "" {
			t := make(map[string][]string)
			t["t"] = []string{tag}
			filters[0].Tags = t
		}
	} else {
		return "", errors.New("could not create client filters")
	}

	return c.subscribeWithFiler(filters)
}

func (c *Client) fetchEventsWithFilter(filters nostr.Filters) ([]RelayEvent, error) {
	relays := c.server.clientRelays(c.Id())
	if len(relays) == 0 {
		return nil, ErrNoRelayConnected
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mu := sync.Mutex{}
	events := []RelayEvent{}
	wg := sync.WaitGroup{}
	wg.Add(len(relays))

	for _, relay := range relays {
		log.Debug().Msgf("NOSTR: Connected to relay %s", relay.URL)
		sub, err := relay.Subscribe(ctx, filters)
		if err != nil {
			log.Error().Msgf("error subscribing to relay: %s", err.Error())
			return nil, errors.Wrapf(err, "could not subscribe to relay %s", relay.URL)
		}

		go func() {
			<-sub.EndOfStoredEvents
			cancel()
			log.Debug().Msg("End of stored events")
		}()

		go func(relay *nostr.Relay) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case ev := <-sub.Events:
					{
						log.Debug().Msgf("NOSTR: Received event from relay, kind: %d", ev.Kind)

						// Decrypt direct messages
						if ev.Kind == kindDirectMessage {
							log.Debug().Msgf("NOSTR: Decrypting message from relay")

							ss, err := nip04.ComputeSharedSecret(ev.PubKey, c.sk)
							if err != nil {
								log.Error().Msgf("could not compute shared secret for receiver %s", err.Error())
								continue
							}
							msg, err := nip04.Decrypt(ev.Content, ss)
							if err != nil {
								log.Error().Msgf("could not decrypt message %s", err.Error())
								continue
							}

							// Set decrypted content
							ev.Content = msg
						}

						mu.Lock()
						events = append(events, RelayEvent{
							Relay: relay.URL,
							Event: *ev,
						})
						mu.Unlock()
					}
				}
			}
		}(relay)
	}

	wg.Wait()

	return events, nil
}

func (c *Client) subscribeWithFiler(filters nostr.Filters) (string, error) {
	relays := c.server.clientRelays(c.Id())
	if len(relays) == 0 {
		return "", ErrNoRelayConnected
	}

	subs := []*nostr.Subscription{}

	ctx := context.Background()
	buf := newEventBuffer()

	for _, relay := range relays {
		log.Debug().Msgf("NOSTR: Connected to relay %s", relay.URL)
		sub, err := relay.Subscribe(ctx, filters)
		if err != nil {
			log.Error().Msgf("error subscribing to relay: %s", err.Error())
			return "", errors.Wrapf(err, "could not subscribe to relay %s", relay.URL)
		}

		subs = append(subs, sub)

		go func() {
			<-sub.EndOfStoredEvents
			log.Debug().Msg("End of stored events")
		}()

		go func() {
			for ev := range sub.Events {
				log.Debug().Msgf("NOSTR: Received event from relay, kind: %d", ev.Kind)

				// Decrypt direct messages
				if ev.Kind == kindDirectMessage {
					log.Debug().Msgf("NOSTR: Decrypting message from relay")

					ss, err := nip04.ComputeSharedSecret(ev.PubKey, c.sk)
					if err != nil {
						log.Error().Msgf("could not compute shared secret for receiver %s", err.Error())
						continue
					}
					msg, err := nip04.Decrypt(ev.Content, ss)
					if err != nil {
						log.Error().Msgf("could not decrypt message %s", err.Error())
						continue
					}

					// Set decrypted content
					ev.Content = msg
				}

				buf.push(ev)
			}
		}()
	}

	sub := &Subscription{
		id:     randString(SUB_ID_LENGTH),
		buffer: buf,
		subs:   subs,
	}

	c.server.manageSubscription(c.Id(), sub)

	return sub.id, nil
}

// TODO: Remove once subsciptions are more porper
func (c *Client) SubscribeDirectMessagesDirect(swapTag string) (<-chan NostrEvent, error) {
	var filters nostr.Filters
	if _, v, err := nip19.Decode(c.Id()); err == nil {
		t := make(map[string][]string)
		t["p"] = []string{v.(string)}
		t["s"] = []string{swapTag}
		filters = []nostr.Filter{{
			Kinds: []int{kindDirectMessage},
			Limit: 0,
			Tags:  t,
		}}
	} else {
		return nil, errors.New("could not create client filters")
	}

	relays := c.server.clientRelays(c.Id())
	if len(relays) == 0 {
		log.Error().Msg("No relays connected to subscribe for direct messages")
		return nil, ErrNoRelayConnected
	}

	ctx := context.Background()

	log.Debug().Msgf("Connecting to relays to subscribe to direct messages, with filters %+v", filters)
	ch := make(chan NostrEvent)
	for _, relay := range relays {
		log.Debug().Msgf("NOSTR: Connected to relay %s", relay.URL)
		sub, err := relay.Subscribe(ctx, filters)
		if err != nil {
			log.Error().Msgf("error subscribing to relay: %s", err.Error())
			return nil, errors.Wrapf(err, "could not subscribe to relay %s", relay.URL)
		}

		go func() {
			<-sub.EndOfStoredEvents
			log.Debug().Msg("End of stored events")
		}()

		go func() {
			for ev := range sub.Events {
				log.Debug().Msgf("NOSTR: Received event from relay, kind: %d", ev.Kind)

				// Decrypt direct messages
				log.Debug().Msgf("NOSTR: Decrypting message from relay")

				ss, err := nip04.ComputeSharedSecret(ev.PubKey, c.sk)
				if err != nil {
					log.Error().Msgf("could not compute shared secret for receiver %s", err.Error())
					continue
				}
				msg, err := nip04.Decrypt(ev.Content, ss)
				if err != nil {
					log.Error().Msgf("could not decrypt message %s", err.Error())
					continue
				}

				// Set decrypted content
				ev.Content = msg

				ch <- *ev

			}
		}()
	}

	return ch, nil
}

// Get all historic events on active subscriptions for the client.
// Note that only a limited amount of events are kept. If the actual client waits
// too long to call this, events might be dropped.
// returned events are sorted from oldes to newest
func (c *Client) GetEvents() []NostrEvent {
	subs := c.server.subscriptions(c.Id())
	var events []NostrEvent
	for _, sub := range subs {
		events = append(events, sub.buffer.take()...)
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].CreatedAt.Unix() < events[j].CreatedAt.Unix()
	})

	return events
}

// GetSubscriptionEvents for a subscription with the given ID. Events are removed from the subscription
func (c *Client) GetSubscriptionEvents(id string) []NostrEvent {
	subs := c.server.subscriptions(c.Id())
	for _, sub := range subs {
		if sub.id == id {
			return sub.buffer.take()
		}
	}
	return nil
}

// GetSubscriptionEventsWithCount returns a number of events for a subscription with the given ID. Returned events are removed from the subscription
func (c *Client) GetSubscriptionEventsWithCount(id string, count uint32) []NostrEvent {
	subs := c.server.subscriptions(c.Id())
	for _, sub := range subs {
		if sub.id == id {
			return sub.buffer.consume(count)
		}
	}
	return nil
}

// Get the ID's of all active subscriptions
func (c *Client) SubscriptionIds() []string {
	subs := c.server.subscriptions(c.Id())
	var ids []string
	for _, sub := range subs {
		ids = append(ids, sub.id)
	}
	return ids
}

// CloseSubscription managed by the server for this client, based on its ID.
func (c *Client) CloseSubscription(id string) {
	log.Debug().Msg("NOSTR: calling close subscription")
	sub := c.server.removeSubscription(c.Id(), id)
	if sub != nil {
		sub.Close()
	}
}

// Close an open subscription
func (s *Subscription) Close() {
	for _, sub := range s.subs {
		sub.Unsub()
	}
}

// go random string, source: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randString(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// end random string code copy
