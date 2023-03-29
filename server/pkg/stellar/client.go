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
	// ErrClientNotConnected indicates stellar client is not yet connected to an ethereum node and or the client does not have a private key loaded yet.
	ErrClientNotConnected struct{}
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

// Error implements Error interface
func (e ErrClientNotConnected) Error() string {
	return "client not connected yet"
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

// Transer an amount of Eth from the loaded account to the destination. The transaction ID is returned.
func (c *Client) Transfer(ctx context.Context, amount string, destination string, memo string) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return ErrClientNotConnected{}
	}

	return state.client.Transfer(destination, memo, amount)
}
