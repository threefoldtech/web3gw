package explorer

import (
	proxy "github.com/threefoldtech/tfgrid-sdk-go/grid-proxy/pkg/client"
	proxyTypes "github.com/threefoldtech/tfgrid-sdk-go/grid-proxy/pkg/types"
)

var (
	endpoints = map[string]string{
		"main": "https://gridproxy.grid.tf",
		"test": "https://gridproxy.test.grid.tf",
		"qa":   "https://gridproxy.qa.grid.tf",
		"dev":  "https://gridproxy.dev.grid.tf",
	}
)

type ExplorerClient struct {
	proxyClient proxy.Client
}

func NewClient() *ExplorerClient {
	return &ExplorerClient{}
}

func (c *ExplorerClient) Load(net string) {
	endpoint := endpoints["dev"]
	if res, ok := endpoints[net]; ok {
		endpoint = res
	}
	c.proxyClient = proxy.NewClient(endpoint)
}

func (c *ExplorerClient) Ping() error {
	return c.proxyClient.Ping()
}

func (c *ExplorerClient) Nodes(params NodesRequestParams) (NodesResult, error) {
	nodes, count, err := c.proxyClient.Nodes(params.Filters, params.Pagination)
	res := NodesResult{
		Nodes:      nodes,
		TotalCount: count,
	}
	return res, err
}

func (c *ExplorerClient) Farms(params FarmsRequestParams) (FarmsResult, error) {
	farms, count, err := c.proxyClient.Farms(params.Filters, params.Pagination)
	res := FarmsResult{
		Farms:      farms,
		TotalCount: count,
	}
	return res, err
}

func (c *ExplorerClient) Contracts(params ContractsRequestParams) (ContractsResult, error) {
	contracts, count, err := c.proxyClient.Contracts(params.Filters, params.Pagination)
	res := ContractsResult{
		Contracts:  contracts,
		TotalCount: count,
	}
	return res, err
}

func (c *ExplorerClient) Twins(params TwinsRequestParams) (TwinsResult, error) {
	twins, count, err := c.proxyClient.Twins(params.Filters, params.Pagination)
	res := TwinsResult{
		Twins:      twins,
		TotalCount: count,
	}
	return res, err
}

func (c *ExplorerClient) Node(nodeID uint32) (proxyTypes.NodeWithNestedCapacity, error) {
	return c.proxyClient.Node(nodeID)
}

func (c *ExplorerClient) NodeStatus(nodeID uint32) (proxyTypes.NodeStatus, error) {
	return c.proxyClient.NodeStatus(nodeID)
}

func (c *ExplorerClient) Counters(filters proxyTypes.StatsFilter) (proxyTypes.Counters, error) {
	return c.proxyClient.Counters(filters)
}
