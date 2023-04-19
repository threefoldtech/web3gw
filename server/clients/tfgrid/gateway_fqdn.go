package tfgrid

import (
	"context"
	"fmt"
	"strconv"

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

func (r *Client) GatewayFQDNDeploy(ctx context.Context, gatewayFQDNModel GatewayFQDNModel, projectName string) (GatewayFQDNModel, error) {
	if err := r.validateProjectName(ctx, projectName); err != nil {
		return GatewayFQDNModel{}, err
	}

	gatewayFQDN := workloads.GatewayFQDNProxy{
		NodeID:         gatewayFQDNModel.NodeID,
		Backends:       gatewayFQDNModel.Backends,
		FQDN:           gatewayFQDNModel.FQDN,
		Name:           gatewayFQDNModel.Name,
		TLSPassthrough: gatewayFQDNModel.TLSPassthrough,
		Description:    gatewayFQDNModel.Description,
		SolutionType:   projectName,
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

func (r *Client) GatewayFQDNGet(ctx context.Context, projectName string) (GatewayFQDNModel, error) {
	contracts, err := r.client.GetProjectContracts(ctx, projectName)
	if err != nil {
		return GatewayFQDNModel{}, errors.Wrapf(err, "failed to get project %s contracts", projectName)
	}

	if len(contracts.NodeContracts) != 1 {
		return GatewayFQDNModel{}, fmt.Errorf("node contracts for project %s should be 1, but %d were found", projectName, len(contracts.NodeContracts))
	}

	nodeID := contracts.NodeContracts[0].NodeID

	nodeClient, err := r.client.GetNodeClient(nodeID)
	if err != nil {
		return GatewayFQDNModel{}, errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	nodeContractID, err := strconv.ParseUint(contracts.NodeContracts[0].ContractID, 0, 64)
	if err != nil {
		return GatewayFQDNModel{}, errors.Wrapf(err, "could not parse contract %s into uint64", contracts.NodeContracts[0].ContractID)
	}

	dl, err := nodeClient.DeploymentGet(ctx, nodeContractID)
	if err != nil {
		return GatewayFQDNModel{}, errors.Wrapf(err, "failed to get deployment with contract id %d", nodeContractID)
	}

	if len(dl.Workloads) != 1 {
		return GatewayFQDNModel{}, errors.Wrapf(err, "deployment should include only one gateway workload, but %d were found", len(dl.Workloads))
	}

	wl := &dl.Workloads[0]
	dataI, err := wl.WorkloadData()
	if err != nil {
		return GatewayFQDNModel{}, errors.Wrap(err, "failed to get workload data")
	}

	data, ok := dataI.(*zos.GatewayFQDNProxy)
	if !ok {
		return GatewayFQDNModel{}, fmt.Errorf("could not create gateway fqdn proxy workload from data %v", dataI)
	}

	return GatewayFQDNModel{
		NodeID:         nodeID,
		Name:           wl.Name.String(),
		TLSPassthrough: data.TLSPassthrough,
		Backends:       data.Backends,
		FQDN:           data.FQDN,
		Description:    wl.Description,
		ContractID:     nodeContractID,
	}, nil
}
