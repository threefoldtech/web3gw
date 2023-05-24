package atomicswap

import (
	"context"

	"github.com/pkg/errors"
	"github.com/stellar/go/txnbuild"
	stellargoclient "github.com/threefoldtech/web3_proxy/server/clients/stellar"
)

type (
	// StellarDriver implements stellar specific atomic swap logic
	StellarDriver struct {
		stellar *stellargoclient.Client
	}
)

// InitTFTTransfer implements SellChain
func (s *StellarDriver) InitTFTTransfer(ctx context.Context, details NegotiatedTrade, sharedSecret SwapSecretHash) (any, error) {
	return nil, errors.New("TODO")
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
