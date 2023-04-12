package tfgrid

import (
	"context"

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
)

type (
	// Client exposing tfgrid methods
	Client struct {
		state *state.StateManager[tfgridState]
	}

	tfgridState struct {
		cl *tfgridBase.Runner
	}

	MachinesDeploy struct {
		Model       tfgridBase.MachinesModel `json:"model"`
		ProjectName string                   `json:"project_name"`
	}

	MachinesGet struct {
		ModelName   string `json:"model_name"`
		ProjectName string `json:"project_name"`
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
	tfgrid_client := tfgridBase.Runner{}
	err := tfgrid_client.Login(ctx, tfgridBase.Credentials{
		Mnemonics: mnemonic,
		Network:   network,
	})
	if err != nil {
		return err
	}
	gs := tfgridState{
		cl: &tfgrid_client,
	}

	c.state.Set(state.IDFromContext(ctx), gs)

	return nil
}

func (c *Client) MachinesDeploy(ctx context.Context, args MachinesDeploy) (tfgridBase.MachinesModel, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.cl == nil {
		return tfgridBase.MachinesModel{}, pkg.ErrClientNotConnected{}
	}
	return state.cl.MachinesDeploy(ctx, args.Model, args.ProjectName)
}

func (c *Client) MachinesGet(ctx context.Context, args MachinesGet) (tfgridBase.MachinesModel, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.cl == nil {
		return tfgridBase.MachinesModel{}, pkg.ErrClientNotConnected{}
	}
	return state.cl.MachinesGet(ctx, args.ModelName, args.ProjectName)
}

func (c *Client) MachinesDelete(ctx context.Context, name string) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}
	return state.cl.MachinesDelete(ctx, name)
}
