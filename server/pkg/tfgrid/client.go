package tfgrid

import (
	"context"

	"github.com/threefoldtech/web3_proxy/server/clients/tfgrid"
	"github.com/threefoldtech/web3_proxy/server/pkg/state"
)

type (
	// Client exposing tfgrid methods
	Client struct {
		state *state.StateManager[tfgridState]
	}

	tfgridState struct {
		cl *tfgrid.Client
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
	cl, err := tfgrid.NewClient(mnemonic, network)
	if err != nil {
		return err
	}

	gs := tfgridState{
		cl: cl,
	}

	c.state.Set(state.IDFromContext(ctx), gs)

	return nil
}
