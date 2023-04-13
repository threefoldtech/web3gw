package nostr

import (
	"context"

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

type Input struct {
	tags    []string
	content string
}

func (c *Client) PublishEventToRelays(ctx context.Context, input Input) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.PublishEventToRelays(ctx, input.tags, input.content)
}

func (c *Client) SubscribeRelays(ctx context.Context) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.SubscribeRelays(ctx)
}

func (c *Client) GetEvents(ctx context.Context) ([]*nostr.NostrEvent, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetEvents(), nil
}
