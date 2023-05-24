package atomicswap

import (
	"context"
	"strconv"

	"github.com/pkg/errors"
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
	ParticipateOutput struct {
		HoldingAccount string
		RefundTx       string
	}
)

// InitTFTTransfer implements SellChain
func (s *StellarDriver) InitTFTTransfer(ctx context.Context, details NegotiatedTrade, sharedSecret SwapSecretHash, destination string) (any, error) {
	kp := s.stellar.KeyPair()
	participateOutput, err := stellar.Participate(s.networkPassphrase, &kp, destination, strconv.FormatUint(uint64(details.Amount), 10), sharedSecret[:], s.asset, s.horizonClient)
	if err != nil {
		return nil, errors.Wrap(err, "could not participate on stellar side")
	}
	return ParticipateOutput{
		HoldingAccount: participateOutput.HoldingAccountAddress,
		RefundTx:       participateOutput.RefundTransaction,
	}, nil
}

// ValidateTFTTranser implements SellChain
func (s *StellarDriver) ValidateTFTTranser(ctx context.Context, initTransferResult any, sharedSecret SwapSecretHash) error {
	return errors.New("TODO")
}

// ClaimTFT implements SellChain
func (s *StellarDriver) ClaimTFT(ctx context.Context, initTransferResult any, secret SwapSecret) error {
	return errors.New("TODO")
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
