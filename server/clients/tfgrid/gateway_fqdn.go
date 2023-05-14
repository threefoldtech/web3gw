package tfgrid

import (
	"context"

	"github.com/pkg/errors"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/state"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/workloads"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"
)

// GatewayFQDNModel for gateway FQDN proxy
type GatewayFQDNModel struct {
	// required
	NodeID uint32 `json:"node_id"`
	// Backends are list of backend ips
	Backends []zos.Backend `json:"backends"`
	// FQDN deployed on the node
	FQDN string `json:"fqdn"`
	// Name is the workload name
	Name string `json:"name"`

	// optional
	// Passthrough whether to pass tls traffic or not
	TLSPassthrough bool   `json:"tls_passthrough"`
	Description    string `json:"description"`

	// computed
	ContractID uint64 `json:"contract_id"`
}

func (c *Client) GatewayFQDNDeploy(ctx context.Context, gw GatewayFQDNModel) (GatewayFQDNModel, error) {
	if err := c.validateProjectName(ctx, gw.Name); err != nil {
		return GatewayFQDNModel{}, err
	}

	gridGW := workloads.GatewayFQDNProxy{
		NodeID:         gw.NodeID,
		Backends:       gw.Backends,
		FQDN:           gw.FQDN,
		Name:           gw.Name,
		TLSPassthrough: gw.TLSPassthrough,
		Description:    gw.Description,
		SolutionType:   generateProjectName(gw.Name),
	}

	if err := c.deployGWFQDN(ctx, &gridGW); err != nil {
		return GatewayFQDNModel{}, errors.Wrapf(err, "failed to deploy gateway %s", gw.Name)
	}

	gw.ContractID = gridGW.ContractID

	return gw, nil
}

func (c *Client) deployGWFQDN(ctx context.Context, gridGW *workloads.GatewayFQDNProxy) error {
	if err := c.client.DeployGWFQDN(ctx, gridGW); err != nil {
		return err
	}

	projectName := generateProjectName(gridGW.Name)

	projectState := map[uint32]state.ContractIDs{
		gridGW.NodeID: {gridGW.ContractID},
	}

	c.Projects[projectName] = ProjectState{
		nodeContracts: projectState,
	}

	return nil
}

func (c *Client) GatewayFQDNDelete(ctx context.Context, modelName string) error {
	if err := c.client.CancelProject(ctx, modelName); err != nil {
		return errors.Wrapf(err, "failed to delete gateway fqdn model contracts")
	}

	return nil
}

func (c *Client) GatewayFQDNGet(ctx context.Context, modelName string) (GatewayFQDNModel, error) {
	gw, err := c.loadGWFQDN(ctx, modelName)
	if err != nil {
		return GatewayFQDNModel{}, err
	}

	ret := GatewayFQDNToModel(gw)

	return ret, nil
}

func GatewayFQDNToModel(gw workloads.GatewayFQDNProxy) GatewayFQDNModel {
	return GatewayFQDNModel{
		NodeID:         gw.NodeID,
		Backends:       gw.Backends,
		FQDN:           gw.FQDN,
		Name:           gw.Name,
		TLSPassthrough: gw.TLSPassthrough,
		Description:    gw.Description,
		ContractID:     gw.ContractID,
	}
}
