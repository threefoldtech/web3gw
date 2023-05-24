package atomicswap

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/threefoldtech/atomicswap/eth"
	goethclient "github.com/threefoldtech/web3_proxy/server/clients/eth"
)

type (
	// EthDriver implements Ethereum specific atomic swap logic
	EthDriver struct {
		eth    *goethclient.Client
		client *eth.EthClient
		sct    *eth.SwapContractTransactor
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

	// ErrTxUnconfirmed indicates the transaction is not confirmed after waiting for some amount of time
	ErrTxUnconfirmed          = errors.New("transaction is not confirmed after waiting")
	ErrContractUndervalued    = errors.New("contract has less than expected value")
	ErrWrongContract          = errors.New("call is for wrong contract")
	ErrDifferentSwapReceiver  = errors.New("swap is for a different receiver")
	ErrWrongRefundAddress     = errors.New("contract refund address is wrong")
	ErrContractExpiresTooSoon = errors.New("contract expires too soon")
)

// newEthDriver creates a new eth driver
func newEthDriver(ctx context.Context, cl *goethclient.Client) (*EthDriver, error) {
	dialCtx, cancel := context.WithTimeout(ctx, dialTimeout)
	defer cancel()
	client, err := eth.DialClient(dialCtx, cl.Url) // TODO: should probably be able to construct this from the existing client
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial eth node")
	}

	driver := &EthDriver{
		//eth:    cl,
		client: client,
	}

	if err := driver.initSct(ctx); err != nil {
		return nil, errors.Wrap(err, "could not initialize swap contract transactor")
	}

	return driver, nil
}

// initSct initializes the swap contract transactor
func (e *EthDriver) initSct(ctx context.Context) error {
	sct, err := eth.NewSwapContractTransactor(ctx, e.client, contractAddress, e.eth.Key, sepoliaChainId)
	if err != nil {
		return errors.Wrap(err, "failed to construct swap contract transactor")
	}
	e.sct = &sct

	return nil
}

// InitPayment implements BuyChain
func (e *EthDriver) InitPayment(ctx context.Context, tftAmount uint64, tftPrice uint64, destination string) (any, SwapSecret, SwapSecretHash, error) {
	if len(destination) < 40 {
		return nil, SwapSecret{}, SwapSecretHash{}, fmt.Errorf("truncated destination address %s", destination)
	}
	dest := common.HexToAddress(destination)
	// total wei = swap.amount * swap.price
	cost := big.NewInt(0).Mul(big.NewInt(int64(tftAmount)), big.NewInt(int64(tftPrice)))
	output, err := eth.Initiate(ctx, *e.sct, dest, cost)
	if err != nil {
		return nil, SwapSecret{}, SwapSecretHash{}, errors.Wrap(err, "failed to initiate ETH swap")
	}

	return InitiateEthOutput{
		EthAddress:          output.InitiatorAddress,
		InitiateTransaction: &output.ContractTransaction,
	}, output.Secret, output.SecretHash, nil
}

// ValidateInitPaymentResult implements BuyChain
func (e *EthDriver) ValidateInitPaymentResult(ctx context.Context, initResult any, details NegotiatedTrade) (SwapSecretHash, error) {
	req, ok := initResult.(InitiateEthOutput)
	if !ok {
		return SwapSecretHash{}, errors.New("eth init result is not the proper type")
	}
	deadline := time.Now().Add(time.Minute * 5)
	var auditOutput eth.AuditContractOutput
	var err error
	for {
		auditOutput, err = eth.AuditContract(ctx, *e.sct, req.InitiateTransaction)
		if err != nil {
			if errors.Is(err, eth.ErrTxPending) {
				if time.Now().After(deadline) {
					return SwapSecretHash{}, ErrTxUnconfirmed
				}
				log.Debug().Msg("Tx not confirmed yet, sleeping and trying again")
				time.Sleep(time.Second * 15)
				continue
			}
			return SwapSecretHash{}, errors.Wrap(err, "could not audit ETH contract")
		}
		break
	}

	if auditOutput.ContractAddress != contractAddress {
		log.Warn().Msg("Call is for wrong contract, ignore")
		return SwapSecretHash{}, ErrWrongContract
	}

	// Check the Eth locked in the contract. Notice that we will shamelessly accept if the buyer pays too much
	expectedEthValue := big.NewInt(0).Mul(big.NewInt(int64(details.Amount)), big.NewInt(int64(details.Price)))
	if auditOutput.ContractValue.Cmp(expectedEthValue) == -1 {
		return SwapSecretHash{}, ErrContractUndervalued
	}

	if auditOutput.RecipientAddress != e.eth.AddressFromKey() {
		return SwapSecretHash{}, ErrDifferentSwapReceiver
	}

	// TODO: Strictly speaking we don't really care for this
	if auditOutput.RefundAddress != req.EthAddress {
		return SwapSecretHash{}, ErrWrongRefundAddress
	}

	if time.Unix(auditOutput.Locktime, 0).Before(time.Now().Add(time.Hour * 2)) {
		return SwapSecretHash{}, ErrContractExpiresTooSoon
	}

	return auditOutput.SecretHash, nil
}

// Claim implements BuyChain
func (e *EthDriver) Claim(ctx context.Context, initResult any, secretHash SwapSecretHash, secret SwapSecret) (string, error) {
	redeemOutput, err := eth.Redeem(ctx, *e.sct, secretHash, secret)
	if err != nil {
		return "", errors.Wrap(err, "could not redeem contract")
	}

	return redeemOutput.RedeemTxHash.Hex(), nil
}
