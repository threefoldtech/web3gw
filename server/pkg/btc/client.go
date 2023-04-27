package btc

import (
	"context"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	btcRpcClient "github.com/btcsuite/btcd/rpcclient"
	"github.com/threefoldtech/web3_proxy/server/pkg"
)

const (
	// NostrID is the ID for state of a btc client in the connection state.
	BtcID = "btc"
)

type (
	// Client exposes nostr related functionality
	Client struct {
	}
	// state managed by nostr client
	btcState struct {
		client *btcRpcClient.Client
	}

	Load struct {
		Host string `json:"host"`
		User string `json:"user"`
		Pass string `json:"pass"`
	}

	ImportAddressRescan struct {
		Address string `json:"address"`
		Account string `json:"account"`
		Rescan  bool   `json:"rescan"`
	}

	ImportPrivKeyLabel struct {
		WIF   string `json:"wif"`
		Label string `json:"label"`
	}

	ImportPrivKeyRescan struct {
		WIF    string `json:"wif"`
		Label  string `json:"label"`
		Rescan bool   `json:"rescan"`
	}

	ImportPubKeyRescan struct {
		PubKey string `json:"pub_key"`
		Rescan bool   `json:"rescan"`
	}

	RenameAccount struct {
		OldAccount string `json:"old_account"`
		NewAccount string `json:"new_account"`
	}

	SubmitBlock struct {
		Block   *btcutil.Block              `json:"block"`
		Options *btcjson.SubmitBlockOptions `json:"options"`
	}

	SendToAddress struct {
		Address   string         `json:"address"`
		Amount    btcutil.Amount `json:"amount"`
		Comment   string         `json:"comment"`
		CommentTo string         `json:"comment_to"`
	}

	EstimateSmartFee struct {
		ConfTarget int64                        `json:"conf_target"`
		Mode       btcjson.EstimateSmartFeeMode `json:"mode"`
	}

	GenerateToAddress struct {
		NumBlocks int64  `json:"num_blocks"`
		Address   string `json:"address"`
		MaxTries  int64  `json:"max_tries"`
	}
)

// State from a connection. If no state is present, it is initialized
func State(conState jsonrpc.State) *btcState {
	raw, exists := conState[BtcID]
	if !exists {
		ns := &btcState{}
		conState[BtcID] = ns
		return ns
	}
	ns, ok := raw.(*btcState)
	if !ok {
		// This means the invariant is violated, so panic here is ok
		panic("Invalid saved state for btc")
	}
	return ns
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Load(ctx context.Context, conState jsonrpc.State, args Load) error {
	client, err := btcRpcClient.New(
		&btcRpcClient.ConnConfig{
			Host:         args.Host,
			User:         args.User,
			Pass:         args.Pass,
			HTTPPostMode: true,
			DisableTLS:   true,
		}, nil)
	if err != nil {
		return err
	}
	state := State(conState)
	state.client = client

	return nil
}

func (c *Client) CreateNewAccount(ctx context.Context, conState jsonrpc.State, account string) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.CreateNewAccount(account)
}

func (c *Client) CreateEncryptedWallet(ctx context.Context, conState jsonrpc.State, passphrase string) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.CreateEncryptedWallet(passphrase)
}

func (c *Client) ImportAddress(ctx context.Context, conState jsonrpc.State, address string) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ImportAddress(address)
}

func (c *Client) ImportAddressRescan(ctx context.Context, conState jsonrpc.State, args ImportAddressRescan) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ImportAddressRescan(args.Address, args.Account, args.Rescan)
}

func (c *Client) ImportPrivKey(ctx context.Context, conState jsonrpc.State, wif string) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	privKeyWIF, err := btcutil.DecodeWIF(wif)
	if err != nil {
		return err
	}

	return state.client.ImportPrivKey(privKeyWIF)
}

func (c *Client) ImportPrivKeyLabel(ctx context.Context, conState jsonrpc.State, args ImportPrivKeyLabel) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	privKeyWIF, err := btcutil.DecodeWIF(args.WIF)
	if err != nil {
		return err
	}

	return state.client.ImportPrivKeyLabel(privKeyWIF, args.Label)
}

func (c *Client) ImportPrivKeyRescan(ctx context.Context, conState jsonrpc.State, args ImportPrivKeyRescan) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	privKeyWIF, err := btcutil.DecodeWIF(args.WIF)
	if err != nil {
		return err
	}

	return state.client.ImportPrivKeyRescan(privKeyWIF, args.Label, args.Rescan)
}

func (c *Client) ImportPubKey(ctx context.Context, conState jsonrpc.State, pubKey string) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ImportPubKey(pubKey)
}

func (c *Client) ImportPubKeyRescan(ctx context.Context, conState jsonrpc.State, args ImportPubKeyRescan) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ImportPubKeyRescan(args.PubKey, args.Rescan)
}

func (c *Client) InvalidateBlock(ctx context.Context, conState jsonrpc.State, hash string) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	blockHash, err := chainhash.NewHashFromStr(hash)
	if err != nil {
		return err
	}

	return state.client.InvalidateBlock(blockHash)
}

func (c *Client) RenameAccount(ctx context.Context, conState jsonrpc.State, args RenameAccount) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.RenameAccount(args.OldAccount, args.NewAccount)
}

func (c *Client) SubmitBlock(ctx context.Context, conState jsonrpc.State, args SubmitBlock) error {
	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.SubmitBlock(args.Block, args.Options)
}

func (c *Client) SendToAddress(ctx context.Context, conState jsonrpc.State, args SendToAddress) (*chainhash.Hash, error) {
	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	address, err := btcutil.DecodeAddress(args.Address, nil)
	if err != nil {
		return nil, err
	}

	return state.client.SendToAddress(address, args.Amount)
}

func (c *Client) SendToAddressComment(ctx context.Context, conState jsonrpc.State, args SendToAddress) (*chainhash.Hash, error) {
	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	address, err := btcutil.DecodeAddress(args.Address, nil)
	if err != nil {
		return nil, err
	}

	return state.client.SendToAddressComment(address, args.Amount, args.Comment, args.CommentTo)
}

func (c *Client) EstimateFee(ctx context.Context, conState jsonrpc.State, numBlocks int64) (float64, error) {
	state := State(conState)
	if state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.EstimateFee(numBlocks)
}

func (c *Client) EstimateSmartFee(ctx context.Context, conState jsonrpc.State, args EstimateSmartFee) (*btcjson.EstimateSmartFeeResult, error) {
	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.EstimateSmartFee(args.ConfTarget, &args.Mode)
}

func (c *Client) Generate(ctx context.Context, conState jsonrpc.State, numBlocks uint32) ([]*chainhash.Hash, error) {
	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.Generate(numBlocks)
}

func (c *Client) GenerateToAddress(ctx context.Context, conState jsonrpc.State, args GenerateToAddress) ([]*chainhash.Hash, error) {
	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	address, err := btcutil.DecodeAddress(args.Address, nil)
	if err != nil {
		return nil, err
	}

	return state.client.GenerateToAddress(args.NumBlocks, address, &args.MaxTries)
}

func (c *Client) GetAccount(ctx context.Context, conState jsonrpc.State, address string) (string, error) {
	state := State(conState)
	if state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	decodedAddress, err := btcutil.DecodeAddress(address, nil)
	if err != nil {
		return "", err
	}

	return state.client.GetAccount(decodedAddress)
}

func (c *Client) GetAccountAddress(ctx context.Context, conState jsonrpc.State, account string) (string, error) {
	state := State(conState)
	if state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	address, err := state.client.GetAccountAddress(account)
	if err != nil {
		return "", err
	}

	return address.EncodeAddress(), nil
}

func (c *Client) GetAddressInfo(ctx context.Context, conState jsonrpc.State, address string) (*btcjson.GetAddressInfoResult, error) {
	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetAddressInfo(address)
}

func (c *Client) GetAddressesByAccount(ctx context.Context, conState jsonrpc.State, account string) ([]string, error) {
	state := State(conState)
	if state.client == nil {
		return []string{}, pkg.ErrClientNotConnected{}
	}

	addresses, err := state.client.GetAddressesByAccount(account)
	if err != nil {
		return []string{}, err
	}

	addressesEncoded := []string{}
	for _, address := range addresses {
		addressesEncoded = append(addressesEncoded, address.EncodeAddress())
	}

	return addressesEncoded, nil
}

func (c *Client) GetBalance(ctx context.Context, conState jsonrpc.State, account string) (btcutil.Amount, error) {
	state := State(conState)
	if state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.GetBalance(account)
}

func (c *Client) GetBlockCount(ctx context.Context, conState jsonrpc.State, account string) (int64, error) {
	state := State(conState)
	if state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.GetBlockCount()
}

func (c *Client) GetBlockHash(ctx context.Context, conState jsonrpc.State, blockHeight int64) (*chainhash.Hash, error) {
	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetBlockHash(blockHeight)
}

func (c *Client) GetBlockStats(ctx context.Context, conState jsonrpc.State, hash string) (*btcjson.GetBlockStatsResult, error) {
	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	blockHash, err := chainhash.NewHashFromStr(hash)
	if err != nil {
		return nil, err
	}

	return state.client.GetBlockStats(blockHash, nil)
}

func (c *Client) GetBlockVerboseTx(ctx context.Context, conState jsonrpc.State, hash string) (*btcjson.GetBlockVerboseTxResult, error) {
	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	blockHash, err := chainhash.NewHashFromStr(hash)
	if err != nil {
		return nil, err
	}

	return state.client.GetBlockVerboseTx(blockHash)
}
