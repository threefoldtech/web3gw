package tfgrid

import (
	"context"

	"github.com/pkg/errors"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/workloads"
	"github.com/threefoldtech/zos/pkg/gridtypes"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func (r *Client) deployNetwork(ctx context.Context, modelName string, nodes []uint32, IPRange string, WGAccess bool, projectName string) (*workloads.ZNet, error) {
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

func doesNetworkIncludeNode(networkNodes []uint32, nodeID uint32) bool {
	for _, node := range networkNodes {
		if node == nodeID {
			return true
		}
	}

	return false
}

func (c *Client) ensureNodeBelongsToNetwork(ctx context.Context, networkName string, networkContracts map[uint32]uint64, nodeID uint32) error {
	znet, err := c.client.LoadNetwork(networkName, networkContracts)
	if err != nil {
		return errors.Wrapf(err, "failed to load network %s", networkName)
	}

	if !doesNetworkIncludeNode(znet.Nodes, nodeID) {
		znet.Nodes = append(znet.Nodes, nodeID)
		err = c.client.DeployNetwork(ctx, &znet)
		if err != nil {
			return errors.Wrap(err, "failed to deploy network")
		}
	}

	return nil
}
