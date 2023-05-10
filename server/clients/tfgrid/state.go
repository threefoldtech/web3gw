package tfgrid

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/state"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/workloads"
)

type modelState struct {
	nodeContracts map[uint32]state.ContractIDs
	// used in gateway names
	nameContracts map[uint32]uint64
}

func (c *Client) loadNetwork(modelName string, networkContracts map[uint32]uint64) (workloads.ZNet, error) {
	networkState := map[uint32]state.ContractIDs{}
	for nodeID, contractID := range networkContracts {
		networkState[nodeID] = state.ContractIDs{contractID}
	}

	return c.client.LoadNetwork(generateNetworkName(modelName))
}

func (c *Client) loadK8s(masterName string, nodeContracts map[uint32][]uint64) (workloads.K8sCluster, error) {
	k8sState := map[uint32]state.ContractIDs{}
	nodeIDs := []uint32{}
	for node, contracts := range nodeContracts {
		k8sState[node] = contracts
		nodeIDs = append(nodeIDs, node)
	}

	c.client.SetNodeDeploymentState(k8sState)

	return c.client.LoadK8s(masterName, nodeIDs)
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

	return c.client.LoadGatewayFQDN(modelName, nodeID)
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

	return c.client.LoadGatewayName(modelName, nodeID)
}

func (c *Client) loadDeployment(modelName string, nodeID uint32, contractID uint64) (workloads.Deployment, error) {
	return c.client.LoadDeployment(modelName, nodeID)
}

func (c *Client) loadZDB(modelName string, nodeID uint32, contractID uint64) (workloads.ZDB, error) {
	return c.client.LoadZDB(modelName, nodeID)
}

func (c *Client) loadModelContracts(ctx context.Context, modelName string) (modelState, error) {
	projectName := generateProjectName(modelName)
	modState := modelState{}
	projectState, ok := c.projects[projectName]
	if ok {
		modState = projectState
		c.client.SetContractState(modState.nodeContracts)
		return modState, nil
	}

	projectContracts, err := c.client.GetProjectContracts(ctx, projectName)
	if err != nil {
		return modelState{}, errors.Wrapf(err, "failed to retreive contracts with project name %s", projectName)
	}

	nodeContracts := map[uint32]state.ContractIDs{}
	for _, c := range projectContracts.NodeContracts {
		contractID, err := strconv.ParseUint(c.ContractID, 10, 64)
		if err != nil {
			return modelState{}, err
		}

		nodeContracts[c.NodeID] = append(nodeContracts[c.NodeID], contractID)
	}

	c.client.SetContractState(nodeContracts)

	nameContracts := map[uint32]uint64{}
	for _, c := range projectContracts.NameContracts {
		contractID, err := strconv.ParseUint(c.ContractID, 10, 64)
		if err != nil {
			return modelState{}, err
		}

		nameContracts[c.NodeID] = contractID
	}

	return modelState{
		nodeContracts: nodeContracts,
		nameContracts: nameContracts,
	}, nil
}

func (c *Client) loadGridMachinesMadel(modelName string, nodeDeployments map[uint32]uint64, networkDeployments map[uint32]uint64) (gridMachinesModel, error) {
	znet, err := c.loadNetwork(modelName, networkDeployments)
	if err != nil {
		return gridMachinesModel{}, err
	}

	deployments := map[uint32]workloads.Deployment{}
	for nodeID, contractID := range nodeDeployments {
		dl, err := c.loadDeployment(modelName, nodeID, contractID)
		if err != nil {
			return gridMachinesModel{}, errors.Wrap(err, "failed to load deployments")
		}
		deployments[nodeID] = dl
	}

	return gridMachinesModel{
		modelName:   modelName,
		network:     &znet,
		deployments: deployments,
	}, nil
}

func (c *Client) setNetworkState(g *gridMachinesModel) error {
	subnets := map[uint32]string{}
	for nodeID, subnet := range g.network.NodesIPRange {
		subnets[nodeID] = subnet.String()
	}

	usedIPs := state.NodeDeploymentHostIDs{}
	for nodeID, dl := range g.deployments {
		nodeUsedIPs := state.DeploymentHostIDs{}
		for _, vm := range dl.Vms {
			slices := strings.SplitAfter(vm.IP, ".")
			hostID := slices[len(slices)-1]
			id, err := strconv.ParseUint(hostID, 10, 8)
			if err != nil {
				return err
			}
			nodeUsedIPs[dl.ContractID] = append(nodeUsedIPs[dl.ContractID], byte(id))
		}
		usedIPs[nodeID] = nodeUsedIPs
	}

	c.client.SetNetworkState(generateNetworkName(g.modelName), subnets, usedIPs)

	return nil
}
