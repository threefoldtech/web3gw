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
	NostrState struct {
		Client *nostr.Client
	}
)

// State from a connection. If no state is present, it is initialized
func State(conState jsonrpc.State) *NostrState {
	raw, exists := conState[NostrID]
	if !exists {
		ns := &NostrState{
			Client: nil,
		}
		conState[NostrID] = ns
		return ns
	}
	ns, ok := raw.(*NostrState)
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
	state.Client = cl

	return nil
}

// GetPublicKey returns the nostr ID for the client
func (c *Client) GetId(ctx context.Context, conState jsonrpc.State) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.Id(), nil
}

// GetPublicKey returns the public key of the client in hex
func (c *Client) GetPublicKey(ctx context.Context, conState jsonrpc.State) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.PublicKey(), nil
}

// ConnectRelay connects to an authenticated relay with a given url
func (c *Client) ConnectAuthRelay(ctx context.Context, conState jsonrpc.State, url string) error {
	state := State(conState)
	if state.Client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.Client.ConnectAuthRelay(ctx, url)
}

// ConnectRelay connects to a relay with a given url
func (c *Client) ConnectRelay(ctx context.Context, conState jsonrpc.State, url string) error {
	state := State(conState)
	if state.Client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.Client.ConnectRelay(ctx, url)
}

// GenerateKeyPair generates a new keypair
func (c *Client) GenerateKeyPair(ctx context.Context) (string, error) {
	return nostr.GenerateKeyPair(), nil
}

// ConnectToRelay connects to a relay with a given url
func (c *Client) ConnectToRelay(ctx context.Context, conState jsonrpc.State, url string) error {
	state := State(conState)
	if state.Client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.Client.ConnectAuthRelay(ctx, url)
}

// TextNote is a text note published on a relay
type TextInput struct {
	Tags    []string `json:"tags"`
	Content string   `json:"content"`
}

// PublishTextNote publishes a text note to all relays
func (c *Client) PublishTextNote(ctx context.Context, conState jsonrpc.State, input TextInput) error {
	state := State(conState)

	if state.Client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.Client.PublishTextNote(ctx, input.Tags, input.Content)
}

// MetadataInput is metadata published on a relay
type MetadataInput struct {
	Tags     []string       `json:"tags"`
	Metadata nostr.Metadata `json:"metadata"`
}

// PublishMetadata publishes metadata to all relays
func (c *Client) PublishMetadata(ctx context.Context, conState jsonrpc.State, input MetadataInput) error {
	state := State(conState)
	if state.Client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.Client.PublishMetadata(ctx, input.Tags, input.Metadata)
}

type DirectMessageInput struct {
	Receiver string   `json:"receiver"`
	Tags     []string `json:"tags"`
	Content  string   `json:"content"`
}

// PublishDirectMessage publishes a direct message to a receiver
func (c *Client) PublishDirectMessage(ctx context.Context, conState jsonrpc.State, input DirectMessageInput) error {
	state := State(conState)
	if state.Client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.Client.PublishDirectMessage(ctx, input.Receiver, input.Tags, input.Content)
}

// SubscribeTextNotes subscribes to text notes on all relays
func (c *Client) SubscribeTextNotes(ctx context.Context, conState jsonrpc.State) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.SubscribeTextNotes()
}

// SubscribeDirectMessages subscribes to direct messages on all relays and decrypts them
func (c *Client) SubscribeDirectMessages(ctx context.Context, conState jsonrpc.State) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.SubscribeMessages()
}

// SubscribeStallCreation subscribes to stall creation on all relays
func (c *Client) SubscribeStallCreation(ctx context.Context, conState jsonrpc.State) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.SubscribeStallCreation("")
}

// SubscribeProductCreation subscribes to product creation on all relays
func (c *Client) SubscribeProductCreation(ctx context.Context, conState jsonrpc.State) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.SubscribeProductCreation("")
}

// CloseSubscription closes a subscription by id
func (c *Client) CloseSubscription(ctx context.Context, conState jsonrpc.State, id string) error {
	state := State(conState)
	if state.Client == nil {
		return pkg.ErrClientNotConnected{}
	}

	state.Client.CloseSubscription(id)

	return nil
}

// GetSubscriptionIds returns all subscription ids
func (c *Client) GetSubscriptionIds(ctx context.Context, conState jsonrpc.State) ([]string, error) {
	state := State(conState)
	if state.Client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.Client.SubscriptionIds(), nil
}

// GetEvents returns all events for all subscriptions
func (c *Client) GetEvents(ctx context.Context, conState jsonrpc.State) ([]nostr.NostrEvent, error) {
	state := State(conState)
	if state.Client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	evs := state.Client.GetEvents()

	return evs, nil
}

type StallInput struct {
	Tags  []string    `json:"tags"`
	Stall nostr.Stall `json:"stall"`
}

func (c *Client) PublishStall(ctx context.Context, conState jsonrpc.State, input StallInput) error {
	state := State(conState)
	if state.Client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.Client.PublishStall(ctx, input.Tags, input.Stall)
}

type ProductInput struct {
	Tags    []string      `json:"tags"`
	Product nostr.Product `json:"product"`
}

func (c *Client) PublishProduct(ctx context.Context, conState jsonrpc.State, input ProductInput) error {
	state := State(conState)
	if state.Client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.Client.PublishProduct(ctx, input.Tags, input.Product)
}
