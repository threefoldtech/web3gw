package explorer

import (
	"context"
	"fmt"

	"github.com/LeeSmet/go-jsonrpc"
	proxy "github.com/threefoldtech/tfgrid-sdk-go/grid-proxy/pkg/client"
	proxyTypes "github.com/threefoldtech/tfgrid-sdk-go/grid-proxy/pkg/types"
	"github.com/threefoldtech/web3_proxy/server/pkg"
	"github.com/threefoldtech/web3_proxy/server/pkg/state"
)

var (
	endpoints = map[string]string{
		"main": "https://gridproxy.grid.tf",
		"test": "https://gridproxy.test.grid.tf",
		"qa":   "https://gridproxy.qa.grid.tf",
		"dev":  "https://gridproxy.dev.grid.tf",
	}
)

const ExplorerID = "explorer"

type (
	Client struct {
		state *state.StateManager[*explorerState]
	}

	explorerState struct {
		cl proxy.Client
	}
)

func (e *explorerState) Close() {

}

func NewClient() *Client {
	return &Client{
		state: state.NewStateManager[*explorerState](),
	}
}

// State from a connection. If no state is present, it is initialized
func State(conState jsonrpc.State) *explorerState {
	raw, exists := conState[ExplorerID]
	if !exists {
		ns := &explorerState{
			cl: nil,
		}
		conState[ExplorerID] = ns
		return ns
	}

	ns, ok := raw.(*explorerState)
	if !ok {
		// This means the invariant is violated, so panic here is ok
		panic("Invalid saved state for explorer")
	}
	return ns
}

// Load an identity for the explorer with the given network
func (c *Client) Load(ctx context.Context, conState jsonrpc.State, network string) error {
	state := State(conState)
	if state.cl != nil {
		state.Close()
	}

	endpoint, ok := endpoints[network]
	if !ok {
		return fmt.Errorf("network %s is not supported", network)
	}

	state.cl = proxy.NewClient(endpoint)

	return nil
}

func (c *Client) Ping(ctx context.Context, conState jsonrpc.State) error {
	state := State(conState)
	if state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.cl.Ping()
}

func (c *Client) Nodes(ctx context.Context, conState jsonrpc.State, params NodesRequestParams) (NodesResult, error) {
	state := State(conState)
	if state.cl == nil {
		return NodesResult{}, pkg.ErrClientNotConnected{}
	}

	nodes, count, err := state.cl.Nodes(params.Filters, params.Pagination)
	res := NodesResult{
		Nodes:      nodes,
		TotalCount: count,
	}
	return res, err
}

func (c *Client) Farms(ctx context.Context, conState jsonrpc.State, params FarmsRequestParams) (FarmsResult, error) {
	state := State(conState)
	if state.cl == nil {
		return FarmsResult{}, pkg.ErrClientNotConnected{}
	}

	farms, count, err := state.cl.Farms(params.Filters, params.Pagination)
	res := FarmsResult{
		Farms:      farms,
		TotalCount: count,
	}
	return res, err
}

func (c *Client) Contracts(ctx context.Context, conState jsonrpc.State, params ContractsRequestParams) (ContractsResult, error) {
	state := State(conState)
	if state.cl == nil {
		return ContractsResult{}, pkg.ErrClientNotConnected{}
	}

	contracts, count, err := state.cl.Contracts(params.Filters, params.Pagination)
	res := ContractsResult{
		Contracts:  contracts,
		TotalCount: count,
	}
	return res, err
}

func (c *Client) Twins(ctx context.Context, conState jsonrpc.State, params TwinsRequestParams) (TwinsResult, error) {
	state := State(conState)
	if state.cl == nil {
		return TwinsResult{}, pkg.ErrClientNotConnected{}
	}

	twins, count, err := state.cl.Twins(params.Filters, params.Pagination)
	res := TwinsResult{
		Twins:      twins,
		TotalCount: count,
	}
	return res, err
}

func (c *Client) Node(ctx context.Context, conState jsonrpc.State, nodeID uint32) (proxyTypes.NodeWithNestedCapacity, error) {
	state := State(conState)
	if state.cl == nil {
		return proxyTypes.NodeWithNestedCapacity{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.Node(nodeID)
}

func (c *Client) NodeStatus(ctx context.Context, conState jsonrpc.State, nodeID uint32) (proxyTypes.NodeStatus, error) {
	state := State(conState)
	if state.cl == nil {
		return proxyTypes.NodeStatus{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.NodeStatus(nodeID)
}

func (c *Client) Counters(ctx context.Context, conState jsonrpc.State, filters proxyTypes.StatsFilter) (proxyTypes.Counters, error) {
	state := State(conState)
	if state.cl == nil {
		return proxyTypes.Counters{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.Counters(filters)
}
