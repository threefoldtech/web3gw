package eth

import (
	"context"

	goethclient "github.com/threefoldtech/web3_proxy/server/clients/eth"
	"github.com/threefoldtech/web3_proxy/server/pkg"
	"github.com/threefoldtech/web3_proxy/server/pkg/state"
)

type (
	// Client exposes ethereum related functionality
	Client struct {
		state *state.StateManager[ethState]
	}
	// state managed by ethereum client
	ethState struct {
		client *goethclient.Client
	}

	Transfer struct {
		Amount      int64  `json:"amount"`
		Destination string `json:"destination"`
	}
)

// NewClient creates a new Client ready for use
func NewClient() *Client {
	return &Client{
		state: state.NewStateManager[ethState](),
	}
}

// Load a client, connecting to the rpc endpoint at the given URL and loading a keypair from the given secret
func (c *Client) Load(ctx context.Context, url string, secret string) error {
	cl, err := goethclient.NewClient(url, secret)
	if err != nil {
		return err
	}

	es := ethState{
		client: cl,
	}

	c.state.Set(state.IDFromContext(ctx), es)

	return nil
}

// Balance of an address
func (c *Client) Balance(ctx context.Context, address string) (int64, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	balance, err := state.client.GetBalance(address)
	if err != nil {
		return 0, err
	}

	return balance.Int64(), nil
}

// Height of the chain for the connected rpc remote
func (c *Client) Height(ctx context.Context) (uint64, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.GetCurrentHeight()
}

// Transer an amount of Eth from the loaded account to the destination. The transaction ID is returned.
func (c *Client) Transfer(ctx context.Context, args Transfer) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.TransferEth(args.Amount, args.Destination)
}
