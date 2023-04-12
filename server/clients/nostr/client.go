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

	// TODO: check if we already manage such a relay?
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

	// Either we succeeded and all is good, or we did not and should tear down the connection properly
	switch auth_status {
	case nostr.PublishStatusFailed:
		// Best effort cleanup
		relay.Close()
		return ErrRelayAuthFailed
	case nostr.PublishStatusSent:
		// Best effort cleanup
		relay.Close()
		return ErrRelayAuthTimeout
	case nostr.PublishStatusSucceeded:
		c.server.manageRelay(c.Id(), relay)
		return nil
	}

	// Go doesn't understand we checked all cases above (and that's why you have proper enums), we could just return nil
	// here but this panic is a nice sanity check.
	panic("Unreachable")
}
