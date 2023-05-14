package tfgrid

import (
	"context"
	"math/rand"

	"github.com/pkg/errors"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/state"
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

func (c *Client) GatewayNameDeploy(ctx context.Context, gw GatewayNameModel) (GatewayNameModel, error) {
	// validate that no other project is deployed with this name
	if err := c.validateProjectName(ctx, gw.Name); err != nil {
		return GatewayNameModel{}, err
	}

	if err := c.ensureGatewayNodeIDExist(&gw); err != nil {
		return GatewayNameModel{}, err
	}

	// deploy gridGW
	gridGW := toGridGWName(gw)

	if err := c.deployGWName(ctx, &gridGW); err != nil {
		return GatewayNameModel{}, err
	}

	return c.GatewayNameGet(ctx, gw.Name)
}

func (c *Client) deployGWName(ctx context.Context, gridGW *workloads.GatewayNameProxy) error {
	if err := c.client.DeployGWName(ctx, gridGW); err != nil {
		return errors.Wrapf(err, "failed to deploy gateway %s", gridGW.Name)
	}

	projectName := generateProjectName(gridGW.Name)

	projectState := map[uint32]state.ContractIDs{
		gridGW.NodeID: {gridGW.ContractID},
	}

	c.Projects[projectName] = ProjectState{
		nodeContracts: projectState,
		nameContracts: map[uint32]uint64{gridGW.NodeID: gridGW.NameContractID},
	}

	return nil
}

func toGridGWName(model GatewayNameModel) workloads.GatewayNameProxy {
	return workloads.GatewayNameProxy{
		NodeID:         model.NodeID,
		Name:           model.Name,
		Backends:       model.Backends,
		TLSPassthrough: model.TLSPassthrough,
		Description:    model.Description,
		SolutionType:   generateProjectName(model.Name),
	}
}

func (c *Client) GatewayNameDelete(ctx context.Context, modelName string) error {
	if err := c.cancelModel(ctx, modelName); err != nil {
		return errors.Wrapf(err, "failed to cancel gateway %s", modelName)
	}

	return nil
}

func (c *Client) GatewayNameGet(ctx context.Context, modelName string) (GatewayNameModel, error) {
	gw, err := c.loadGWName(ctx, modelName)
	if err != nil {
		return GatewayNameModel{}, err
	}

	ret := fromGridGWName(gw)

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

func fromGridGWName(gw workloads.GatewayNameProxy) GatewayNameModel {
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
