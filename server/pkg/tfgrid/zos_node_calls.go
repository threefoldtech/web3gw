package tfgrid

import (
	"context"
	"encoding/json"
	"net"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/pkg/errors"
	client "github.com/threefoldtech/tfgrid-sdk-go/grid-client/node"
	"github.com/threefoldtech/web3_proxy/server/clients/tfgrid"
	"github.com/threefoldtech/web3_proxy/server/pkg"
	"github.com/threefoldtech/zos/pkg/capacity/dmi"
	"github.com/threefoldtech/zos/pkg/gridtypes"
)

type ZOSNodeRequest struct {
	NodeID uint32          `json:"node_id"`
	Data   json.RawMessage `json:"data"`
}

func (c *Client) ZOSDeploymentDeploy(ctx context.Context, conState jsonrpc.State, request ZOSNodeRequest) error {
	state := State(conState)
	if state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}

	dl := gridtypes.Deployment{}
	if err := json.Unmarshal(request.Data, &dl); err != nil {
		return errors.Wrap(err, "failed to parse deployment data")
	}

	return state.cl.ZOSDeploymentDeploy(ctx, request.NodeID, dl)
}

func (c *Client) ZOSDeploymentGet(ctx context.Context, conState jsonrpc.State, request ZOSNodeRequest) (gridtypes.Deployment, error) {
	state := State(conState)
	if state.cl == nil {
		return gridtypes.Deployment{}, pkg.ErrClientNotConnected{}
	}

	contractID := uint64(0)
	if err := json.Unmarshal(request.Data, &contractID); err != nil {
		return gridtypes.Deployment{}, errors.Wrap(err, "failed to parse deployment data")
	}

	return state.cl.ZOSDeploymentGet(ctx, request.NodeID, contractID)
}

func (c *Client) ZOSDeploymentDelete(ctx context.Context, conState jsonrpc.State, request ZOSNodeRequest) error {
	state := State(conState)
	if state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}

	contractID := uint64(0)
	if err := json.Unmarshal(request.Data, &contractID); err != nil {
		return errors.Wrap(err, "failed to parse deployment data")
	}

	return state.cl.ZOSDeploymentDelete(ctx, request.NodeID, contractID)
}

func (c *Client) ZOSDeploymentUpdate(ctx context.Context, conState jsonrpc.State, request ZOSNodeRequest) error {
	state := State(conState)
	if state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}

	dl := gridtypes.Deployment{}
	if err := json.Unmarshal(request.Data, &dl); err != nil {
		return errors.Wrap(err, "failed to parse deployment data")
	}

	return state.cl.ZOSDeploymentUpdate(ctx, request.NodeID, dl)
}

func (c *Client) ZOSDeploymentChanges(ctx context.Context, conState jsonrpc.State, request ZOSNodeRequest) ([]gridtypes.Workload, error) {
	state := State(conState)
	if state.cl == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	contractID := uint64(0)
	if err := json.Unmarshal(request.Data, &contractID); err != nil {
		return nil, errors.Wrap(err, "failed to parse deployment data")
	}

	return state.cl.ZOSDeploymentChanges(ctx, request.NodeID, contractID)
}

func (c *Client) ZOSStatisticsGet(ctx context.Context, conState jsonrpc.State, request ZOSNodeRequest) (tfgrid.Statistics, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgrid.Statistics{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.ZOSStatisticsGet(ctx, request.NodeID)
}

func (c *Client) ZOSNetworkListWGPorts(ctx context.Context, conState jsonrpc.State, request ZOSNodeRequest) ([]uint16, error) {
	state := State(conState)
	if state.cl == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.cl.ZOSNetworkListWGPorts(ctx, request.NodeID)
}

func (c *Client) ZOSNetworkInterfaces(ctx context.Context, conState jsonrpc.State, request ZOSNodeRequest) (map[string][]net.IP, error) {
	state := State(conState)
	if state.cl == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.cl.ZOSNetworkInterfaces(ctx, request.NodeID)
}

func (c *Client) ZOSNetworkPublicConfigGet(ctx context.Context, conState jsonrpc.State, request ZOSNodeRequest) (client.PublicConfig, error) {
	state := State(conState)
	if state.cl == nil {
		return client.PublicConfig{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.ZOSNetworkPublicConfigGet(ctx, request.NodeID)
}

func (c *Client) ZOSSystemDMI(ctx context.Context, conState jsonrpc.State, request ZOSNodeRequest) (dmi.DMI, error) {
	state := State(conState)
	if state.cl == nil {
		return dmi.DMI{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.ZOSSystemDMI(ctx, request.NodeID)
}

func (c *Client) ZOSSystemHypervisor(ctx context.Context, conState jsonrpc.State, request ZOSNodeRequest) (string, error) {
	state := State(conState)
	if state.cl == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.cl.ZOSSystemHypervisor(ctx, request.NodeID)
}

func (c *Client) ZOSSystemVersion(ctx context.Context, conState jsonrpc.State, request ZOSNodeRequest) (client.Version, error) {
	state := State(conState)
	if state.cl == nil {
		return client.Version{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.ZOSVersion(ctx, request.NodeID)
}

func (c *Client) ZOSStoragePools(ctx context.Context, conState jsonrpc.State, request ZOSNodeRequest) ([]client.PoolMetrics, error) {
	state := State(conState)
	if state.cl == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.cl.ZOSStoragePools(ctx, request.NodeID)
}

func (c *Client) ZOSHasPublicIPv6(ctx context.Context, conState jsonrpc.State, request ZOSNodeRequest) (bool, error) {
	state := State(conState)
	if state.cl == nil {
		return false, pkg.ErrClientNotConnected{}
	}

	return state.cl.ZosHasPublicIPv6(ctx, request.NodeID)
}

func (c *Client) ZOSNetworkListAllInterfaces(ctx context.Context, conState jsonrpc.State, request ZOSNodeRequest) (map[string]client.Interface, error) {
	state := State(conState)
	if state.cl == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.cl.ZOSNetworkListAllInterfaces(ctx, request.NodeID)
}

func (c *Client) ZOSNetworkSetPublicExitDevice(ctx context.Context, conState jsonrpc.State, request ZOSNodeRequest) error {
	state := State(conState)
	if state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}

	iface := ""
	if err := json.Unmarshal(request.Data, &iface); err != nil {
		return errors.Wrap(err, "failed to parse deployment data")
	}

	return state.cl.ZOSNetworkSetPublicExitDevice(ctx, request.NodeID, iface)
}

func (c *Client) ZOSNetworkGetPublicExitDevice(ctx context.Context, conState jsonrpc.State, request ZOSNodeRequest) (client.ExitDevice, error) {
	state := State(conState)
	if state.cl == nil {
		return client.ExitDevice{}, pkg.ErrClientNotConnected{}
	}

	iface := ""
	if err := json.Unmarshal(request.Data, &iface); err != nil {
		return client.ExitDevice{}, errors.Wrap(err, "failed to parse deployment data")
	}

	return state.cl.ZOSNetworkGetPublicExitDevice(ctx, request.NodeID)
}
