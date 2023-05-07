package tfgrid

import (
	"encoding/json"
	"strconv"

	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/graphql"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/workloads"
)

type contractsInfo struct {
	networkContracts    map[uint32]uint64
	deploymentContracts map[uint32][]uint64
}

func getContractsInfo(contracts graphql.Contracts, networkName string) (contractsInfo, error) {
	networkState := map[uint32]uint64{}
	deploymentState := map[uint32][]uint64{}
	for _, c := range contracts.NodeContracts {
		var deploymentData workloads.DeploymentData
		err := json.Unmarshal([]byte(c.DeploymentData), &deploymentData)
		if err != nil {
			return contractsInfo{}, err
		}

		contractID, err := strconv.ParseUint(c.ContractID, 10, 64)
		if err != nil {
			return contractsInfo{}, err
		}

		if deploymentData.Type != "network" || deploymentData.Name != networkName {
			deploymentState[c.NodeID] = append(deploymentState[c.NodeID], contractID)
			continue
		}

		networkState[c.NodeID] = contractID
	}

	return contractsInfo{
		networkContracts:    networkState,
		deploymentContracts: deploymentState,
	}, nil
}

func (c *contractsInfo) getNodeIDs() []uint32 {
	idSet := map[uint32]struct{}{}
	nodeIDs := []uint32{}
	for nodeID := range c.deploymentContracts {
		if _, ok := idSet[nodeID]; !ok {
			idSet[nodeID] = struct{}{}
			nodeIDs = append(nodeIDs, nodeID)
		}
	}
	for nodeID := range c.networkContracts {
		if _, ok := idSet[nodeID]; !ok {
			idSet[nodeID] = struct{}{}
			nodeIDs = append(nodeIDs, nodeID)
		}
	}
	return nodeIDs
}
