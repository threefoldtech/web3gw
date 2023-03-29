package stellar

import (
	"context"

	stellargoclient "github.com/threefoldtech/web3_proxy/server/clients/stellar"
	"github.com/threefoldtech/web3_proxy/server/pkg/state"
)

const (
	stellarNetworkPublic  = "public"
	stellarNetworkTestnet = "testnet"
)

type (
	// ErrUnknownNetwork indicates a client was requested for an unknown network
	ErrUnknownNetwork struct{}
	// Client exposing stellar methods
	Client struct {
		state *state.StateManager[stellarState]
	}
	stellarState struct {
		client  *stellargoclient.Client
		network string
	}
)

// Error implements the error interface
func (e ErrUnknownNetwork) Error() string {
	return "only 'public' and 'testnet' networks are supported"
}

// NewClient creates a new Client ready for use
func NewClient() *Client {
	return &Client{
		state: state.NewStateManager[stellarState](),
	}
}

// Load a client, connecting to the rpc endpoint at the given URL and loading a keypair from the given secret
func (c *Client) Load(ctx context.Context, network string, secret string) error {
	if network != stellarNetworkTestnet && network != stellarNetworkPublic {
		return ErrUnknownNetwork{}
	}
	cl, err := stellargoclient.NewClient(secret, network)
	if err != nil {
		return err
	}

	ss := stellarState{
		client:  cl,
		network: network,
	}

	c.state.Set(state.IDFromContext(ctx), ss)

	return nil
}
