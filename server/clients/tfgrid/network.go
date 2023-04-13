package tfgrid

import (
	"context"

	"github.com/pkg/errors"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/workloads"
	"github.com/threefoldtech/zos/pkg/gridtypes"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func (r *Runner) deployNetwork(ctx context.Context, modelName string, nodes []uint32, IPRange string, WGAccess bool, projectName string) (*workloads.ZNet, error) {
	nodeList := []uint32{}
	nodeSet := map[uint32]struct{}{}
	for _, node := range nodes {
		if _, ok := nodeSet[node]; !ok {
			nodeList = append(nodeList, node)
			nodeSet[node] = struct{}{}
		}
	}

	ipRange, err := gridtypes.ParseIPNet(IPRange)
	if err != nil {
		return nil, errors.Wrapf(err, "network ip range (%s) is invalid", IPRange)
	}

	znet := workloads.ZNet{
		Name:         generateNetworkName(modelName),
		Nodes:        nodeList,
		IPRange:      ipRange,
		AddWGAccess:  WGAccess,
		SolutionType: projectName,
	}

	if znet.AddWGAccess {
		privateKey, err := wgtypes.GeneratePrivateKey()
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate wireguard private key")
		}
		znet.ExternalSK = privateKey
	}

	err = r.client.DeployNetwork(ctx, &znet)
	if err != nil {
		return nil, errors.Wrap(err, "failed to deploy network")
	}

	return &znet, nil
}
