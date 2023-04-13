package tfgrid

/*
import (
	"context"
	"fmt"

	tfgridBase "github.com/threefoldtech/tfgrid-sdk-go/grid-client/deployer"
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
)

// NewClient creates a new Client ready for use
func NewClient() *Client {
	return &Client{
		state: state.NewStateManager[tfgridState](),
	}
}

func generateProjectName(modelName string) (projectName string) {
	return fmt.Sprintf("%s.web3proxy", modelName)
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
*/
