package nostr

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
	"github.com/nbd-wtf/go-nostr/nip42"
	"github.com/pkg/errors"
)

type (
	// TODO: break out into individual files

	// Server is a persistent client keeping open connections to relays.
	Server struct {
		connectedRelays     map[string][]*nostr.Relay
		clientSubscriptions map[string]*Subscription

		mutex sync.RWMutex
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
		Events <-chan nostr.Event
	}
)

var (
	// The default duration to try and connect to a relay
	relayConnectTimeout = time.Second * 5
	// The default duration to try and authenticate to a relay
	relayAuthTimeout = time.Second * 5

	// ErrRelayAuthFailed indicates the authentication on a relay completed, but failed
	ErrRelayAuthFailed = errors.New("Failed to authenticate to the relay")
	// ErrRelayAuthTimeout indicates the authentication on a relay did not complete in time
	ErrRelayAuthTimeout = errors.New("Timeout authenticating to the relay")
	// ErrFailedToPublishEvent indicates the event could not be published to the relay
	ErrFailedToPublishEvent = errors.New("Failed to publish event to relay")
)

// NewServer managing relay connections and subscriptions for possibly different peers
func NewServer() *Server {
	return &Server{
		connectedRelays:     make(map[string][]*nostr.Relay),
		clientSubscriptions: make(map[string]*Subscription),
	}
}

// NewClient for a server, authenticated by the private key of the client. Private key is passed as hex bytes
func (s *Server) NewClient(sk string) (*Client, error) {
	pk, err := nostr.GetPublicKey(sk)
	if err != nil {
		return nil, errors.Wrap(err, "could not get public key from provided private key")
	}

	return &Client{
		server: s,
		sk:     sk,
		pk:     pk,
	}, nil
}

func (s *Server) manageRelay(id string, relay *nostr.Relay) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, r := range s.connectedRelays[id] {
		if r == relay {
			return
		}
	}
	s.connectedRelays[id] = append(s.connectedRelays[id], relay)
}

// Id of the client, this is the enocded public key in NIP19 format
func (c *Client) Id() string {
	id, err := nip19.EncodePublicKey(c.pk)
	if err != nil {
		panic(fmt.Sprintf("Can't encode public key, although this was previously validated. This should not happen (%s)", err))
	}

	return id
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

// Add function to publish events to a set of relays
func (c *Client) PublishEventToRelays(ctx context.Context, tags []string, content string) error {
	c.server.mutex.RLock()
	defer c.server.mutex.RUnlock()

	relays := c.server.connectedRelays[c.Id()]
	if len(relays) == 0 {
		return errors.New("No relays connected")
	}

	ev := nostr.Event{
		PubKey:    c.pk,
		CreatedAt: time.Now(),
		Kind:      1,
		Tags:      make(nostr.Tags, len(tags)),
		Content:   content,
	}

	// calling Sign sets the event ID field and the event Sig field
	ev.Sign(c.sk)

	for _, relay := range c.server.connectedRelays[c.Id()] {
		status, err := relay.Publish(ctx, ev)
		if err != nil {
			return errors.Wrap(err, ErrFailedToPublishEvent.Error())
		}

		if status != nostr.PublishStatusSucceeded {
			return ErrFailedToPublishEvent
		}
	}

	return nil
}
