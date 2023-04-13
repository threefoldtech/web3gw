package tfchain

import (
	"context"
	"errors"

	substrate "github.com/threefoldtech/tfchain/clients/tfchain-client-go"
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

	Transfer struct {
		amount      uint64
		destination string
	}

	CreateTwin struct {
		relay string
		pk    []byte
	}

	AcceptTermsAndConditions struct {
		link string
		hash string
	}

	GetContractWithHash struct {
		nodeID uint32
		hash   substrate.HexHash
	}

	CreateNodeContract struct {
		nodeID             uint32
		body               string
		hash               string
		publicIPs          uint32
		solutionProviderID *uint64
	}

	CreateRentContract struct {
		nodeID             uint32
		solutionProviderID *uint64
	}

	ServiceContractCreate struct {
		service  substrate.AccountID
		consumer substrate.AccountID
	}

	ServiceContractBill struct {
		contractID     uint64
		variableAmount uint64
		metadata       string
	}

	SetServiceContractFees struct {
		contractID  uint64
		baseFee     uint64
		variableFee uint64
	}

	ServiceContractSetMetadata struct {
		contractID uint64
		metadata   string
	}

	CreateFarm struct {
		name      string
		publicIPs []substrate.PublicIPInput
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

func (c *Client) Height(ctx context.Context) (uint32, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.GetCurrentHeight()
}

// Transer an amount of TFT from the loaded account to the destination.
func (c *Client) Transfer(ctx context.Context, args Transfer) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	dest, err := substrate.FromAddress(args.destination)
	if err != nil {
		return err
	}

	return state.client.Transfer(*state.identity, args.amount, dest)
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

func (c *Client) GetTwinByPubKey(ctx context.Context, pk []byte) (uint32, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.GetTwinByPubKey(pk)
}

func (c *Client) CreateTwin(ctx context.Context, args CreateTwin) (uint32, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.CreateTwin(*state.identity, args.relay, args.pk)
}

func (c *Client) AcceptTermsAndConditions(ctx context.Context, args AcceptTermsAndConditions) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.AcceptTermsAndConditions(*state.identity, args.link, args.hash)
}

func (c *Client) GetNode(ctx context.Context, id uint32) (*substrate.Node, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetNode(id)
}

func (c *Client) GetNodes(ctx context.Context, farm_id uint32) ([]uint32, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return []uint32{}, pkg.ErrClientNotConnected{}
	}

	return state.client.GetNodes(farm_id)
}

func (c *Client) CreateNode(ctx context.Context, node *substrate.Node) (uint32, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}
	return state.client.CreateNode(*state.identity, *node)
}

func (c *Client) GetFarm(ctx context.Context, id uint32) (*substrate.Farm, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetFarm(id)
}

func (c *Client) GetFarmByName(ctx context.Context, name string) (uint32, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.GetFarmByName(name)
}

func (c *Client) CreateFarm(ctx context.Context, args CreateFarm) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.CreateFarm(*state.identity, args.name, args.publicIPs)
}

func (c *Client) GetContract(ctx context.Context, contract_id uint64) (*substrate.Contract, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetContract(contract_id)
}

func (c *Client) GetContractIDByNameRegistration(ctx context.Context, name string) (uint64, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.GetContractIDByNameRegistration(name)
}

func (c *Client) GetContractWithHash(ctx context.Context, args GetContractWithHash) (uint64, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.GetContractWithHash(args.nodeID, args.hash)
}

func (c *Client) CreateNameContract(ctx context.Context, name string) (uint64, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.CreateNameContract(*state.identity, name)
}

func (c *Client) CreateNodeContract(ctx context.Context, args CreateNodeContract) (uint64, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.CreateNodeContract(*state.identity, args.nodeID, args.body, args.hash, args.publicIPs, args.solutionProviderID)
}

func (c *Client) CreateRentContract(ctx context.Context, args CreateRentContract) (uint64, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.CreateRentContract(*state.identity, args.nodeID, args.solutionProviderID)
}

func (c *Client) ServiceContractCreate(ctx context.Context, args ServiceContractCreate) (uint64, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.ServiceContractCreate(*state.identity, args.service, args.consumer)
}

func (c *Client) ServiceContractApprove(ctx context.Context, contract_id uint64) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ServiceContractApprove(*state.identity, contract_id)
}

func (c *Client) ServiceContractBill(ctx context.Context, args ServiceContractBill) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ServiceContractBill(*state.identity, args.contractID, args.variableAmount, args.metadata)
}

func (c *Client) ServiceContractCancel(ctx context.Context, contract_id uint64) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ServiceContractCancel(*state.identity, contract_id)
}

func (c *Client) ServiceContractReject(ctx context.Context, contract_id uint64) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ServiceContractReject(*state.identity, contract_id)
}

func (c *Client) ServiceContractSetFees(ctx context.Context, args SetServiceContractFees) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ServiceContractSetFees(*state.identity, args.contractID, args.baseFee, args.variableFee)
}

func (c *Client) ServiceContractSetMetadata(ctx context.Context, args ServiceContractSetMetadata) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ServiceContractSetMetadata(*state.identity, args.contractID, args.metadata)
}

func (c *Client) CancelContract(ctx context.Context, contract_id uint64) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.CancelContract(*state.identity, contract_id)
}

func (c *Client) GetZosVersion(ctx context.Context) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.GetZosVersion()
}
