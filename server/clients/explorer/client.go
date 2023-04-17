package explorer

import (
	"context"

	proxy "github.com/threefoldtech/grid_proxy_server/pkg/client"
	proxyTypes "github.com/threefoldtech/grid_proxy_server/pkg/types"
)

var (
	endpoints = map[string]string{
		"dev":  "https://gridproxy.grid.tf",
		"qa":   "https://gridproxy.test.grid.tf",
		"test": "https://gridproxy.qa.grid.tf",
		"main": "https://gridproxy.dev.grid.tf",
	}
)

type Explorer interface {
	Ping() error
	Nodes(filter proxyTypes.NodeFilter, pagination proxyTypes.Limit) (res []proxyTypes.Node, totalCount int, err error)
	Farms(filter proxyTypes.FarmFilter, pagination proxyTypes.Limit) (res []proxyTypes.Farm, totalCount int, err error)
	Contracts(filter proxyTypes.ContractFilter, pagination proxyTypes.Limit) (res []proxyTypes.Contract, totalCount int, err error)
	Twins(filter proxyTypes.TwinFilter, pagination proxyTypes.Limit) (res []proxyTypes.Twin, totalCount int, err error)
	Node(nodeID uint32) (res proxyTypes.NodeWithNestedCapacity, err error)
	NodeStatus(nodeID uint32) (res proxyTypes.NodeStatus, err error)
	Counters(filter proxyTypes.StatsFilter) (res proxyTypes.Counters, err error)
}

type ExplorerClient struct {
	client proxy.Client
}

func (ec *ExplorerClient) Load(net string) {
	ec.client = proxy.NewClient(endpoints[net])
}

func (ec *ExplorerClient) Ping() error {
	return ec.client.Ping()
}

func (ec *ExplorerClient) Nodes(ctx context.Context, filter proxyTypes.NodeFilter, pagination proxyTypes.Limit) ([]proxyTypes.Node, error) {
	res, _, err := ec.client.Nodes(proxyTypes.NodeFilter{}, proxyTypes.Limit{})
	return res, err
}

func (ec *ExplorerClient) Farms(ctx context.Context, filter proxyTypes.FarmFilter, pagination proxyTypes.Limit) ([]proxyTypes.Farm, error) {
	res, _, err := ec.client.Farms(proxyTypes.FarmFilter{}, proxyTypes.Limit{})
	return res, err
}

func (ec *ExplorerClient) Contracts(ctx context.Context, filter proxyTypes.ContractFilter, pagination proxyTypes.Limit) ([]proxyTypes.Contract, error) {
	res, _, err := ec.client.Contracts(proxyTypes.ContractFilter{}, proxyTypes.Limit{})
	return res, err
}

func (ec *ExplorerClient) Twins(ctx context.Context, filter proxyTypes.TwinFilter, pagination proxyTypes.Limit) ([]proxyTypes.Twin, error) {
	res, _, err := ec.client.Twins(proxyTypes.TwinFilter{}, proxyTypes.Limit{})
	return res, err
}

func (ec *ExplorerClient) Node(nodeID uint32) (res proxyTypes.NodeWithNestedCapacity, err error) {
	return ec.client.Node(nodeID)
}

func (ec *ExplorerClient) NodeStatus(nodeID uint32) (res proxyTypes.NodeStatus, err error) {
	return ec.client.NodeStatus(nodeID)
}

func (ec *ExplorerClient) Counters(filter proxyTypes.StatsFilter) (res proxyTypes.Counters, err error) {
	return ec.client.Counters(filter)
}
