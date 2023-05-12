package stellargoclient

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/support/errors"
	"github.com/stellar/go/txnbuild"
)

const (
	// Eth
	stellarPublicNetworkEthBridgeAddress  = "GARQ6KUXUCKDPIGI7NPITDN55J23SVR5RJ5RFOOU3ZPLMRJYOQRNMOIJ"
	stellarTestnetNetworkEthBridgeAddress = "GAXPJGADXTP2FXUYASUOE5MQ6SSCEMBU2PPD27ZG55MKKPJRAVASBNJI"
	// BSC
	// stellarPublicNetworkBscBridgeAddress = "GBFFWXWBZDILJJAMSINHPJEUJKB3H4UYXRWNB4COYQAF7UUQSWSBUXW5"
	// Tfchain
	// stellarPublicNetworkTfchainBridgeAddress = "GBNOTAYUMXVO5QDYWYO2SOCOYIJ3XFIP65GKOQN7H65ZZSO6BK4SLWSC"
)

func (c *Client) Transfer(destination, memo string, amount string) error {
	accountRequest := horizonclient.AccountRequest{AccountID: c.kp.Address()}
	hAccount, err := c.horizon.AccountDetail(accountRequest)
	if err != nil {
		return errors.Wrap(err, "account does not exist")
	}

	if !hasTftTrustline(hAccount) {
		return errors.New("source account does not have trustline")
	}

	destAccountRequest := horizonclient.AccountRequest{AccountID: destination}
	destHAccount, err := c.horizon.AccountDetail(destAccountRequest)
	if err != nil {
		return errors.Wrap(err, "account does not exist")
	}

	if !hasTftTrustline(destHAccount) {
		return errors.New("destination account does not have trustline")
	}

	transferTx := txnbuild.Payment{
		Destination:   destination,
		Amount:        amount,
		Asset:         c.GetTftAsset(),
		SourceAccount: c.kp.Address(),
	}

	params := txnbuild.TransactionParams{
		SourceAccount:        &hAccount,
		IncrementSequenceNum: true,
		Operations:           []txnbuild.Operation{&transferTx},
		BaseFee:              txnbuild.MinBaseFee,
		Memo:                 txnbuild.MemoText(memo),
		Preconditions: txnbuild.Preconditions{
			TimeBounds: txnbuild.NewInfiniteTimeout(),
		},
	}
	tx, err := txnbuild.NewTransaction(params)
	if err != nil {
		return err
	}

	return c.SignAndSubmit(tx)
}

func (c *Client) TransferToEthBridge(destination, amount string) error {
	b, err := hex.DecodeString(strings.TrimPrefix(destination, "0x"))
	if err != nil {
		return err
	}

	bridgeAddr, err := c.GetEthBridgeAddress()
	if err != nil {
		return err
	}

	return c.Transfer(bridgeAddr, fmt.Sprintf("%s=", base64.RawStdEncoding.EncodeToString(b)), amount)
}

func (c *Client) GetEthBridgeAddress() (string, error) {
	if c.stellarNetwork == "public" {
		return stellarPublicNetworkEthBridgeAddress, nil
	} else if c.stellarNetwork == "testnet" {
		return stellarTestnetNetworkEthBridgeAddress, nil
	} else {
		return "", errors.New("eth bridge address not available for networks other than public")
	}
}

// Reinstate later

// func (c *Client) TransferToBscBridge(destination, amount string) error {
// 	b, err := hex.DecodeString(strings.TrimPrefix(destination, "0x"))
// 	if err != nil {
// 		return err
// 	}

// 	bridgeAddr, err := c.GetEthBridgeAddress()
// 	if err != nil {
// 		return err
// 	}

// 	return c.Transfer(bridgeAddr, base64.RawStdEncoding.EncodeToString(b), amount)
// }

func (c *Client) TransferToTfchainBridge(destination, amount string, twinID uint32) error {
	bridgeAddr, err := c.GetTfchainBridgeAddress()
	if err != nil {
		return err
	}

	return c.Transfer(bridgeAddr, fmt.Sprintf("twin_%d", twinID), amount)
}

// func (c *Client) GetBscBridgeAddress() (string, error) {
// 	if c.stellarNetwork == "public" {
// 		return stellarPublicNetworkBscBridgeAddress, nil
// 	} else {
// 		return "", errors.New("bsc address not available for networks other than public")
// 	}
// }

func (c *Client) GetTfchainBridgeAddress() (string, error) {
	if c.stellarNetwork == "public" {
		return stellarPublicNetworkTfchainBridgeAddress, nil
	} else {
		return "", errors.New("bsc address not available for networks other than public")
	}
}
