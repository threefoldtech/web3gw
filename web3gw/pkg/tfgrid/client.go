package tfgrid

import (
	"context"

	"github.com/LeeSmet/go-jsonrpc"
	tfgridBase "github.com/threefoldtech/3bot/web3gw/server/clients/tfgrid"
	"github.com/threefoldtech/3bot/web3gw/server/pkg"
)

const (
	TFGridID = "tfgrid"
)

type (
	// Client exposing tfgrid methods
	Client struct{}

	tfgridState struct {
		cl *tfgridBase.Client
	}

	Load struct {
		Mnemonic string `json:"mnemonic"`
		Network  string `json:"network"`
	}
)

// NewClient creates a new Client ready for use
func NewClient() *Client {
	return &Client{}
}

// State from a connection. If no state is present, it is initialized
func State(conState jsonrpc.State) *tfgridState {
	raw, exists := conState[TFGridID]
	if !exists {
		ns := &tfgridState{
			cl: nil,
		}
		conState[TFGridID] = ns
		return ns
	}

	ns, ok := raw.(*tfgridState)
	if !ok {
		// This means the invariant is violated, so panic here is ok
		panic("Invalid saved state for tfchain")
	}
	return ns
}

// Close implements jsonrpc.Closer
func (s *tfgridState) Close() {
	s.cl.GridClient.Close()
}

// Load an identity for the tfgrid with the given network
func (c *Client) Load(ctx context.Context, conState jsonrpc.State, args Load) error {
	state := State(conState)
	if state.cl != nil {
		state.Close()
	}

	tfgrid_client := tfgridBase.Client{
		Projects: make(map[string]tfgridBase.ProjectState),
	}

	err := tfgrid_client.Login(ctx, tfgridBase.Credentials{
		Mnemonics: args.Mnemonic,
		Network:   args.Network,
	})
	if err != nil {
		return err
	}

	state.cl = &tfgrid_client

	return nil
}

func (c *Client) DeployVM(ctx context.Context, conState jsonrpc.State, args tfgridBase.DeployVM) (tfgridBase.VMDeployment, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.VMDeployment{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.DeployVM(ctx, args)
}

func (c *Client) GetVMDeployment(ctx context.Context, conState jsonrpc.State, networkName string) (tfgridBase.VMDeployment, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.VMDeployment{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.GetVMDeployment(ctx, networkName)
}

func (c *Client) CancelVMDeployment(ctx context.Context, conState jsonrpc.State, name string) error {
	state := State(conState)
	if state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.cl.CancelVMDeployment(ctx, name)
}

func (c *Client) DeployNetwork(ctx context.Context, conState jsonrpc.State, args tfgridBase.NetworkDeployment) (tfgridBase.NetworkDeployment, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.NetworkDeployment{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.DeployNetwork(ctx, args)
}

func (c *Client) GetNetworkDeployment(ctx context.Context, conState jsonrpc.State, name string) (tfgridBase.NetworkDeployment, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.NetworkDeployment{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.GetNetworkDeployment(ctx, name)
}

func (c *Client) CancelNetworkDeployment(ctx context.Context, conState jsonrpc.State, name string) error {
	state := State(conState)
	if state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.cl.CancelNetworkDeployment(ctx, name)
}

func (c *Client) AddVMToNetworkDeployment(ctx context.Context, conState jsonrpc.State, args tfgridBase.AddVMToNetworkDeployment) (tfgridBase.NetworkDeployment, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.NetworkDeployment{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.AddVMToNetworkDeployment(ctx, args)
}

func (c *Client) RemoveVMFromNetworkDeployment(ctx context.Context, conState jsonrpc.State, removeMachine tfgridBase.RemoveVMFromNetworkDeployment) (tfgridBase.NetworkDeployment, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.NetworkDeployment{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.RemoveVMFromNetworkDeployment(ctx, removeMachine)
}

func (c *Client) K8sDeploy(ctx context.Context, conState jsonrpc.State, model tfgridBase.K8sCluster) (tfgridBase.K8sCluster, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.K8sCluster{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.K8sDeploy(ctx, model)
}

func (c *Client) K8sGet(ctx context.Context, conState jsonrpc.State, k8sGetInfo tfgridBase.GetClusterParams) (tfgridBase.K8sCluster, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.K8sCluster{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.K8sGet(ctx, k8sGetInfo)
}

func (c *Client) K8sDelete(ctx context.Context, conState jsonrpc.State, modelName string) error {
	state := State(conState)
	if state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.cl.K8sDelete(ctx, modelName)
}

func (c *Client) AddK8sWorker(ctx context.Context, conState jsonrpc.State, workerInfo tfgridBase.AddWorkerParams) (tfgridBase.K8sCluster, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.K8sCluster{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.AddK8sWorker(ctx, workerInfo)
}

func (c *Client) RemoveK8sWorker(ctx context.Context, conState jsonrpc.State, removeWorkerInfo tfgridBase.RemoveWorkerParams) (tfgridBase.K8sCluster, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.K8sCluster{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.RemoveK8sWorker(ctx, removeWorkerInfo)
}

func (c *Client) ZDBDeploy(ctx context.Context, conState jsonrpc.State, model tfgridBase.ZDB) (tfgridBase.ZDB, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.ZDB{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.ZDBDeploy(ctx, model)
}

func (c *Client) ZDBGet(ctx context.Context, conState jsonrpc.State, modelName string) (tfgridBase.ZDB, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.ZDB{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.ZDBGet(ctx, modelName)
}

func (c *Client) ZDBDelete(ctx context.Context, conState jsonrpc.State, modelName string) error {
	state := State(conState)
	if state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.cl.ZDBDelete(ctx, modelName)
}

func (c *Client) GatewayNameDeploy(ctx context.Context, conState jsonrpc.State, model tfgridBase.GatewayNameModel) (tfgridBase.GatewayNameModel, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.GatewayNameModel{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.GatewayNameDeploy(ctx, model)
}

func (c *Client) GatewayNameGet(ctx context.Context, conState jsonrpc.State, modelName string) (tfgridBase.GatewayNameModel, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.GatewayNameModel{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.GatewayNameGet(ctx, modelName)
}

func (c *Client) GatewayNameDelete(ctx context.Context, conState jsonrpc.State, modelName string) error {
	state := State(conState)
	if state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.cl.GatewayNameDelete(ctx, modelName)
}

func (c *Client) GatewayFQDNDeploy(ctx context.Context, conState jsonrpc.State, model tfgridBase.GatewayFQDNModel) (tfgridBase.GatewayFQDNModel, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.GatewayFQDNModel{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.GatewayFQDNDeploy(ctx, model)
}

func (c *Client) GatewayFQDNGet(ctx context.Context, conState jsonrpc.State, modelName string) (tfgridBase.GatewayFQDNModel, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.GatewayFQDNModel{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.GatewayFQDNGet(ctx, modelName)
}

func (c *Client) GatewayFQDNDelete(ctx context.Context, conState jsonrpc.State, modelName string) error {
	state := State(conState)
	if state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.cl.GatewayFQDNDelete(ctx, modelName)
}

func (c *Client) FindNodes(ctx context.Context, conState jsonrpc.State, filters tfgridBase.NodeFilterOptions) ([]uint32, error) {
	state := State(conState)
	if state.cl == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.cl.FilterNodes(ctx, filters)
}

func (c *Client) FindFarms(ctx context.Context, conState jsonrpc.State, filters tfgridBase.FarmFilterOptions) ([]uint32, error) {
	state := State(conState)
	if state.cl == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.cl.FilterFarms(ctx, filters)
}

func (c *Client) FindContracts(ctx context.Context, conState jsonrpc.State, filters tfgridBase.ContractFilterOptions) ([]uint32, error) {
	state := State(conState)
	if state.cl == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.cl.FilterContracts(ctx, filters)
}

func (c *Client) FindTwins(ctx context.Context, conState jsonrpc.State, filters tfgridBase.TwinFilterOptions) ([]uint32, error) {
	state := State(conState)
	if state.cl == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.cl.FilterTwins(ctx, filters)
}

func (c *Client) Statistics(ctx context.Context, conState jsonrpc.State, filters tfgridBase.StatsFilterOptions) (tfgridBase.CountersResult, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.CountersResult{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.GetStatistics(ctx, filters)
}
