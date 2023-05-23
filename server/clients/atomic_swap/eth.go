package atomicswap

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/threefoldtech/atomicswap/eth"
	goethclient "github.com/threefoldtech/web3_proxy/server/clients/eth"
)

type (
	// EthDriver implements Ethereum specific atomic swap logic
	EthDriver struct {
		eth *goethclient.Client
		sct *eth.SwapContractTransactor
	}

	// InitiateEthOutput is the result of initiating a swap on an ethereum chain
	InitiateEthOutput struct {
		// EthAddress of the initiator (which will reclaim the funds if the time passes)
		EthAddress common.Address `json:"ethAddress"`
		// InitiateTransaction is the transaction which initiates the swap, including the passed parameters
		InitiateTransaction *types.Transaction `json:"initiateTransaction"`
	}
)

var (
	// chain ID of the goerli network
	goerliChainID = big.NewInt(5)
	// chain ID of the sepolia network
	sepoliaChainId = big.NewInt(11155111)
	// contract address on the sepolia test network
	sepoliaContractAddress = common.HexToAddress("0x17f54245073bfed168a51c3d13b536e39e406063")
	// contract address on the goerli network
	goerliContractAddress = common.HexToAddress("0x8420c8271d602F6D0B190856Cea8E74D09A0d3cF")
)

// InitPayment implements BuyChain
func (e *EthDriver) InitPayment(ctx context.Context, tftAmount uint, tftPrice float64) (any, error) {
	return nil, errors.New("TODO")
}

// ValidateInitPaymentResult implements BuyChain
func (e *EthDriver) ValidateInitPaymentResult(ctx context.Context, initResult any, details NegotiatedTrade) error {
	return errors.New("TODO")
}

// Claim implements BuyChain
func (e *EthDriver) Claim(ctx context.Context, initResult any, secret SwapSecret) error {
	return errors.New("TODO")
}
