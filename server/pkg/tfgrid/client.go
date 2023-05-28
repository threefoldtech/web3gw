package tfgrid

import (
	"context"

	"github.com/LeeSmet/go-jsonrpc"
	tfgridBase "github.com/threefoldtech/web3_proxy/server/clients/tfgrid"
	"github.com/threefoldtech/web3_proxy/server/pkg"
	"github.com/threefoldtech/web3_proxy/server/pkg/state"
)

const (
	// keyType for the TF grid
	keyType = "sr25519"

	// NetworkMain is the TF grid mainnet
	NetworkMain = "main"
	// NetworkTest is the TF grid testnet
	NetworkTest = "test"
	// NetworkQa is the TF grid qanet
	NetworkQA = "qa"
	// NetworkDev is the TF grid devnet
	NetworkDev = "dev"

	// DeployerTimeoutSeconds is the amount of seconds before deployment operations time out
	DeployerTimeoutSeconds = 600 // 10 minutes

	TFGridID = "tfgrid"
)

type (
	// Client exposing tfgrid methods
	Client struct {
		state *state.StateManager[*tfgridState]
	}

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
	return &Client{
		state: state.NewStateManager[*tfgridState](),
	}
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

func (c *Client) MachinesDeploy(ctx context.Context, conState jsonrpc.State, model tfgridBase.MachinesModel) (tfgridBase.MachinesModel, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.MachinesModel{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.MachinesDeploy(ctx, model)
}

func (c *Client) MachinesGet(ctx context.Context, conState jsonrpc.State, modelName string) (tfgridBase.MachinesModel, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.MachinesModel{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.MachinesGet(ctx, modelName)
}

func (c *Client) MachinesDelete(ctx context.Context, conState jsonrpc.State, modelName string) error {
	state := State(conState)
	if state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.cl.MachinesDelete(ctx, modelName)
}

func (c *Client) MachinesAdd(ctx context.Context, conState jsonrpc.State, machine tfgridBase.AddMachineParams) (tfgridBase.MachinesModel, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.MachinesModel{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.MachineAdd(ctx, machine)
}

func (c *Client) MachinesRemove(ctx context.Context, conState jsonrpc.State, removeMachine tfgridBase.RemoveMachineParams) (tfgridBase.MachinesModel, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.MachinesModel{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.MachineRemove(ctx, removeMachine)
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

func (c *Client) FilterNodes(ctx context.Context, conState jsonrpc.State, filters tfgridBase.FilterOptions) ([]uint32, error) {
	state := State(conState)
	if state.cl == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.cl.FilterNodes(ctx, filters)
}

func (c *Client) Logout(ctx context.Context) error {
	st, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || st.cl == nil {
		return pkg.ErrClientNotConnected{}
	}

	st.cl.GridClient.Close()

	return nil
}
