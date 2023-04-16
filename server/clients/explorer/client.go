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

type ExplorerClient struct {
	client proxy.Client
}

func (ec *ExplorerClient) Load(net string) {
	ec.client = proxy.NewClient(endpoints[net])
}

func (ec *ExplorerClient) Nodes(ctx context.Context) ([]proxyTypes.Node, error) {
	res, _, err := ec.client.Nodes(proxyTypes.NodeFilter{}, proxyTypes.Limit{})
	return res, err
}
