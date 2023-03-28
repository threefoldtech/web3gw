package stellargoclient

import (
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
)

var TestnetTft = txnbuild.CreditAsset{Code: TFT, Issuer: TESTNET_ISSUER}
var MainnetTft = txnbuild.CreditAsset{Code: TFT, Issuer: TESTNET_ISSUER}

var TestnetTftAsset = base.Asset{Type: "credit_alphanum4", Code: TFT, Issuer: TESTNET_ISSUER}
var MainnetTftAsset = base.Asset{Type: "credit_alphanum4", Code: TFT, Issuer: TESTNET_ISSUER}

// hasTftTrustline checks if the account has a trustline for the TFT asset
func hasTftTrustline(hAccount horizon.Account) bool {
	hasTftTrustline := false
	for _, b := range hAccount.Balances {
		if b.Asset == TestnetTftAsset {
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
