package tfgrid

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/pkg/errors"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/workloads"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"

	proxyTypes "github.com/threefoldtech/tfgrid-sdk-go/grid-proxy/pkg/types"
)

// GatewayNameModel struct for gateway name proxy
type GatewayNameModel struct {
	// Required
	NodeID uint32 `json:"node_id"`
	// Name the fully qualified domain name to use (cannot be present with Name)
	Name string `json:"name"`
	// Backends are list of backend ips
	Backends []zos.Backend `json:"backends"`

	// Optional
	// Passthrough whether to pass tls traffic or not
	TLSPassthrough bool   `json:"tls_passthrough"`
	Description    string `json:"description"`

	// computed

	// FQDN deployed on the node
	FQDN           string `json:"fqdn"`
	NameContractID uint64 `json:"name_contract_id"`
	ContractID     uint64 `json:"contract_id"`
}

func (r *Client) GatewayNameDeploy(ctx context.Context, gatewayNameModel GatewayNameModel) (GatewayNameModel, error) {
	projectName := generateProjectName(gatewayNameModel.Name)

	// validate that no other project is deployed with this name
	if err := r.validateProjectName(ctx, projectName); err != nil {
		return GatewayNameModel{}, err
	}

	if err := r.ensureGatewayNodeIDExist(&gatewayNameModel); err != nil {
		return GatewayNameModel{}, err
	}

	// deploy gateway
	gateway := newGWNameProxyFromModel(gatewayNameModel)

	if err := r.client.DeployGWName(ctx, &gateway); err != nil {
		return GatewayNameModel{}, errors.Wrapf(err, "failed to deploy gateway %s", gateway.Name)
	}

	nodeDomain, err := r.client.GetNodeDomain(ctx, gateway.NodeID)
	if err != nil {
		return GatewayNameModel{}, errors.Wrapf(err, "failed to get node %d domain", gateway.NodeID)
	}

	gatewayNameModel.FQDN = fmt.Sprintf("%s.%s", gateway.Name, nodeDomain)
	gatewayNameModel.ContractID = gateway.ContractID
	gatewayNameModel.NameContractID = gateway.NameContractID

	return gatewayNameModel, nil
}

func newGWNameProxyFromModel(model GatewayNameModel) workloads.GatewayNameProxy {
	return workloads.GatewayNameProxy{
		NodeID:         model.NodeID,
		Name:           model.Name,
		Backends:       model.Backends,
		TLSPassthrough: model.TLSPassthrough,
		Description:    model.Description,
		SolutionType:   generateProjectName(model.Name),
	}
}

func (r *Client) GatewayNameDelete(ctx context.Context, projectName string) error {
	if err := r.client.CancelProject(ctx, projectName); err != nil {
		return errors.Wrapf(err, "failed to cancel project %s", projectName)
	}

	return nil
}

func (r *Client) GatewayNameGet(ctx context.Context, modelName string) (GatewayNameModel, error) {
	projectName := generateProjectName(modelName)

	contracts, err := r.client.GetProjectContracts(ctx, projectName)
	if err != nil {
		return GatewayNameModel{}, errors.Wrapf(err, "failed to get project %s contracts", projectName)
	}

	if len(contracts.NodeContracts) != 1 {
		return GatewayNameModel{}, fmt.Errorf("node contracts for project %s should be 1, but %d were found", projectName, len(contracts.NodeContracts))
	}

	if len(contracts.NameContracts) != 1 {
		return GatewayNameModel{}, fmt.Errorf("name contracts for project %s should be 1, but %d were found", projectName, len(contracts.NameContracts))
	}

	nodeContractID, err := strconv.ParseUint(contracts.NodeContracts[0].ContractID, 0, 64)
	if err != nil {
		return GatewayNameModel{}, errors.Wrapf(err, "could not parse contract %s into uint64", contracts.NodeContracts[0].ContractID)
	}

	nodeID := contracts.NodeContracts[0].NodeID

	r.client.SetNodeDeploymentState(map[uint32][]uint64{nodeID: {nodeContractID}})

	gw, err := r.client.LoadGatewayName(nodeID, modelName)
	if err != nil {
		return GatewayNameModel{}, err
	}

	ret := GatewayNameToModel(gw)

	return ret, nil
}

func (r *Client) ensureGatewayNodeIDExist(gatewayNameModel *GatewayNameModel) error {
	if gatewayNameModel.NodeID == 0 {
		nodeId, err := r.getGatewayNode()
		if err != nil {
			return errors.Wrapf(err, "Couldn't find a gateway node")
		}

		gatewayNameModel.NodeID = nodeId
	}
	return nil
}

func (r *Client) getGatewayNode() (uint32, error) {
	options := proxyTypes.NodeFilter{
		Status: &Status,
		IPv4:   &TrueVal,
		Domain: &TrueVal,
	}

	nodes, _, err := r.client.FilterNodes(options, proxyTypes.Limit{})
	if err != nil || len(nodes) == 0 {
		return 0, errors.Wrapf(err, "Couldn't find node for the provided filters: %+v", options)
	}

	return uint32(nodes[rand.Intn(len(nodes))].NodeID), nil
}

func GatewayNameToModel(gw workloads.GatewayNameProxy) GatewayNameModel {
	return GatewayNameModel{
		NodeID:         gw.NodeID,
		Name:           gw.Name,
		Backends:       gw.Backends,
		TLSPassthrough: gw.TLSPassthrough,
		Description:    gw.Description,
		FQDN:           gw.FQDN,
		NameContractID: gw.NameContractID,
		ContractID:     gw.ContractID,
	}
}
