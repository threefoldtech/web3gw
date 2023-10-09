package tfgrid

import (
	"context"

	"github.com/threefoldtech/tfgrid-sdk-go/grid-proxy/pkg/types"
)

type StatsFilterOptions struct {
	Status string
}

type CountersResult struct {
	Nodes             int64            `json:"nodes"`
	Farms             int64            `json:"farms"`
	Countries         int64            `json:"countries"`
	TotalCRU          int64            `json:"totalCru"`
	TotalSRU          int64            `json:"totalSru"`
	TotalMRU          int64            `json:"totalMru"`
	TotalHRU          int64            `json:"totalHru"`
	PublicIPs         int64            `json:"publicIps"`
	AccessNodes       int64            `json:"accessNodes"`
	Gateways          int64            `json:"gateways"`
	Twins             int64            `json:"twins"`
	Contracts         int64            `json:"contracts"`
	NodesDistribution map[string]int64 `json:"nodesDistribution" gorm:"-:all"`
	GPUs              int64            `json:"gpus"`
}

func toCountersResult(res types.Counters) CountersResult {
	return CountersResult{
		Nodes:             res.Nodes,
		Farms:             res.Farms,
		Countries:         res.Countries,
		TotalCRU:          res.TotalCRU,
		TotalSRU:          res.TotalSRU,
		TotalMRU:          res.TotalMRU,
		TotalHRU:          res.TotalHRU,
		PublicIPs:         res.PublicIPs,
		AccessNodes:       res.AccessNodes,
		Gateways:          res.Gateways,
		Twins:             res.Twins,
		Contracts:         res.Contracts,
		NodesDistribution: res.NodesDistribution,
		GPUs:              res.GPUs,
	}
}

func (r *Client) GetStatistics(ctx context.Context, options StatsFilterOptions) (CountersResult, error) {
	filter := types.StatsFilter{
		Status: &options.Status,
	}

	counters, err := r.GridClient.GetCounters(filter)
	return toCountersResult(counters), err

}
