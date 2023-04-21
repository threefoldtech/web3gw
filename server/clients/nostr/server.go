package nostr

import (
	"sync"

	"github.com/nbd-wtf/go-nostr"
	"github.com/pkg/errors"
)

type (

	// Server is a persistent client keeping open connections to relays.
	Server struct {
		connectedRelays     map[string][]*nostr.Relay
		clientSubscriptions map[string][]*Subscription

		mutex sync.RWMutex
	}
)

// NewServer managing relay connections and subscriptions for possibly different peers
func NewServer() *Server {
	return &Server{
		connectedRelays:     make(map[string][]*nostr.Relay),
		clientSubscriptions: make(map[string][]*Subscription),
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

// Manage an active relay connection for a client
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

// Get the list of all relays managed for the given client. These relays must have been
// added first through the manageRelay method
func (s *Server) clientRelays(id string) []*nostr.Relay {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.connectedRelays[id]
}

// Manage an active subscription for a client
func (s *Server) manageSubscription(id string, sub *Subscription) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, oldSub := range s.clientSubscriptions[id] {
		if oldSub.id == sub.id {
			return
		}
	}

	s.clientSubscriptions[id] = append(s.clientSubscriptions[id], sub)
}

// Get a list of all the subscriptions being managed for a client
func (s *Server) subscriptions(id string) []*Subscription {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.clientSubscriptions[id]
}

// Remove an active subscription for a client by its ID
func (s *Server) removeSubscription(id string, subID string) *Subscription {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i, sub := range s.clientSubscriptions[id] {
		if sub.id == subID {
			s.clientSubscriptions[id] = append(s.clientSubscriptions[id][:i], s.clientSubscriptions[id][i+1:]...)
			return sub
		}
	}

	return nil
}
