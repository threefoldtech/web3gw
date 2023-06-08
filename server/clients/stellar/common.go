package stellargoclient

import (
	"errors"
	"strings"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/protocols/horizon"
	"github.com/stellar/go/protocols/horizon/base"
	"github.com/stellar/go/txnbuild"
)

const (
	TFT            = "TFT"
	TESTNET_ISSUER = "GA47YZA3PKFUZMPLQ3B5F2E3CJIB57TGGU7SPCQT2WAEYKN766PWIMB3"
	MAINNET_ISSUER = "GBOVQKJYHXRR3DX6NOX2RRYFRCUMSADGDESTDNBDS6CDVLGVESRTAC47"
	BaseFee        = 1000000
)

var TestnetTft = txnbuild.CreditAsset{Code: TFT, Issuer: TESTNET_ISSUER}
var MainnetTft = txnbuild.CreditAsset{Code: TFT, Issuer: MAINNET_ISSUER}

var TestnetTftAsset = base.Asset{Type: "credit_alphanum4", Code: TFT, Issuer: TESTNET_ISSUER}
var MainnetTftAsset = base.Asset{Type: "credit_alphanum4", Code: TFT, Issuer: MAINNET_ISSUER}

// hasTrustline checks if the account has a trustline a specific asset
func hasTrustline(hAccount horizon.Account, asset base.Asset) bool {
	hasTftTrustline := false
	for _, b := range hAccount.Balances {
		if b.Asset == asset {
			hasTftTrustline = true
			break
		}
	}

	return hasTftTrustline
}

// GetHorizonClient returns the horizon client for the stellar network
func GetHorizonClient(stellarNetwork string) *horizonclient.Client {
	if stellarNetwork == "testnet" {
		return horizonclient.DefaultTestNetClient
	} else if stellarNetwork == "public" {
		return horizonclient.DefaultPublicNetClient
	} else {
		return horizonclient.DefaultTestNetClient
	}
}

// GetTftAsset returns the tft asset for the stellar network
func (c *Client) GetTftAsset() txnbuild.CreditAsset {
	if c.stellarNetwork == "testnet" {
		return TestnetTft
	} else if c.stellarNetwork == "public" {
		return MainnetTft
	} else {
		return TestnetTft
	}
}

// GetTftAsset returns the tft asset for the stellar network
func (c *Client) GetTftBaseAsset() base.Asset {
	if c.stellarNetwork == "testnet" {
		return TestnetTftAsset
	} else if c.stellarNetwork == "public" {
		return MainnetTftAsset
	} else {
		return TestnetTftAsset
	}
}

func (c *Client) GetXlmAsset() txnbuild.CreditAsset {
	return txnbuild.CreditAsset{}
}

// GetTftAsset returns the tft asset for the stellar network
func (c *Client) GetAssetFromString(asset string) (txnbuild.Asset, error) {
	assetLower := strings.ToLower(asset)
	if assetLower == "tft" {
		return c.GetTftAsset(), nil
	} else if assetLower == "xlm" {
		return txnbuild.NativeAsset{}, nil
	} else {
		return txnbuild.CreditAsset{}, errors.New("unsupported asset")
	}
}

// GetStellarNetworkPassphrase returns the passphrase for the stellar network
func (c *Client) GetStellarNetworkPassphrase() string {
	if c.stellarNetwork == "testnet" {
		return network.TestNetworkPassphrase
	} else if c.stellarNetwork == "public" {
		return network.PublicNetworkPassphrase
	} else {
		return network.TestNetworkPassphrase
	}
}

func GetKeypairFromSeed(seed string) (*keypair.Full, error) {
	kp, err := keypair.Parse(seed)
	if err != nil {
		return nil, err
	}

	return kp.(*keypair.Full), nil
}
