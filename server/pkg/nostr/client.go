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

func (c *Client) Load(ctx context.Context, secret string, relayUrl string) error {
	srv := nostr.NewServer()

	cl, err := srv.NewClient(secret)
	if err != nil {
		return err
	}

	err = cl.ConnectAuthRelay(ctx, relayUrl)
	if err != nil {
		return err
	}

	ns := nostrState{
		client: cl,
	}

	c.state.Set(state.IDFromContext(ctx), ns)

	return nil
}

func (c *Client) ConnectToRelay(ctx context.Context, url string) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ConnectAuthRelay(ctx, url)
}

func (c *Client) PublishEventToRelays(ctx context.Context, tags []string, content string) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.PublishEventToRelays(ctx, tags, content)
}
