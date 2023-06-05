package stellar

import (
	"context"

	"github.com/LeeSmet/go-jsonrpc"
	stellargoclient "github.com/threefoldtech/web3_proxy/server/clients/stellar"
	"github.com/threefoldtech/web3_proxy/server/pkg"
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
	}

	StellarState struct {
		Client  *stellargoclient.Client
		network string
	}

	Load struct {
		Network string `json:"network"`
		Secret  string `json:"secret"`
	}

	Transfer struct {
		Amount      string `json:"amount"`
		Destination string `json:"destination"`
		Memo        string `json:"memo"`
	}

	BridgeTransfer struct {
		Amount      string `json:"amount"`
		Destination string `json:"destination"`
	}

	TfchainBridgeTransfer struct {
		Amount string `json:"amount"`
		TwinId uint32 `json:"twin_id"`
	}
)

const (
	// StellarID is the ID for state of a stellar client in the connection state.
	StellarID = "stellar"
)

// Close implements jsonrpc.Closer
func (s *StellarState) Close() {}

// Error implements the error interface
func (e ErrUnknownNetwork) Error() string {
	return "only 'public' and 'testnet' networks are supported"
}

// State from a connection. If no state is present, it is initialized
func State(conState jsonrpc.State) *StellarState {
	raw, exists := conState[StellarID]
	if !exists {
		ns := &StellarState{
			Client:  nil,
			network: stellarNetworkTestnet,
		}
		conState[StellarID] = ns
		return ns
	}
	ns, ok := raw.(*StellarState)
	if !ok {
		// This means the invariant is violated, so panic here is ok
		panic("Invalid saved state for stellar")
	}
	return ns
}

// NewClient creates a new Client ready for use
func NewClient() *Client {
	return &Client{}
}

// Load a client, connecting to the rpc endpoint at the given URL and loading a keypair from the given secret
func (c *Client) Load(ctx context.Context, conState jsonrpc.State, args Load) error {
	if args.Network != stellarNetworkTestnet && args.Network != stellarNetworkPublic {
		return ErrUnknownNetwork{}
	}
	cl, err := stellargoclient.NewClient(args.Secret, args.Network)
	if err != nil {
		return err
	}

	state := State(conState)
	state.Client = cl
	state.network = args.Network

	return nil
}

// Get the public address of the loaded stellar secret
func (c *Client) Address(ctx context.Context, conState jsonrpc.State) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.Address(), nil
}

// Transer an amount of TFT from the loaded account to the destination.
func (c *Client) Transfer(ctx context.Context, conState jsonrpc.State, args Transfer) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.Transfer(args.Destination, args.Memo, args.Amount)
}

// Balance of an account for TFT on stellar.
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

// BridgeToEth transfers TFT from the loaded account to eth bridge and deposits into the destination ethereum account.
func (c *Client) BridgeToEth(ctx context.Context, conState jsonrpc.State, args BridgeTransfer) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.TransferToEthBridge(args.Destination, args.Amount)
}

// Reinstate later

// // BridgeToBsc transfers TFT from the loaded account to bsc bridge and deposits into the destination bsc account.
// func (c *Client) BridgeToBsc(ctx context.Context, conState jsonrpc.State, args BridgeTransfer) error {
// 	state := State(conState)
// 	if state.Client == nil {
// 		return pkg.ErrClientNotConnected{}
// 	}

// 	return state.Client.TransferToBscBridge(args.Destination, args.Amount)
// }

// BridgeToTfchain transfers TFT from the loaded account to tfchain bridge and deposits into a twin account.
func (c *Client) BridgeToTfchain(ctx context.Context, conState jsonrpc.State, args TfchainBridgeTransfer) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.TransferToTfchainBridge(args.Amount, args.TwinId)
}

// Await till a transaction is processed on ethereum bridge that contains a specific memo
func (c *Client) AwaitTransactionOnEthBridge(ctx context.Context, conState jsonrpc.State, memo string) error {
	state := State(conState)
	if state.Client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.Client.AwaitTransactionWithMemoOnEthBridge(memo, 60)
}
