package explorer

import (
	"context"
	"errors"

	"github.com/threefoldtech/grid_proxy_server/pkg/types"
	proxyTypes "github.com/threefoldtech/grid_proxy_server/pkg/types"
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

func (c *Client) Ping(ctx context.Context) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.cl == nil {
		return errors.New("Error pinging")
	}

	return state.cl.Ping()
}

func (c *Client) Nodes(ctx context.Context, filter proxyTypes.NodeFilter, pagination proxyTypes.Limit) ([]types.Node, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.cl == nil {
		return []types.Node{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.Nodes(ctx, filter, pagination)
}

func (c *Client) Farms(ctx context.Context, filter proxyTypes.FarmFilter, pagination proxyTypes.Limit) ([]types.Farm, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.cl == nil {
		return []types.Farm{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.Farms(ctx, filter, pagination)
}
func (c *Client) Contracts(ctx context.Context, filter proxyTypes.ContractFilter, pagination proxyTypes.Limit) ([]types.Contract, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.cl == nil {
		return []types.Contract{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.Contracts(ctx, filter, pagination)
}
func (c *Client) Twins(ctx context.Context, filter proxyTypes.TwinFilter, pagination proxyTypes.Limit) ([]types.Twin, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.cl == nil {
		return []types.Twin{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.Twins(ctx, filter, pagination)
}

func (c *Client) Node(ctx context.Context, nodeID uint32) (res proxyTypes.NodeWithNestedCapacity, err error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.cl == nil {
		return types.NodeWithNestedCapacity{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.Node(nodeID)
}

func (c *Client) NodeStatus(ctx context.Context, nodeID uint32) (res proxyTypes.NodeStatus, err error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.cl == nil {
		return proxyTypes.NodeStatus{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.NodeStatus(nodeID)
}

func (c *Client) Counters(ctx context.Context, filter proxyTypes.StatsFilter) (res proxyTypes.Counters, err error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.cl == nil {
		return proxyTypes.Counters{}, pkg.ErrClientNotConnected{}
	}
	return state.cl.Counters(filter)
}
