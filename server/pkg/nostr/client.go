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

func NewClient() *Client {
	return &Client{
		server: nostr.NewServer(),
	}
}

func (c *Client) Load(ctx context.Context, conState jsonrpc.State, secret string) error {
	cl, err := c.server.NewClient(secret)
	if err != nil {
		return err
	}

	state := State(conState)
	state.client = cl

	return nil
}

func (c *Client) ConnectAuthRelay(ctx context.Context, conState jsonrpc.State, url string) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ConnectAuthRelay(ctx, url)
}

func (c *Client) ConnectRelay(ctx context.Context, conState jsonrpc.State, url string) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ConnectRelay(ctx, url)
}

func (c *Client) GenerateKeyPair(ctx context.Context) (string, error) {
	return nostr.GenerateKeyPair(), nil
}

func (c *Client) ConnectToRelay(ctx context.Context, conState jsonrpc.State, url string) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ConnectAuthRelay(ctx, url)
}

type EventInput struct {
	Tags    []string `json:"tags"`
	Content string   `json:"content"`
}

func (c *Client) PublishEventToRelays(ctx context.Context, conState jsonrpc.State, input EventInput) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.PublishTextNote(ctx, input.Tags, input.Content)
}

func (c *Client) SubscribeRelays(ctx context.Context, conState jsonrpc.State) (string, error) {
	state := State(conState)
	if state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.SubscribeRelays()
}

func (c *Client) CloseSubscription(ctx context.Context, conState jsonrpc.State, id string) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	state.client.CloseSubscription(id)

	return nil
}

func (c *Client) GetSubscriptionIds(ctx context.Context, conState jsonrpc.State) ([]string, error) {
	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.SubscriptionIds(), nil
}

func (c *Client) GetEvents(ctx context.Context, conState jsonrpc.State) ([]nostr.NostrEvent, error) {
	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	evs := state.client.GetEvents()

	return evs, nil
}