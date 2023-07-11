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

	FetchChannelMessageInput struct {
		ChannelId string `json:"channel_id"`
	}

	SubscribeChannelMessageInput struct {
		// ID of the channel or message for which the reply is intended
		ID string `json:"id"`
	}

	CreateChannelMessageInput struct {
		ChannelID string `json:"channel_id"`
		Content   string `json:"content"`
		// MessageID is used for replies
		MessageID string `json:"message_id"`
		// PublicKey of author to reploy to
		PublicKey string `json:"public_key"`
	}

	CreateChannelInput struct {
		Tags    []string `json:"tags"`
		Name    string   `json:"name"`
		About   string   `json:"about"`
		Picture string   `json:"picture"`
	}

	ProductInput struct {
		Tags    []string      `json:"tags"`
		Product nostr.Product `json:"product"`
	}

	StallInput struct {
		Tags  []string    `json:"tags"`
		Stall nostr.Stall `json:"stall"`
	}
	DirectMessageInput struct {
		Receiver string   `json:"receiver"`
		Tags     []string `json:"tags"`
		Content  string   `json:"content"`
	}

	// TextNote is a text note published on a relay
	TextInput struct {
		Tags    []string `json:"tags"`
		Content string   `json:"content"`
	}

	// MetadataInput is metadata published on a relay
	MetadataInput struct {
		Tags     []string       `json:"tags"`
		Metadata nostr.Metadata `json:"metadata"`
	}

	// GetSubscriptionEventsInput specifies subscription events retrieval information
	GetSubscriptionEventsInput struct {
		ID    string `json:"id"`
		Count uint32 `json:"count"`
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

// Close implements jsonrpc.Closer
func (s *NostrState) Close() {}

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

// PublishTextNote publishes a text note to all relays
func (c *Client) PublishTextNote(ctx context.Context, conState jsonrpc.State, input TextInput) error {
	state := State(conState)

	if state.Client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.Client.PublishTextNote(ctx, input.Tags, input.Content)
}

// PublishMetadata publishes metadata to all relays
func (c *Client) PublishMetadata(ctx context.Context, conState jsonrpc.State, input MetadataInput) error {
	state := State(conState)
	if state.Client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.Client.PublishMetadata(ctx, input.Tags, input.Metadata)
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

// GetSubscriptionEvents returns all events for a subscription with the specified id
func (c *Client) GetSubscriptionEvents(ctx context.Context, conState jsonrpc.State, args GetSubscriptionEventsInput) ([]nostr.NostrEvent, error) {
	state := State(conState)
	if state.Client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	evs := state.Client.GetSubscriptionEventsWithCount(args.ID, args.Count)

	return evs, nil
}

// PublishStall publishes a new stall to the relay
func (c *Client) PublishStall(ctx context.Context, conState jsonrpc.State, input StallInput) error {
	state := State(conState)
	if state.Client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.Client.PublishStall(ctx, input.Tags, input.Stall)
}

// PublishProduct publishes a new product to the relay
func (c *Client) PublishProduct(ctx context.Context, conState jsonrpc.State, input ProductInput) error {
	state := State(conState)
	if state.Client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.Client.PublishProduct(ctx, input.Tags, input.Product)
}

// CreateChannel creates a new channel
func (c *Client) CreateChannel(ctx context.Context, conState jsonrpc.State, input CreateChannelInput) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.CreateChannel(ctx, input.Tags, nostr.Channel{Name: input.Name, About: input.About, Picture: input.Picture})
}

// SubscribeChannelCreation subscribes to channel creation events on the relay
func (c *Client) SubscribeChannelCreation(ctx context.Context, conState jsonrpc.State) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.SubscribeChannelCreation()
}

// CreateChannelMessage creates a channel message
func (c *Client) CreateChannelMessage(ctx context.Context, conState jsonrpc.State, input CreateChannelMessageInput) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.CreateChannelRootMessage(ctx, nostr.ChannelMessage{Content: input.Content, ChannelID: input.ChannelID, MessageID: input.MessageID, PublicKey: input.PublicKey})
}

// SubscribeChannelMessage subscribes to a channel messages or message replies, depending on the the id provided
func (c *Client) SubscribeChannelMessage(ctx context.Context, conState jsonrpc.State, input SubscribeChannelMessageInput) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.SubscribeChannelMessages(input.ID)
}

// ListChannels on connected relays
func (c *Client) ListChannels(ctx context.Context, conState jsonrpc.State) ([]nostr.RelayChannel, error) {
	state := State(conState)
	if state.Client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.Client.FetchChannelCreation()
}

// GetChannelMessages returns channel messages
func (c *Client) GetChannelMessages(ctx context.Context, conState jsonrpc.State, input FetchChannelMessageInput) ([]nostr.RelayChannelMessage, error) {
	state := State(conState)
	if state.Client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.Client.FetchChannelMessages(input.ChannelId)
}
