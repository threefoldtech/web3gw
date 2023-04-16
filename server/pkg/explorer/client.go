package explorer

import (
	"context"

	"github.com/threefoldtech/grid_proxy_server/pkg/types"
	"github.com/threefoldtech/web3_proxy/server/clients/explorer"
	"github.com/threefoldtech/web3_proxy/server/pkg"
	"github.com/threefoldtech/web3_proxy/server/pkg/state"
)

type (
	Client struct {
		state *state.StateManager[explorerState]
	}

	explorerState struct {
		cl *explorer.ExplorerClient
	}
)

func NewClient() *Client {
	return &Client{
		state: state.NewStateManager[explorerState](),
	}
}

func (c *Client) Load(ctx context.Context, net string) error {
	gpc := explorer.ExplorerClient{}
	gpc.Load(net)

	gs := explorerState{
		cl: &gpc,
	}

	c.state.Set(state.IDFromContext(ctx), gs)

	return nil
}

func (c *Client) Nodes(ctx context.Context) ([]types.Node, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.cl == nil {
		return []types.Node{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.Nodes(ctx)
}
