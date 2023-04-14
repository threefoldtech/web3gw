package nostr

import (
	"context"
	"fmt"

	"github.com/threefoldtech/web3_proxy/server/clients/nostr"
	"github.com/threefoldtech/web3_proxy/server/pkg"
	"github.com/threefoldtech/web3_proxy/server/pkg/state"
)

type (
	// Client exposes nostr related functionality
	Client struct {
		state *state.StateManager[nostrState]
	}
	// state managed by nostr client
	nostrState struct {
		client *nostr.Client
	}
)

func NewClient() *Client {
	return &Client{
		state: state.NewStateManager[nostrState](),
	}
}

func (c *Client) Load(ctx context.Context, secret string) error {
	srv := nostr.NewServer()

	cl, err := srv.NewClient(secret)
	if err != nil {
		return err
	}

	ns := nostrState{
		client: cl,
	}

	c.state.Set(state.IDFromContext(ctx), ns)

	return nil
}

func (c *Client) ConnectAuthRelay(ctx context.Context, url string) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ConnectAuthRelay(ctx, url)
}

func (c *Client) ConnectRelay(ctx context.Context, url string) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ConnectRelay(ctx, url)
}

func (c *Client) GenerateKeyPair(ctx context.Context) (string, error) {
	return nostr.GenerateKeyPair(), nil
}

func (c *Client) ConnectToRelay(ctx context.Context, url string) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ConnectAuthRelay(ctx, url)
}

type EventInput struct {
	Tags    []string `json:"tags"`
	Content string   `json:"content"`
}

func (c *Client) PublishEventToRelays(ctx context.Context, input EventInput) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.PublishEventToRelays(ctx, input.Tags, input.Content)
}

func (c *Client) SubscribeRelays(ctx context.Context) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.SubscribeRelays()
}

func (c *Client) CloseSubscription(ctx context.Context, id string) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	state.client.CloseSubscription(id)

	return nil
}

func (c *Client) GetSubscriptionIds(ctx context.Context) ([]string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.SubscriptionIds(), nil
}

func (c *Client) GetEvents(ctx context.Context) ([]nostr.NostrEvent, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	evs := state.client.GetEvents()

	fmt.Printf("events: %v", evs)

	return evs, nil
}
