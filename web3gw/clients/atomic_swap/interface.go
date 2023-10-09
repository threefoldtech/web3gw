package atomicswap

import (
	"context"
	"crypto/sha256"

	"github.com/pkg/errors"
)

type (
	// BuyChain holds all logic regarding the chain used to pay for TFT
	BuyChain interface {
		// Address on the buy chain of the loaded account
		Address() string
		// InitPayment initializes a payment transaction on the chain
		InitPayment(ctx context.Context, tftAmount uint64, tftPrice uint64, destination string) (any, SwapSecret, SwapSecretHash, error)
		// ValidateInitPaymentResult validates the result of an init payment call
		ValidateInitPaymentResult(ctx context.Context, initResult any, details NegotiatedTrade) (SwapSecretHash, error)
		// Claim payment on chain
		Claim(ctx context.Context, initResult any, secretHash SwapSecretHash, secret SwapSecret) (string, error)
	}

	// SellChain holds all logic regarding the chain on which TFT are sold
	SellChain interface {
		// Address on the sell chain of the loaded account
		Address() string
		// InitTFTTransfer locks TFT's in the contract
		InitTFTTransfer(ctx context.Context, details NegotiatedTrade, sharedSecret SwapSecretHash, destination string) (any, error)
		// ValidateTFTTranser validates the locked TFT's
		ValidateTFTTranser(ctx context.Context, initTransferResult any, details NegotiatedTrade, sharedSecret SwapSecretHash) error
		// ClaimTFT claims the locked TFT's
		ClaimTFT(ctx context.Context, initTransferResult any, secret SwapSecret) error
	}

	// NegotiatedTrade holds data both parties agreed on in a trade
	NegotiatedTrade struct {
		// Amount of TFT
		Amount uint64 `json:"amount"`
		// Price of 1 TFT in the smallest unit of the buying currency
		Price uint64 `json:"price"`
	}

	// SwapSecret generated for an atomic swap
	SwapSecret [32]byte
	// SwapSecretHash is the sha256 hash of the swap secret used in a trade
	SwapSecretHash [sha256.Size]byte
)

var (
	// ErrTxUnconfirmed indicates the transaction is not confirmed after waiting for some amount of time
	ErrTxUnconfirmed          = errors.New("transaction is not confirmed after waiting")
	ErrContractUndervalued    = errors.New("contract has less than expected value")
	ErrWrongContract          = errors.New("call is for wrong contract")
	ErrDifferentSwapReceiver  = errors.New("swap is for a different receiver")
	ErrWrongRefundAddress     = errors.New("contract refund address is wrong")
	ErrContractExpiresTooSoon = errors.New("contract expires too soon")
	ErrCorruptRefundTx        = errors.New("could not decode refund transaction")
	ErrCorruptContractValue   = errors.New("could not parse contract value, this is an internal coding error")
	ErrWrongSecret            = errors.New("wrong secret hash in contract")
)
