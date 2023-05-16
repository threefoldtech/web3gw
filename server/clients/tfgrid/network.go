package tfgrid

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/workloads"
	"github.com/threefoldtech/zos/pkg/gridtypes"
	"golang.org/x/exp/slices"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func (c *Client) deployZnet(ctx context.Context, znet *workloads.ZNet) error {
	if znet.AddWGAccess {
		privateKey, err := wgtypes.GeneratePrivateKey()
		if err != nil {
			return errors.Wrap(err, "failed to generate wireguard private key")
		}
		znet.ExternalSK = privateKey
	}

	if err := c.client.DeployNetwork(ctx, znet); err != nil {
		return errors.Wrap(err, "failed to deploy network")
	}

	return nil
}

func (r *Client) deployNetwork(ctx context.Context, modelName string, nodes []uint32, IPRange string, WGAccess bool) (*workloads.ZNet, error) {
	projectName := generateProjectName(modelName)

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

func (c *Client) ensureNodeBelongsToNetwork(ctx context.Context, znet *workloads.ZNet, nodeID uint32) error {
	log.Info().Msgf("ensuring node in network: %+v", znet)

	if !slices.Contains(znet.Nodes, nodeID) {
		znet.Nodes = append(znet.Nodes, nodeID)
		err := c.client.DeployNetwork(ctx, znet)
		if err != nil {
			return errors.Wrap(err, "failed to deploy network")
		}
	}

	return nil
}

func (c *Client) removeNodeFromNetwork(ctx context.Context, znet *workloads.ZNet, nodeID uint32) error {
	for idx, node := range znet.Nodes {
		if node == nodeID {
			znet.Nodes = append(znet.Nodes[:idx], znet.Nodes[idx+1:]...)
			return c.client.DeployNetwork(ctx, znet)
		}
	}

	return nil
}

func generateNetworkName(modelName string) string {
	return fmt.Sprintf("%s_network", modelName)
}
