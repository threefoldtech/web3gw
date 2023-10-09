package stellar

import (
	"context"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/protocols/horizon"
	stellargoclient "github.com/threefoldtech/3bot/web3gw/server/clients/stellar"
	"github.com/threefoldtech/3bot/web3gw/server/pkg"
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

	Swap struct {
		Amount           string `json:"amount"`
		SourceAsset      string `json:"source_asset"`
		DestinationAsset string `json:"destination_asset"`
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

	Transactions struct {
		Account       string `json:"account"`
		Limit         uint   `json:"limit"`
		IncludeFailed bool   `json:"include_failed"`
		Cursor        string `json:"cursor"`
		Ascending     bool   `json:"ascending"`
	}

	AccountData struct {
		Account string `json:"account"`
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
	state := State(conState)
	if state.Client == nil {
		state := State(conState)
		state.Client = stellargoclient.NewClient(args.Network)
		state.network = args.Network
	}

	return state.Client.Load(args.Secret)
}

func (c *Client) CreateAccount(ctx context.Context, conState jsonrpc.State, network string) (string, error) {
	if network != stellarNetworkTestnet && network != stellarNetworkPublic {
		return "", ErrUnknownNetwork{}
	}
	state := State(conState)
	if state.Client == nil {
		state := State(conState)
		state.Client = stellargoclient.NewClient(network)
		state.network = network
	}

	return state.Client.CreateAccount()
}

// Get the public address of the loaded stellar secret
func (c *Client) Address(ctx context.Context, conState jsonrpc.State) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.Address(), nil
}

// Swap some amount from one asset to the other (for example from tft to xlm)
func (c *Client) Swap(ctx context.Context, conState jsonrpc.State, args Swap) error {
	state := State(conState)
	if state.Client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.Client.Swap(args.SourceAsset, args.DestinationAsset, args.Amount)
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

	return state.Client.AwaitTransactionWithMemoOnEthBridge(ctx, memo, 300)
}

// Get the last transactions of your account
func (c *Client) Transactions(ctx context.Context, conState jsonrpc.State, args Transactions) ([]horizon.Transaction, error) {
	state := State(conState)
	if state.Client == nil {
		return []horizon.Transaction{}, pkg.ErrClientNotConnected{}
	}
	if args.Account == "" {
		args.Account = state.Client.Address()
	}

	order := horizonclient.OrderDesc
	if args.Ascending {
		order = horizonclient.OrderAsc
	}

	return state.Client.Transactions(args.Account, args.Limit, args.IncludeFailed, args.Cursor, order)
}

// Get data related to a stellar account
func (c *Client) AccountData(ctx context.Context, conState jsonrpc.State, account string) (horizon.Account, error) {
	state := State(conState)
	if state.Client == nil {
		return horizon.Account{}, pkg.ErrClientNotConnected{}
	}

	if account == "" {
		account = state.Client.Address()
	}

	return state.Client.AccountData(account)
}
