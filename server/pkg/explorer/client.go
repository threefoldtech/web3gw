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

func (c *ExplorerClient) Load(net string) error {
	endpoint := endpoints["dev"]
	if res, ok := endpoints[net]; ok {
		endpoint = res
	}
	c.proxyClient = proxy.NewClient(endpoint)
	return nil
}

func (c *ExplorerClient) Ping() error {
	return c.proxyClient.Ping()
}

func (c *ExplorerClient) Nodes(params NodesRequestParams) ([]proxyTypes.Node, error) {
	nodes, _, err := c.proxyClient.Nodes(params.Filters, params.Pagination)
	return nodes, err
}

func (c *ExplorerClient) Farms(params FarmsRequestParams) ([]proxyTypes.Farm, error) {
	farms, _, err := c.proxyClient.Farms(params.Filters, params.Pagination)
	return farms, err
}

func (c *ExplorerClient) Contracts(params ContractsRequestParams) ([]proxyTypes.Contract, error) {
	contracts, _, err := c.proxyClient.Contracts(params.Filters, params.Pagination)
	return contracts, err
}

func (c *ExplorerClient) Twins(params TwinsRequestParams) ([]proxyTypes.Twin, error) {
	twins, _, err := c.proxyClient.Twins(params.Filters, params.Pagination)
	return twins, err
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
