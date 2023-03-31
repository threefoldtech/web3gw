package tfgrid

import (
	"context"

	"github.com/threefoldtech/grid3-go/deployer"
	procedure "github.com/threefoldtech/tf-grid-cli/pkg/server/procedures"
	"github.com/threefoldtech/tf-grid-cli/pkg/server/types"
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
)

type (
	// Client exposing tfgrid methods
	Client struct {
		state *state.StateManager[tfgridState]
	}

	tfgridState struct {
		//cl *tfgrid.Client
		cl *deployer.TFPluginClient
	}
)

// NewClient creates a new Client ready for use
func NewClient() *Client {
	return &Client{
		state: state.NewStateManager[tfgridState](),
	}
}

// Load an identity for the tfgrid with the given network
func (c *Client) Load(ctx context.Context, mnemonic string, network string) error {
	cl, err := deployer.NewTFPluginClient(mnemonic, keyType, network, deployer.SubstrateURLs[network], deployer.RelayURLS[network], deployer.RMBProxyURLs[network], DeployerTimeoutSeconds, true, false)
	if err != nil {
		return err
	}

	gs := tfgridState{
		cl: &cl,
	}

	c.state.Set(state.IDFromContext(ctx), gs)

	return nil
}

func (c *Client) MachinesDeploy(ctx context.Context, model types.MachinesModel) (types.MachinesModel, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.cl == nil {
		return types.MachinesModel{}, pkg.ErrClientNotConnected{}
	}
	return procedure.MachinesDeploy(ctx, model, state.cl)
}

func (c *Client) MachinesGet(ctx context.Context, name string) (types.MachinesModel, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.cl == nil {
		return types.MachinesModel{}, pkg.ErrClientNotConnected{}
	}
	return procedure.MachinesGet(ctx, name, state.cl)
}

func (c *Client) MachinesDelete(ctx context.Context, name string) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}
	return procedure.MachinesDelete(ctx, name, state.cl)
}
