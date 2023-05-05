package btc

import (
	"context"
	"errors"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	btcRpcClient "github.com/btcsuite/btcd/rpcclient"
	"github.com/rs/zerolog/log"
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
		Host   string `json:"host"`
		User   string `json:"user"`
		Pass   string `json:"pass"`
		Wallet string `json:"wallet"`
	}

	ImportAddress struct {
		Address string `json:"address"`
		Label   string `json:"label"`
		Rescan  bool   `json:"rescan"`
		P2SH    bool   `json:"p2sh"`
	}

	ImportPrivKey struct {
		WIF    string `json:"wif"`
		Label  string `json:"label"`
		Rescan bool   `json:"rescan"`
	}

	ImportPubKey struct {
		PubKey string `json:"pub_key"`
		Label  string `json:"label"`
		Rescan bool   `json:"rescan"`
	}

	RenameAccount struct {
		OldAccount string `json:"old_account"`
		NewAccount string `json:"new_account"`
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

	GetChainTxStatsNBlocksBlockHash struct {
		AmountOfBlocks int32  `json:"amount_of_blocks"`
		BlockHashEnd   string `json:"block_hash_end"`
	}

	CreateWallet struct {
		Name               string `json:"name"`
		DisablePrivateKeys bool   `json:"disable_private_keys"`
		CreateBlackWallet  bool   `json:"create_blank_wallet"`
		Passphrase         string `json:"passphrase"`
		AvoidReuse         bool   `json:"avoid_reuse"`
	}

	GetNewAddress struct {
		Label       string `json:"label"`
		AddressType string `json:"address_type"`
	}

	Move struct {
		FromAccount      string         `json:"from_account"`
		ToAccount        string         `json:"to_account"`
		Amount           btcutil.Amount `json:"amount"`
		MinConfirmations int            `json:"min_confirmations"`
		Comment          string         `json:"comment"`
	}

	SetLabel struct {
		Address string `json:"address"`
		Label   string `json:"label"`
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

// Close implements jsonrpc.Closer
func (s *btcState) Close() {}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Load(ctx context.Context, conState jsonrpc.State, args Load) error {
	log.Debug().Msgf("BTC: connecting to btc node %s", args.Host)

	client, err := btcRpcClient.New(
		&btcRpcClient.ConnConfig{
			Host:         args.Host + "/wallet/" + args.Wallet,
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

func (c *Client) ImportAddress(ctx context.Context, conState jsonrpc.State, args ImportAddress) error {
	log.Debug().Msgf("BTC: importing address %s with label %s (rescan: %t, p2sh: %s)", args.Address, args.Label, args.Rescan, args.P2SH)

	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ImportAddressRescan(args.Address, args.Label, args.Rescan)
}

func (c *Client) ImportPrivKeyRescan(ctx context.Context, conState jsonrpc.State, args ImportPrivKey) error {
	log.Debug().Msgf("BTC: importing private key to label %s (rescan: %t)", args.Label, args.Rescan)

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

func (c *Client) ImportPubKey(ctx context.Context, conState jsonrpc.State, args ImportPubKey) error {
	log.Debug().Msgf("BTC: importing public key (rescan: %t)", args.Rescan)

	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ImportPubKeyRescan(args.PubKey, args.Rescan)
}

func (c *Client) ListLabels(ctx context.Context, conState jsonrpc.State) (map[string]btcutil.Amount, error) {
	log.Debug().Msg("BTC: listing labels")

	state := State(conState)
	if state.client == nil {
		return map[string]btcutil.Amount{}, pkg.ErrClientNotConnected{}
	}

	return state.client.ListAccounts()
}

func (c *Client) RenameAccount(ctx context.Context, conState jsonrpc.State, args RenameAccount) error {
	log.Debug().Msgf("BTC: renaming account from %s to %s", args.OldAccount, args.NewAccount)

	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.RenameAccount(args.OldAccount, args.NewAccount)
}

func (c *Client) SendToAddress(ctx context.Context, conState jsonrpc.State, args SendToAddress) (string, error) {
	log.Debug().Msgf("BTC: sending %d to address %s with comment %s and commentTo %s", args.Amount, args.Address, args.Comment, args.CommentTo)

	state := State(conState)
	if state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	address, err := btcutil.DecodeAddress(args.Address, nil)
	if err != nil {
		return "", err
	}

	var blockHash *chainhash.Hash = nil
	if args.Comment != "" {
		blockHash, err = state.client.SendToAddressComment(address, args.Amount, args.Comment, args.CommentTo)
	} else {
		blockHash, err = state.client.SendToAddress(address, args.Amount)
	}
	if err != nil || blockHash == nil {
		return "", err
	}
	return blockHash.String(), err
}

func (c *Client) EstimateSmartFee(ctx context.Context, conState jsonrpc.State, args EstimateSmartFee) (*btcjson.EstimateSmartFeeResult, error) {
	log.Debug().Msgf("BTC: estimating smart fee for %s blocks with estimation mode %s", args.ConfTarget, args.Mode)

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.EstimateSmartFee(args.ConfTarget, &args.Mode)
}

func hashesToStrings(hashes []*chainhash.Hash) []string {
	var blockHashes = []string{}
	for _, hash := range hashes {
		if hash == nil {
			blockHashes = append(blockHashes, "")
		} else {
			blockHashes = append(blockHashes, hash.String())
		}
	}
	return blockHashes
}

func (c *Client) GenerateBlocks(ctx context.Context, conState jsonrpc.State, numBlocks uint32) ([]string, error) {
	log.Debug().Msgf("BTC: generating %d blocks", numBlocks)

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	hashes, err := state.client.Generate(numBlocks)
	if err != nil {
		return []string{}, err
	}

	return hashesToStrings(hashes), err
}

func (c *Client) GenerateBlocksToAddress(ctx context.Context, conState jsonrpc.State, args GenerateToAddress) ([]string, error) {
	log.Debug().Msgf("BTC: generating %d blocks for address %s", args.NumBlocks, args.Address)

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	address, err := btcutil.DecodeAddress(args.Address, nil)
	if err != nil {
		return nil, err
	}

	hashes, err := state.client.GenerateToAddress(args.NumBlocks, address, &args.MaxTries)
	if err != nil {
		return nil, err
	}

	return hashesToStrings(hashes), err
}

func (c *Client) SetLabel(ctx context.Context, conState jsonrpc.State, args SetLabel) error {
	log.Debug().Msgf("BTC: setting label on address %s: %s", args.Address, args.Label)

	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	decodedAddress, err := btcutil.DecodeAddress(args.Address, nil)
	if err != nil {
		return err
	}

	return state.client.SetAccount(decodedAddress, args.Label)
}

func (c *Client) GetAddressInfo(ctx context.Context, conState jsonrpc.State, address string) (*btcjson.GetAddressInfoResult, error) {
	log.Debug().Msgf("BTC: getting address info for %s", address)

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetAddressInfo(address)
}

func (c *Client) GetAddressesByLabel(ctx context.Context, conState jsonrpc.State, label string) ([]string, error) {
	log.Debug().Msgf("BTC: getting addresses by label %s", label)

	state := State(conState)
	if state.client == nil {
		return []string{}, pkg.ErrClientNotConnected{}
	}

	addresses, err := state.client.GetAddressesByAccount(label)
	if err != nil {
		return []string{}, err
	}

	addressesEncoded := []string{}
	for _, address := range addresses {
		addressesEncoded = append(addressesEncoded, address.EncodeAddress())
	}

	return addressesEncoded, nil
}

func (c *Client) GetBalance(ctx context.Context, conState jsonrpc.State) (btcutil.Amount, error) {
	log.Debug().Msgf("BTC: getting balance of wallet")

	state := State(conState)
	if state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.GetBalance("*")
}

func (c *Client) GetBlockCount(ctx context.Context, conState jsonrpc.State) (int64, error) {
	log.Debug().Msg("BTC: getting block count")

	state := State(conState)
	if state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.GetBlockCount()
}

func (c *Client) GetBlockHash(ctx context.Context, conState jsonrpc.State, blockHeight int64) (string, error) {
	log.Debug().Msgf("BTC: getting block hash for block at height %d", blockHeight)

	state := State(conState)
	if state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	blockHash, err := state.client.GetBlockHash(blockHeight)
	if err != nil {
		return "", err
	}

	return blockHash.String(), nil
}

func (c *Client) GetBlockStats(ctx context.Context, conState jsonrpc.State, hash string) (*btcjson.GetBlockStatsResult, error) {
	log.Debug().Msgf("BTC: getting block stats for block with hash %s", hash)

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetBlockStats(hash, nil)
}

func (c *Client) GetBlockChainInfo(ctx context.Context, conState jsonrpc.State) (*btcjson.GetBlockChainInfoResult, error) {
	log.Debug().Msg("BTC: getting blockchain info")

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetBlockChainInfo()
}

func (c *Client) GetBlockVerboseTx(ctx context.Context, conState jsonrpc.State, hash string) (*btcjson.GetBlockVerboseTxResult, error) {
	log.Debug().Msgf("BTC: getting block verbose tx for block at height %s", hash)

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

func (c *Client) GetChainTxStats(ctx context.Context, conState jsonrpc.State, args GetChainTxStatsNBlocksBlockHash) (*btcjson.GetChainTxStatsResult, error) {
	log.Debug().Msgf("BTC: getting chain transaction statistics (amount_of_blocks:%d, hash_block_end:%s)", args.AmountOfBlocks, args.BlockHashEnd)

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	if args.AmountOfBlocks > 0 {
		if args.BlockHashEnd != "" {
			blockHash, err := chainhash.NewHashFromStr(args.BlockHashEnd)
			if err != nil {
				return nil, err
			}
			return state.client.GetChainTxStatsNBlocksBlockHash(args.AmountOfBlocks, *blockHash)
		}
		return state.client.GetChainTxStatsNBlocks(args.AmountOfBlocks)
	}
	return state.client.GetChainTxStats()
}

func (c *Client) GetConnectionCount(ctx context.Context, conState jsonrpc.State) (int64, error) {
	log.Debug().Msg("BTC: connection count")

	state := State(conState)
	if state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.GetConnectionCount()
}

func (c *Client) GetDifficulty(ctx context.Context, conState jsonrpc.State) (float64, error) {
	log.Debug().Msg("BTC: getting difficulty")

	state := State(conState)
	if state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.GetDifficulty()
}

func (c *Client) GetMiningInfo(ctx context.Context, conState jsonrpc.State) (*btcjson.GetMiningInfoResult, error) {
	log.Debug().Msg("BTC: getting mining info")

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetMiningInfo()
}

func (c *Client) GetNewAddress(ctx context.Context, conState jsonrpc.State, args GetNewAddress) (string, error) {
	log.Debug().Msgf("BTC: getting new address of type %s for label %s", args.AddressType, args.Label)

	state := State(conState)
	if state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	var address btcutil.Address
	var err error
	if args.AddressType != "" {
		address, err = state.client.GetNewAddressType(args.Label, args.AddressType)
	} else {
		address, err = state.client.GetNewAddress(args.Label)
	}
	if err != nil {
		return "", err
	}

	return address.EncodeAddress(), nil
}

func (c *Client) GetNodeAddresses(ctx context.Context, conState jsonrpc.State) ([]btcjson.GetNodeAddressesResult, error) {
	log.Debug().Msg("BTC: getting node addresses")

	state := State(conState)
	if state.client == nil {
		return []btcjson.GetNodeAddressesResult{}, pkg.ErrClientNotConnected{}
	}

	return state.client.GetNodeAddresses(nil)
}

func (c *Client) GetPeerInfo(ctx context.Context, conState jsonrpc.State) ([]btcjson.GetPeerInfoResult, error) {
	log.Debug().Msg("BTC: getting peer info")

	state := State(conState)
	if state.client == nil {
		return []btcjson.GetPeerInfoResult{}, pkg.ErrClientNotConnected{}
	}

	return state.client.GetPeerInfo()
}

func (c *Client) GetRawTransaction(ctx context.Context, conState jsonrpc.State, txHash string) (*btcutil.Tx, error) {
	log.Debug().Msgf("BTC: getting raw transaction with hash %s", txHash)

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	txHashDecoded, err := chainhash.NewHashFromStr(txHash)
	if err != nil {
		return nil, err
	}

	return state.client.GetRawTransaction(txHashDecoded)
}

func (c *Client) GetReceivedByLabel(ctx context.Context, conState jsonrpc.State, label string) (btcutil.Amount, error) {
	log.Debug().Msgf("BTC: getting amount received by label %s", label)

	state := State(conState)
	if state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.GetReceivedByAccount(label)
}

func (c *Client) LoadWallet(ctx context.Context, conState jsonrpc.State, walletName string) (*btcjson.LoadWalletResult, error) {
	log.Debug().Msgf("BTC: loading wallet %s", walletName)

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.LoadWallet(walletName)
}

func (c *Client) GetWalletInfo(ctx context.Context, conState jsonrpc.State) (*btcjson.GetWalletInfoResult, error) {
	log.Debug().Msg("BTC: getting wallet info")

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetWalletInfo()
}

func (c *Client) ListReceivedByLabel(ctx context.Context, conState jsonrpc.State) ([]btcjson.ListReceivedByAccountResult, error) {
	log.Debug().Msg("BTC: listing received transactions by label")

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.ListReceivedByAccount()
}

func (c *Client) ListReceivedByAddress(ctx context.Context, conState jsonrpc.State) ([]btcjson.ListReceivedByAddressResult, error) {
	log.Debug().Msg("BTC: listing received transactions by address")

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.ListReceivedByAddress()
}

func (c *Client) ListSinceBlock(ctx context.Context, conState jsonrpc.State, hash string) (*btcjson.ListSinceBlockResult, error) {
	log.Debug().Msg("BTC: listing since block")

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	blockHash, err := chainhash.NewHashFromStr(hash)
	if err != nil {
		return nil, err
	}

	return state.client.ListSinceBlock(blockHash)
}

func (c *Client) ListTransactions(ctx context.Context, conState jsonrpc.State, label string) ([]btcjson.ListTransactionsResult, error) {
	log.Debug().Msg("BTC: listing transactions for label %s")

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.ListTransactions(label)
}

func (c *Client) CreateWallet(ctx context.Context, conState jsonrpc.State, args CreateWallet) (*btcjson.CreateWalletResult, error) {
	log.Debug().Msgf("BTC: creating wallet with name %s (AvoidReuse:%t, CreateBlackWallet:%t, DisablePrivateKeys:%t)", args.Name, args.AvoidReuse, args.CreateBlackWallet, args.DisablePrivateKeys)

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}
	if args.Passphrase == "" {
		return nil, errors.New("passphrase cannot be empty")
	}

	options := []btcRpcClient.CreateWalletOpt{}
	options = append(options, btcRpcClient.WithCreateWalletPassphrase(args.Passphrase))
	if args.DisablePrivateKeys {
		options = append(options, btcRpcClient.WithCreateWalletDisablePrivateKeys())
	}
	if args.AvoidReuse {
		options = append(options, btcRpcClient.WithCreateWalletAvoidReuse())
	}
	if args.CreateBlackWallet {
		options = append(options, btcRpcClient.WithCreateWalletBlank())
	}

	return state.client.CreateWallet(args.Name, options...)
}

func (c *Client) SetTxFee(tx context.Context, conState jsonrpc.State, fee btcutil.Amount) error {
	log.Debug().Msg("BTC: setting transaction fee")

	state := State(conState)
	if state.client == nil {
		return nil
	}

	return state.client.SetTxFee(fee)
}

func (c *Client) Move(ctx context.Context, conState jsonrpc.State, args Move) (bool, error) {
	log.Debug().Msgf("BTC: moving %d from account %s to %s", args.Amount, args.FromAccount, args.ToAccount)

	state := State(conState)
	if state.client == nil {
		return false, pkg.ErrClientNotConnected{}
	}

	if args.MinConfirmations > 0 {
		if args.Comment != "" {
			return state.client.MoveComment(args.FromAccount, args.ToAccount, args.Amount, args.MinConfirmations, args.Comment)
		}
		return state.client.MoveMinConf(args.FromAccount, args.ToAccount, args.Amount, args.MinConfirmations)
	}
	return state.client.Move(args.FromAccount, args.ToAccount, args.Amount)
}
