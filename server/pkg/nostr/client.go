package nostr

import (
	"context"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/threefoldtech/web3_proxy/server/clients/nostr"
	"github.com/threefoldtech/web3_proxy/server/pkg"
)

const (
	// NostrID is the ID for state of a nostr client in the connection state.
	NostrID = "nostr"
)

type (
	// Client exposes nostr related functionality
	Client struct {
		server *nostr.Server
	}
	// state managed by nostr client
	nostrState struct {
		client *nostr.Client
	}
)

// State from a connection. If no state is present, it is initialized
func State(conState jsonrpc.State) *nostrState {
	raw, exists := conState[NostrID]
	if !exists {
		ns := &nostrState{
			client: nil,
		}
		conState[NostrID] = ns
		return ns
	}
	ns, ok := raw.(*nostrState)
	if !ok {
		// This means the invariant is violated, so panic here is ok
		panic("Invalid saved state for nostr")
	}
	return ns
}

// NewClient creates a new client
func NewClient() *Client {
	return &Client{
		server: nostr.NewServer(),
	}
}

// Load a client from a connection state
func (c *Client) Load(ctx context.Context, conState jsonrpc.State, secret string) error {
	cl, err := c.server.NewClient(secret)
	if err != nil {
		return err
	}

	state := State(conState)
	state.client = cl

	return nil
}

// GetPublicKey returns the nostr ID for the client
func (c *Client) GetId(ctx context.Context, conState jsonrpc.State) (string, error) {
	state := State(conState)
	if state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.Id(), nil
}

// GetPublicKey returns the public key of the client in hex
func (c *Client) GetPublicKey(ctx context.Context, conState jsonrpc.State) (string, error) {
	state := State(conState)
	if state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.PublicKey(), nil
}

// ConnectRelay connects to an authenticated relay with a given url
func (c *Client) ConnectAuthRelay(ctx context.Context, conState jsonrpc.State, url string) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ConnectAuthRelay(ctx, url)
}

// ConnectRelay connects to a relay with a given url
func (c *Client) ConnectRelay(ctx context.Context, conState jsonrpc.State, url string) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ConnectRelay(ctx, url)
}

// GenerateKeyPair generates a new keypair
func (c *Client) GenerateKeyPair(ctx context.Context) (string, error) {
	return nostr.GenerateKeyPair(), nil
}

// ConnectToRelay connects to a relay with a given url
func (c *Client) ConnectToRelay(ctx context.Context, conState jsonrpc.State, url string) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ConnectAuthRelay(ctx, url)
}

// TextNote is a text note published on a relay
type TextInput struct {
	Tags    []string `json:"tags"`
	Content string   `json:"content"`
}

// PublishTextNote publishes a text note to all relays
func (c *Client) PublishTextNote(ctx context.Context, conState jsonrpc.State, input TextInput) error {
	state := State(conState)

	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.PublishTextNote(ctx, input.Tags, input.Content)
}

// MetadataInput is metadata published on a relay
type MetadataInput struct {
	Tags     []string       `json:"tags"`
	Metadata nostr.Metadata `json:"metadata"`
}

// PublishMetadata publishes metadata to all relays
func (c *Client) PublishMetadata(ctx context.Context, conState jsonrpc.State, input MetadataInput) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.PublishMetadata(ctx, input.Tags, input.Metadata)
}

type DirectMessageInput struct {
	Receiver string   `json:"receiver"`
	Tags     []string `json:"tags"`
	Content  string   `json:"content"`
}

// PublishDirectMessage publishes a direct message to a receiver
func (c *Client) PublishDirectMessage(ctx context.Context, conState jsonrpc.State, input DirectMessageInput) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.PublishDirectMessage(ctx, input.Receiver, input.Tags, input.Content)
}

// SubscribeRelays subscribes to text notes on all relays
func (c *Client) SubscribeRelays(ctx context.Context, conState jsonrpc.State) (string, error) {
	state := State(conState)
	if state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.SubscribeRelays()
}

// SubscribeDirectMessages subscribes to direct messages on all relays and decrypts them
func (c *Client) SubscribeDirectMessages(ctx context.Context, conState jsonrpc.State) (string, error) {
	state := State(conState)
	if state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.SubscribeMessages()
}

// CloseSubscription closes a subscription by id
func (c *Client) CloseSubscription(ctx context.Context, conState jsonrpc.State, id string) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	state.client.CloseSubscription(id)

	return nil
}

// GetSubscriptionIds returns all subscription ids
func (c *Client) GetSubscriptionIds(ctx context.Context, conState jsonrpc.State) ([]string, error) {
	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.SubscriptionIds(), nil
}

// GetEvents returns all events for all subscriptions
func (c *Client) GetEvents(ctx context.Context, conState jsonrpc.State) ([]nostr.NostrEvent, error) {
	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	evs := state.client.GetEvents()

	return evs, nil
}
