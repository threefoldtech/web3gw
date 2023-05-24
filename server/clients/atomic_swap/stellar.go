package atomicswap

import (
	"context"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/txnbuild"
	"github.com/threefoldtech/atomicswap/stellar"
	stellargoclient "github.com/threefoldtech/web3_proxy/server/clients/stellar"
)

type (
	// StellarDriver implements stellar specific atomic swap logic
	StellarDriver struct {
		stellar       *stellargoclient.Client
		horizonClient *horizonclient.Client

		networkPassphrase string
		asset             txnbuild.Asset
	}

	// ParticipateOutput generated when initiating TFT transfer
	InitTFTTransferResult struct {
		HoldingAccount string `json:"holdingAccount"`
		RefundTx       string `json:"refundTx"`
	}
)

// InitTFTTransfer implements SellChain
func (s *StellarDriver) InitTFTTransfer(ctx context.Context, details NegotiatedTrade, sharedSecret SwapSecretHash, destination string) (any, error) {
	kp := s.stellar.KeyPair()
	participateOutput, err := stellar.Participate(s.networkPassphrase, &kp, destination, strconv.FormatUint(uint64(details.Amount), 10), sharedSecret[:], s.asset, s.horizonClient)
	if err != nil {
		return nil, errors.Wrap(err, "could not participate on stellar side")
	}
	return InitTFTTransferResult{
		HoldingAccount: participateOutput.HoldingAccountAddress,
		RefundTx:       participateOutput.RefundTransaction,
	}, nil
}

// ValidateTFTTranser implements SellChain
func (s *StellarDriver) ValidateTFTTranser(ctx context.Context, initTransferResult any, details NegotiatedTrade, sharedSecret SwapSecretHash) error {
	initResult, ok := initTransferResult.(InitTFTTransferResult)
	if !ok {
		return errors.New("TFT transfer init result is not of proper type")
	}

	refundTx := txnbuild.Transaction{}
	if err := (&refundTx).UnmarshalText([]byte(initResult.RefundTx)); err != nil {
		return ErrCorruptRefundTx
	}
	auditOutput, err := stellar.AuditContract(s.networkPassphrase, refundTx, initResult.HoldingAccount, s.asset, s.horizonClient)
	if err != nil {
		return errors.Wrap(err, "could not audit stellar contract")
	}

	contractValue, err := strconv.ParseFloat(auditOutput.ContractValue, 64)
	if err != nil {
		return ErrCorruptContractValue
	}

	// if the seller wants to give us more TFT than agreed, we will shamelessly accept
	if contractValue < float64(details.Amount) {
		return ErrContractUndervalued
	}

	// Make sure we are the receiver
	if auditOutput.RecipientAddress != s.stellar.Address() {
		return ErrDifferentSwapReceiver
	}

	// Verify that the secret is properly set
	if auditOutput.SecretHash != hex.EncodeToString(sharedSecret[:]) {
		return ErrWrongSecret
	}

	if time.Unix(auditOutput.Locktime, 0).Before(time.Now().Add(time.Hour * 1)) {
		log.Warn().Msg("Contract doesn't leave at least 1 hour to complete, ignore")
		return ErrContractExpiresTooSoon
	}

	return nil
}

// ClaimTFT implements SellChain
func (s *StellarDriver) ClaimTFT(ctx context.Context, initTransferResult any, secret SwapSecret) error {
	initResult, ok := initTransferResult.(InitTFTTransferResult)
	if !ok {
		return errors.New("TFT transfer init result is not of proper type")
	}
	kp := s.stellar.KeyPair()
	_, err := stellar.Redeem(s.networkPassphrase, &kp, initResult.HoldingAccount, secret[:], s.horizonClient)
	if err != nil {
		return errors.Wrap(err, "failed to redeem stellar contract")
	}

	return nil
}

// Parse the stellar testnet TFT asset
func mustStellarTestnetTftAsset() txnbuild.Asset {
	a, err := txnbuild.ParseAssetString("TFT:GA47YZA3PKFUZMPLQ3B5F2E3CJIB57TGGU7SPCQT2WAEYKN766PWIMB3")
	if err != nil {
		panic(err)
	}
	return a
}

// Parse the stellar mainnet TFT asset
func mustStellarTftAsset() txnbuild.Asset {
	a, err := txnbuild.ParseAssetString("TFT:GBOVQKJYHXRR3DX6NOX2RRYFRCUMSADGDESTDNBDS6CDVLGVESRTAC47")
	if err != nil {
		panic(err)
	}
	return a
}
