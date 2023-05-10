package tfgrid

import (
	"context"

	"github.com/pkg/errors"
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

func (r *Client) GatewayFQDNDeploy(ctx context.Context, gatewayFQDNModel GatewayFQDNModel) (GatewayFQDNModel, error) {
	if err := r.validateProjectName(ctx, gatewayFQDNModel.Name); err != nil {
		return GatewayFQDNModel{}, err
	}

	gatewayFQDN := workloads.GatewayFQDNProxy{
		NodeID:         gatewayFQDNModel.NodeID,
		Backends:       gatewayFQDNModel.Backends,
		FQDN:           gatewayFQDNModel.FQDN,
		Name:           gatewayFQDNModel.Name,
		TLSPassthrough: gatewayFQDNModel.TLSPassthrough,
		Description:    gatewayFQDNModel.Description,
		SolutionType:   generateProjectName(gatewayFQDNModel.Name),
	}

	if err := r.client.DeployGWFQDN(ctx, &gatewayFQDN); err != nil {
		return GatewayFQDNModel{}, errors.Wrapf(err, "failed to deploy gateway fqdn")
	}

	gatewayFQDNModel.ContractID = gatewayFQDN.ContractID

	return gatewayFQDNModel, nil
}

func (r *Client) GatewayFQDNDelete(ctx context.Context, projectName string) error {
	if err := r.client.CancelProject(ctx, projectName); err != nil {
		return errors.Wrapf(err, "failed to delete gateway fqdn model contracts")
	}

	return nil
}

func (c *Client) GatewayFQDNGet(ctx context.Context, modelName string) (GatewayFQDNModel, error) {
	// check if state is not present, get from graphql
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
