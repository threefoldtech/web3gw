package eth

import (
	"context"

	"github.com/LeeSmet/go-jsonrpc"
	goethclient "github.com/threefoldtech/web3_proxy/server/clients/eth"
	"github.com/threefoldtech/web3_proxy/server/pkg"
)

type (
	// Client exposes ethereum related functionality
	Client struct {
	}

	// EthState managed by ethereum client
	EthState struct {
		Client *goethclient.Client
	}

	Load struct {
		Url    string `json:"url"`
		Secret string `json:"secret"`
	}

	Transfer struct {
		Amount      string `json:"amount"`
		Destination string `json:"destination"`
	}
)

const (
	// Eth is the ID for state of an eth client in the connection state.
	EthID = "eth"
)

// NewClient creates a new Client ready for use
func NewClient() *Client {
	return &Client{}
}

// State from a connection. If no state is present, it is initialized
func State(conState jsonrpc.State) *EthState {
	raw, exists := conState[EthID]
	if !exists {
		ns := &EthState{
			Client: nil,
		}
		conState[EthID] = ns
		return ns
	}
	ns, ok := raw.(*EthState)
	if !ok {
		// This means the invariant is violated, so panic here is ok
		panic("Invalid saved state for atomic swap")
	}
	return ns
}

// Close implements jsonrpc.Closer
func (s *EthState) Close() {}

// Load a client, connecting to the rpc endpoint at the given URL and loading a keypair from the given secret
func (c *Client) Load(ctx context.Context, conState jsonrpc.State, args Load) error {
	cl, err := goethclient.NewClient(args.Url, args.Secret)
	if err != nil {
		return err
	}

	state := State(conState)

	state.Client = cl

	return nil
}

// Balance of an address
func (c *Client) Balance(ctx context.Context, conState jsonrpc.State, address string) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	balance, err := state.Client.GetBalance(address)
	if err != nil {
		return "", err
	}

	return balance, nil
}

// Height of the chain for the connected rpc remote
func (c *Client) Height(ctx context.Context, conState jsonrpc.State) (uint64, error) {
	state := State(conState)
	if state.Client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.Client.GetCurrentHeight()
}

// Transer an amount of Eth from the loaded account to the destination. The transaction ID is returned.
func (c *Client) Transfer(ctx context.Context, conState jsonrpc.State, args Transfer) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.TransferEth(ctx, args.Amount, args.Destination)
}

// Address of the loaded client
func (c *Client) Address(ctx context.Context, conState jsonrpc.State) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.AddressFromKey().String(), nil
}

func (c *Client) GetHexSeed(ctx context.Context, conState jsonrpc.State) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.GetHexSeed(), nil
}
