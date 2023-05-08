package btclightning

import (
	"context"
	"fmt"
	"time"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcwallet/walletdb"
	"github.com/lightninglabs/neutrino"
	"github.com/lightninglabs/neutrino/banman"
	"github.com/lightninglabs/neutrino/headerfs"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const (
	BtcLightningID = "btc_lightning"
)

type btcLightningState struct {
	service       *neutrino.ChainService
	rescanContext *rescanContext
}

type Client struct {
}

// State from a connection. If no state is present, it is initialized
func State(conState jsonrpc.State) *btcLightningState {
	raw, exists := conState[BtcLightningID]
	if !exists {
		ns := &btcLightningState{}
		conState[BtcLightningID] = ns
		return ns
	}
	ns, ok := raw.(*btcLightningState)
	if !ok {
		// This means the invariant is violated, so panic here is ok
		panic("Invalid saved state for btc")
	}
	return ns
}

func NewClient() *Client {
	return &Client{}
}

// NewChainService returns a new chain service configured to connect to the
// bitcoin network type specified by chainParams.  Use start to begin syncing
// with peers.
func (c *Client) NewChainService(ctx context.Context, conState jsonrpc.State, serviceConfig ChainServiceConfig) error {
	state := State(conState)

	params, err := getChainServiceParams(serviceConfig.ChainNet)
	if err != nil {
		return err
	}

	db, err := getChainServiceDB(serviceConfig.DatabasePath)
	if err != nil {
		return fmt.Errorf("failed to open or create db with path %s", serviceConfig.DatabasePath)
	}

	cfg := neutrino.Config{
		DataDir:          serviceConfig.DataDir,
		Database:         db,
		ChainParams:      params,
		ConnectPeers:     serviceConfig.ConnectPeers,
		AddPeers:         serviceConfig.AddPeers,
		FilterCacheSize:  serviceConfig.FilterCacheSize,
		BlockCacheSize:   serviceConfig.BlockCacheSize,
		PersistToDisk:    serviceConfig.PersistToDisk,
		BroadcastTimeout: time.Duration(serviceConfig.BroadcastTimeout),
	}

	service, err := neutrino.NewChainService(cfg)
	if err != nil {
		return errors.Wrapf(err, "failed to create new chain service")
	}

	if state.service != nil {
		state.service.Stop()
	}

	state.service = service

	return nil
}

func getChainServiceDB(path string) (walletdb.DB, error) {
	db, err := walletdb.Open("bdb", path, true, time.Second*10)
	if err == nil {
		return db, nil
	}

	return walletdb.Create("bdb", path, true, time.Second*10)
}

func getChainServiceParams(net string) (chaincfg.Params, error) {
	switch net {
	case "mainnet":
		return chaincfg.MainNetParams, nil
	case "testnet":
		return chaincfg.TestNet3Params, nil

	}

	return chaincfg.Params{}, fmt.Errorf("chain net should be \"mainnet\" or \"testnet\", but %s was provided", net)
}

func (c *Client) BestBlock(ctx context.Context, conState jsonrpc.State) (BlockStamp, error) {
	state := State(conState)

	if state.service == nil {
		return BlockStamp{}, ChainServiceInstanceNotFoundErr
	}

	blockStamp, err := state.service.BestBlock()
	if err != nil {
		return BlockStamp{}, err
	}

	return BlockStamp{
		Height:    blockStamp.Height,
		Hash:      blockStamp.Hash[:],
		Timestamp: blockStamp.Timestamp.Unix(),
	}, nil
}

func (c *Client) GetBlockHash(ctx context.Context, conState jsonrpc.State, height int64) (string, error) {
	state := State(conState)

	if state.service == nil {
		return "", ChainServiceInstanceNotFoundErr
	}

	hash, err := state.service.GetBlockHash(height)
	if err != nil {
		return "", err
	}

	return hash.String(), nil
}

func (c *Client) GetBlockHeader(ctx context.Context, conState jsonrpc.State, hash string) (BlockHeader, error) {
	state := State(conState)

	if state.service == nil {
		return BlockHeader{}, ChainServiceInstanceNotFoundErr
	}

	chainHash, err := chainhash.NewHash([]byte(hash))
	if err != nil {
		return BlockHeader{}, err
	}

	blockHeader, err := state.service.GetBlockHeader(chainHash)
	if err != nil {
		return BlockHeader{}, err
	}

	return BlockHeader{
		Version:    blockHeader.Version,
		PrevBlock:  blockHeader.PrevBlock[:],
		MerkleRoot: blockHeader.MerkleRoot[:],
		Timestamp:  blockHeader.Timestamp.Unix(),
		Bits:       blockHeader.Bits,
		Nonce:      blockHeader.Nonce,
	}, nil
}

func (c *Client) GetBlockHeight(ctx context.Context, conState jsonrpc.State, hash string) (int32, error) {
	state := State(conState)

	if state.service == nil {
		return 0, ChainServiceInstanceNotFoundErr
	}

	chainHash, err := chainhash.NewHash([]byte(hash))
	if err != nil {
		return 0, err
	}

	return state.service.GetBlockHeight(chainHash)
}

func (c *Client) BanPeer(ctx context.Context, conState jsonrpc.State, peer BanPeerInfo) error {
	state := State(conState)

	if state.service == nil {
		return ChainServiceInstanceNotFoundErr
	}

	return state.service.BanPeer(peer.Address, banman.Reason(peer.Reason))
}

func (c *Client) IsBanned(ctx context.Context, conState jsonrpc.State, address string) (bool, error) {
	state := State(conState)

	if state.service == nil {
		return false, ChainServiceInstanceNotFoundErr
	}

	return state.service.IsBanned(address), nil
}

func (c *Client) AddPeer(ctx context.Context, conState jsonrpc.State, isServerPeerPersistent bool) error {
	state := State(conState)

	if state.service == nil {
		return ChainServiceInstanceNotFoundErr
	}

	peer := neutrino.NewServerPeer(state.service, isServerPeerPersistent)
	state.service.AddPeer(peer)

	return nil
}

func (c *Client) AddBytesSent(ctx context.Context, conState jsonrpc.State, bytesSent uint64) error {
	state := State(conState)

	if state.service == nil {
		return ChainServiceInstanceNotFoundErr
	}

	state.service.AddBytesSent(bytesSent)

	return nil
}

func (c *Client) AddBytesReceived(ctx context.Context, conState jsonrpc.State, bytesReceived uint64) error {
	state := State(conState)

	if state.service == nil {
		return ChainServiceInstanceNotFoundErr
	}

	state.service.AddBytesReceived(bytesReceived)

	return nil
}

func (c *Client) NetTotals(ctx context.Context, conState jsonrpc.State) (NetTotals, error) {
	state := State(conState)

	if state.service == nil {
		return NetTotals{}, ChainServiceInstanceNotFoundErr
	}

	received, sent := state.service.NetTotals()

	return NetTotals{
		Received: received,
		Sent:     sent,
	}, nil
}

func (c *Client) UpdatePeerHeights(ctx context.Context, conState jsonrpc.State, r UpdatePeerHeightsRequest) error {
	state := State(conState)

	if state.service == nil {
		return ChainServiceInstanceNotFoundErr
	}

	hash, err := chainhash.NewHash(r.LatestBlockHash)
	if err != nil {
		return err
	}

	serverPeer := state.service.PeerByAddr(r.SourcePeerAddress)

	state.service.UpdatePeerHeights(hash, r.LatestHeight, serverPeer)

	return nil
}

func (c *Client) StartChainService(ctx context.Context, conState jsonrpc.State) error {
	state := State(conState)

	if state.service == nil {
		return ChainServiceInstanceNotFoundErr
	}

	return state.service.Start()
}

func (c *Client) StopChainService(ctx context.Context, conState jsonrpc.State) error {
	state := State(conState)

	if state.service == nil {
		return ChainServiceInstanceNotFoundErr
	}

	return state.service.Stop()
}

func (c *Client) IsCurrent(ctx context.Context, conState jsonrpc.State) (bool, error) {
	state := State(conState)

	if state.service == nil {
		return false, ChainServiceInstanceNotFoundErr
	}

	return state.service.IsCurrent(), nil
}

func (c *Client) NewRescan(ctx context.Context, conState jsonrpc.State, rescanOptions RescanOptions, queryOptions QueryOptions) error {
	state := State(conState)

	if state.service == nil {
		return ChainServiceInstanceNotFoundErr
	}

	quit := make(chan struct{}, 1)
	options, err := getRescanOptions(state.service, rescanOptions, quit)
	if err != nil {
		return err
	}

	rescan := neutrino.NewRescan(&neutrino.RescanChainSource{ChainService: state.service}, options...)

	if state.rescanContext != nil {
		// TODO: stop running instance of rescan
	}

	state.rescanContext = &rescanContext{
		rescan: rescan,
	}

	return nil
}

func getRescanOptions(service *neutrino.ChainService, rescanOptions RescanOptions, quit <-chan struct{}) ([]neutrino.RescanOption, error) {
	options := []neutrino.RescanOption{}

	endBlock := headerfs.BlockStamp{
		Height:    rescanOptions.EndBlock.Height,
		Hash:      chainhash.Hash(rescanOptions.EndBlock.Hash),
		Timestamp: time.Unix(rescanOptions.EndBlock.Timestamp, 0),
	}
	options = append(options, neutrino.EndBlock(&endBlock))

	startBlock := headerfs.BlockStamp{
		Height:    rescanOptions.StartBlock.Height,
		Hash:      chainhash.Hash(rescanOptions.StartBlock.Hash),
		Timestamp: time.Unix(rescanOptions.StartBlock.Timestamp, 0),
	}
	options = append(options, neutrino.StartBlock(&startBlock))

	startTime := time.Unix(rescanOptions.StartTime, 0)
	options = append(options, neutrino.StartTime(startTime))

	options = append(options, neutrino.TxIdx(rescanOptions.TxIdx))

	inputs := getInputWithScripts(rescanOptions.WatchInputs)
	options = append(options, neutrino.WatchInputs(inputs...))

	options = append(options, neutrino.QuitChan(quit))

	// TODO: implement notification handlers for rescan
	// options = append(options, neutrino.NotificationHandlers(rpcclient.NotificationHandlers{
	// }))

	options = append(options, neutrino.ProgressHandler(func(lastProcessedBlock uint32) {
		log.Debug().Msgf("last processed block: %d", lastProcessedBlock)
	}))

	queryOpts := getQueryOpts(rescanOptions.QueryOptions)
	options = append(options, neutrino.QueryOptions(queryOpts...))

	params := service.ChainParams()
	addresses, err := getAddressesFromPubKeys(&params, rescanOptions.WatchAddresses)
	if err != nil {
		return nil, err
	}
	options = append(options, neutrino.WatchAddrs(addresses...))

	return options, nil
}

func getInputWithScripts(watchInputs []InputWithScript) []neutrino.InputWithScript {
	ret := []neutrino.InputWithScript{}
	for _, input := range watchInputs {
		ret = append(ret, neutrino.InputWithScript{
			OutPoint: wire.OutPoint{
				Hash:  chainhash.Hash(input.OutPoint.Hash),
				Index: input.OutPoint.Index,
			},
			PkScript: []byte(input.PkScript),
		})
	}
	return ret
}

func getQueryOpts(opts QueryOptions) []neutrino.QueryOption {
	options := []neutrino.QueryOption{}

	if opts.InvalidTxThreshold != nil {
		options = append(options, neutrino.InvalidTxThreshold(*opts.InvalidTxThreshold))
	}

	if opts.MaxBatchSize != nil {
		options = append(options, neutrino.MaxBatchSize(*opts.MaxBatchSize))
	}

	if opts.NumRetries != nil {
		options = append(options, neutrino.NumRetries(*opts.NumRetries))
	}

	if opts.OptimisticBatch != nil {
		options = append(options, neutrino.OptimisticBatch())
	}

	if opts.OptimisticReverseBatch != nil {
		options = append(options, neutrino.OptimisticReverseBatch())
	}

	if opts.PeerConnectTimeout != nil {
		options = append(options, neutrino.PeerConnectTimeout(time.Duration(*opts.PeerConnectTimeout)))
	}

	if opts.RejectTimeout != nil {
		options = append(options, neutrino.RejectTimeout(time.Duration(*opts.RejectTimeout)))
	}

	if opts.Timeout != nil {
		options = append(options, neutrino.Timeout(time.Duration(*opts.Timeout)))
	}

	return options
}

func getAddressesFromPubKeys(netParams *chaincfg.Params, pubKeys []string) ([]btcutil.Address, error) {

	addresses := []btcutil.Address{}
	for _, pubKey := range pubKeys {
		key, err := btcec.ParsePubKey([]byte(pubKey))
		if err != nil {
			return nil, err
		}

		pubKeyHash := btcutil.Hash160(key.SerializeCompressed())
		addr, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, netParams)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, addr)
	}

	return addresses, nil
}

func (c *Client) StartRescan(ctx context.Context, conState jsonrpc.State) error {
	state := State(conState)

	if state.rescanContext == nil {
		return RescanInstanceNotFoundErr
	}

	state.rescanContext.errChan = state.rescanContext.rescan.Start()

	return nil
}

func (c *Client) UpdateRescan(ctx context.Context, conState jsonrpc.State, updateOptions RescanUpdateOptions) error {
	state := State(conState)

	if state.rescanContext == nil {
		return RescanInstanceNotFoundErr
	}

	options, err := getUpdateOptions(state.service, updateOptions)
	if err != nil {
		return err
	}

	state.rescanContext.rescan.Update(options...)

	return nil
}

func getUpdateOptions(service *neutrino.ChainService, updateOptions RescanUpdateOptions) ([]neutrino.UpdateOption, error) {
	options := []neutrino.UpdateOption{}

	params := service.ChainParams()
	addresses, err := getAddressesFromPubKeys(&params, updateOptions.AddAddresses)
	if err != nil {
		return nil, err
	}

	options = append(options, neutrino.AddAddrs(addresses...))

	inputs := getInputWithScripts(updateOptions.WatchInputs)
	options = append(options, neutrino.AddInputs(inputs...))

	if updateOptions.DisableDisconnectedNtfns != nil {
		options = append(options, neutrino.DisableDisconnectedNtfns(*updateOptions.DisableDisconnectedNtfns))
	}

	if updateOptions.Rewind != nil {
		options = append(options, neutrino.Rewind(*updateOptions.Rewind))
	}

	return options, nil
}

func (c *Client) RescanShutdown(ctx context.Context, conState jsonrpc.State) error {
	state := State(conState)

	if state.rescanContext == nil {
		return RescanInstanceNotFoundErr
	}

	close(state.rescanContext.quit)

	state.rescanContext.rescan.WaitForShutdown()

	return nil
}

// TODO: implement NewUtxoScanner

// func (c *Client) NewUtxoScanner(ctx context.Context, conState jsonrpc.State) error{

// 	return nil
// }
