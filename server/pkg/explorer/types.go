package explorer

import proxyTypes "github.com/threefoldtech/tfgrid-sdk-go/grid-proxy/pkg/types"

type NodesRequestParams struct {
	Filters    proxyTypes.NodeFilter `json:"filters"`
	Pagination proxyTypes.Limit      `json:"limit"`
}

type FarmsRequestParams struct {
	Filters    proxyTypes.FarmFilter `json:"filters"`
	Pagination proxyTypes.Limit      `json:"limit"`
}

type TwinsRequestParams struct {
	Filters    proxyTypes.TwinFilter `json:"filters"`
	Pagination proxyTypes.Limit      `json:"limit"`
}

type ContractsRequestParams struct {
	Filters    proxyTypes.ContractFilter `json:"filters"`
	Pagination proxyTypes.Limit          `json:"limit"`
}
