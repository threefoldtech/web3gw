package tfgrid

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/state"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/workloads"
)

type ProjectState struct {
	nodeContracts map[uint32]state.ContractIDs
	// used in gateway names
	nameContracts map[uint32]uint64
}

func (c *Client) loadNetwork(modelName string) (workloads.ZNet, error) {
	return c.GridClient.LoadNetwork(generateNetworkName(modelName))
}

func (c *Client) loadK8s(masterName string, nodeContracts map[uint32]state.ContractIDs) (workloads.K8sCluster, error) {
	nodeIDs := []uint32{}
	for node := range nodeContracts {
		nodeIDs = append(nodeIDs, node)
	}

	return c.GridClient.LoadK8s(masterName, nodeIDs)
}

func (c *Client) loadGWFQDN(ctx context.Context, modelName string) (workloads.GatewayFQDNProxy, error) {
	modelContracts, err := c.loadModelContracts(ctx, modelName)
	if err != nil {
		return workloads.GatewayFQDNProxy{}, errors.Wrapf(err, "failed to load gateway %s contracts", modelName)
	}

	if len(modelContracts.nodeContracts) != 1 {
		return workloads.GatewayFQDNProxy{}, fmt.Errorf("node contracts for gateway %s should be 1, but %d were found", modelName, len(modelContracts.nodeContracts))
	}

	var nodeID uint32
	for node := range modelContracts.nodeContracts {
		// there is only one node contract, so this loop will have only one iteration
		nodeID = node
	}

	return c.GridClient.LoadGatewayFQDN(modelName, nodeID)
}

func (c *Client) loadGWName(ctx context.Context, modelName string) (workloads.GatewayNameProxy, error) {
	modelcontracts, err := c.loadModelContracts(ctx, modelName)
	if err != nil {
		return workloads.GatewayNameProxy{}, errors.Wrapf(err, "failed to get gateway %s contracts", modelName)
	}

	if len(modelcontracts.nodeContracts) != 1 {
		return workloads.GatewayNameProxy{}, fmt.Errorf("node contracts for gateway %s should be 1, but %d were found", modelName, len(modelcontracts.nodeContracts))
	}

	if len(modelcontracts.nameContracts) != 1 {
		return workloads.GatewayNameProxy{}, fmt.Errorf("name contracts for gateway %s should be 1, but %d were found", modelName, len(modelcontracts.nameContracts))
	}

	var nodeID uint32
	for node := range modelcontracts.nodeContracts {
		nodeID = node
	}

	return c.GridClient.LoadGatewayName(modelName, nodeID)
}

func (c *Client) loadDeployment(modelName string, nodeID uint32) (workloads.Deployment, error) {
	return c.GridClient.LoadDeployment(modelName, nodeID)
}

func (c *Client) loadZDB(ctx context.Context, modelName string) (workloads.ZDB, uint32, error) {
	modelcontracts, err := c.loadModelContracts(ctx, modelName)
	if err != nil {
		return workloads.ZDB{}, 0, errors.Wrapf(err, "failed to get zdb %s contract", modelName)
	}

	if len(modelcontracts.nodeContracts) != 1 {
		return workloads.ZDB{}, 0, fmt.Errorf("node contracts for zdb %s should be 1, but %d were found", modelName, len(modelcontracts.nodeContracts))
	}

	var nodeID uint32
	for node := range modelcontracts.nodeContracts {
		nodeID = node
	}

	zdb, err := c.GridClient.LoadZDB(modelName, nodeID)
	if err != nil {
		return workloads.ZDB{}, 0, err
	}

	return zdb, nodeID, nil
}

func (c *Client) loadModelContracts(ctx context.Context, modelName string) (ProjectState, error) {
	projectName := generateProjectName(modelName)

	if projectState, ok := c.Projects[projectName]; ok {
		c.GridClient.SetContractState(projectState.nodeContracts)
		return projectState, nil
	}

	newState := ProjectState{
		nodeContracts: make(map[uint32]state.ContractIDs),
		nameContracts: make(map[uint32]uint64),
	}

	projectContracts, err := c.GridClient.GetProjectContracts(ctx, projectName)
	if err != nil {
		return ProjectState{}, errors.Wrapf(err, "failed to retreive contracts with project name %s", projectName)
	}

	for _, c := range projectContracts.NodeContracts {
		contractID, err := strconv.ParseUint(c.ContractID, 10, 64)
		if err != nil {
			return ProjectState{}, err
		}

		newState.nodeContracts[c.NodeID] = append(newState.nodeContracts[c.NodeID], contractID)
	}

	for _, c := range projectContracts.NameContracts {
		contractID, err := strconv.ParseUint(c.ContractID, 10, 64)
		if err != nil {
			return ProjectState{}, err
		}

		newState.nameContracts[c.NodeID] = contractID
	}

	c.GridClient.SetContractState(newState.nodeContracts)

	c.Projects[projectName] = newState

	return newState, nil
}

func (c *Client) loadGridMachinesModel(ctx context.Context, modelName string) (gridMachinesModel, error) {
	modelContracts, err := c.loadModelContracts(ctx, modelName)
	if err != nil {
		return gridMachinesModel{}, errors.Wrapf(err, "failed to get machines model %s contracts", modelName)
	}

	if len(modelContracts.nodeContracts) == 0 {
		delete(c.Projects, generateProjectName(modelName))
		return gridMachinesModel{}, fmt.Errorf("found 0 contracts for model %s", modelName)
	}

	znet, err := c.loadNetwork(modelName)
	if err != nil {
		return gridMachinesModel{}, err
	}

	deployments := map[uint32]*workloads.Deployment{}
	for nodeID := range modelContracts.nodeContracts {
		dl, err := c.loadDeployment(modelName, nodeID)
		if err != nil {
			// no deployment found on this node, we should continue looking for other deployments
			log.Debug().Msgf("no deployment found on node %d", nodeID)
			continue
		}

		deployments[nodeID] = &dl
	}

	g := gridMachinesModel{
		modelName:   modelName,
		network:     &znet,
		deployments: deployments,
	}

	return g, nil
}
