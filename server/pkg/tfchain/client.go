package tfchain

import (
	"context"
	"errors"

	"github.com/threefoldtech/substrate-client"
	"github.com/threefoldtech/web3_proxy/server/pkg"
	"github.com/threefoldtech/web3_proxy/server/pkg/state"
)

const (
	tfchainMainnet = "wss://tfchain.grid.tf"
	tfchainTestnet = "wss://tfchain.test.grid.tf"
	tfchainQanet   = "wss://tfchain.qa.grid.tf"
	tfchainDevnet  = "wss://tfchain.dev.grid.tf"
)

type (
	// ErrUnknownNetwork indicates a client was requested for an unknown network
	ErrUnknownNetwork struct{}
	// Client exposing stellar methods
	Client struct {
		state *state.StateManager[tfchainState]
	}
	tfchainState struct {
		client   *substrate.Substrate
		identity *substrate.Identity
		network  string
	}
)

// Error implements the error interface
func (e ErrUnknownNetwork) Error() string {
	return "only 'public' and 'testnet' networks are supported"
}

// NewClient creates a new Client ready for use
func NewClient() *Client {
	return &Client{
		state: state.NewStateManager[tfchainState](),
	}
}

// Load a client, connecting to the rpc endpoint at the given URL and loading a keypair from the given mnemonic
func (c *Client) Load(ctx context.Context, network string, passphrase string) error {
	url, err := tfchainNetworkFromNetworkString(network)
	if err != nil {
		return err
	}

	mgr := substrate.NewManager(url)
	substrateConnection, err := mgr.Substrate()
	if err != nil {
		return err
	}

	identity, err := substrate.NewIdentityFromSr25519Phrase(passphrase)
	if err != nil {
		return err
	}

	ts := tfchainState{
		client:   substrateConnection,
		identity: &identity,
		network:  network,
	}

	c.state.Set(state.IDFromContext(ctx), ts)

	return nil
}

// Transer an amount of TFT from the loaded account to the destination.
func (c *Client) Transfer(ctx context.Context, amount uint64, destination string, memo string) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	dest, err := substrate.FromAddress(destination)
	if err != nil {
		return err
	}

	return state.client.Transfer(*state.identity, amount, dest)
}

// Balance of an account for TFT on stellar.
func (c *Client) Balance(ctx context.Context, address string) (int64, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	accountId, err := substrate.FromAddress(address)
	if err != nil {
		return 0, err
	}

	balance, err := state.client.GetBalance(accountId)
	if err != nil {
		return 0, err
	}

	return balance.Free.Int64(), nil
}

func (c *Client) GetTwin(ctx context.Context, id uint32) (*substrate.Twin, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetTwin(id)
}

func (c *Client) GetNode(ctx context.Context, id uint32) (*substrate.Node, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetNode(id)
}

func (c *Client) GetFarm(ctx context.Context, id uint32) (*substrate.Farm, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetFarm(id)
}

func tfchainNetworkFromNetworkString(ntwrk string) (string, error) {
	if ntwrk == "mainnet" {
		return tfchainMainnet, nil
	} else if ntwrk == "testnet" {
		return tfchainTestnet, nil
	} else if ntwrk == "qanet" {
		return tfchainQanet, nil
	} else if ntwrk == "devnet" {
		return tfchainDevnet, nil
	}

	return "", errors.New("unsupported network")
}
