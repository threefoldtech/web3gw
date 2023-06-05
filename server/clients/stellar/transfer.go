package stellargoclient

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

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
	stellarPublicNetworkTfchainBridgeAddress  = "GBNOTAYUMXVO5QDYWYO2SOCOYIJ3XFIP65GKOQN7H65ZZSO6BK4SLWSC"
	stellarTestnetNetworkTfchainBridgeAddress = "GDHJP6TF3UXYXTNEZ2P36J5FH7W4BJJQ4AYYAXC66I2Q2AH5B6O6BCFG"
)

func (c *Client) Transfer(destination, memo string, amount string) (string, error) {
	accountRequest := horizonclient.AccountRequest{AccountID: c.kp.Address()}
	hAccount, err := c.horizon.AccountDetail(accountRequest)
	if err != nil {
		return "", errors.Wrap(err, "account does not exist")
	}

	if !hasTrustline(hAccount, c.GetTftBaseAsset()) {
		return "", errors.New("source account does not have trustline")
	}

	destAccountRequest := horizonclient.AccountRequest{AccountID: destination}
	destHAccount, err := c.horizon.AccountDetail(destAccountRequest)
	if err != nil {
		return "", errors.Wrap(err, "account does not exist")
	}

	if !hasTrustline(destHAccount, c.GetTftBaseAsset()) {
		return "", errors.New("destination account does not have trustline")
	}

	transferTx := txnbuild.Payment{
		Destination: destination,
		Amount:      amount,
		Asset:       c.GetTftAsset(),
	}

	params := txnbuild.TransactionParams{
		SourceAccount:        &hAccount,
		IncrementSequenceNum: true,
		Operations:           []txnbuild.Operation{&transferTx},
		BaseFee:              BaseFee,
		Preconditions: txnbuild.Preconditions{
			TimeBounds: txnbuild.NewInfiniteTimeout(),
		},
	}
	if memo != "" {
		params.Memo = txnbuild.MemoText(memo)
	}

	tx, err := txnbuild.NewTransaction(params)
	if err != nil {
		return "", err
	}

	err = c.SignAndSubmit(tx)
	if err != nil {
		return "", err
	}
	hash, err := tx.HashHex(c.stellarNetwork)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (c *Client) TransferToEthBridge(destination, amount string) (string, error) {
	b, err := hex.DecodeString(strings.TrimPrefix(destination, "0x"))
	if err != nil {
		return "", err
	}

	bridgeAddr, err := c.GetEthBridgeAddress()
	if err != nil {
		return "", err
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

func (c *Client) TransferToTfchainBridge(amount string, twinID uint32) (string, error) {
	bridgeAddr, err := c.GetTfchainBridgeAddress()
	if err != nil {
		return "", err
	}
	memo := fmt.Sprintf("twin_%d", twinID)
	return c.Transfer(bridgeAddr, memo, amount)
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
	} else if c.stellarNetwork == "testnet" {
		return stellarTestnetNetworkTfchainBridgeAddress, nil
	} else {
		return "", errors.New("bsc address not available for networks other than public")
	}
}

func (c *Client) AwaitTransactionWithMemo(account string, memo string, timeout int) error {
	memo = strings.TrimPrefix(memo, "0x")
	for i := 0; i < int(timeout); i++ {
		transactionRequest := horizonclient.TransactionRequest{
			ForAccount: account,
			Order:      horizonclient.OrderDesc,
		}
		txs, err := c.horizon.Transactions(transactionRequest)
		if err != nil {
			return err
		}
		for _, tx := range txs.Embedded.Records {
			decodedMemo, err := base64.StdEncoding.DecodeString(tx.Memo)
			if err == nil {
				hexDecodedMemo := hex.EncodeToString(decodedMemo)
				if hexDecodedMemo == memo {
					return nil
				}
			}
		}
		time.Sleep(time.Second * 1)
	}
	return errors.New("transaction not found")
}

func (c *Client) AwaitTransactionWithMemoOnEthBridge(memo string, timeout int) error {
	bridgeAddress, err := c.GetEthBridgeAddress()
	if err != nil {
		return err
	}
	return c.AwaitTransactionWithMemo(bridgeAddress, memo, timeout)
}

func (c *Client) AwaitForTransactionWithMemoOnTfchainBridge(memo string, timeout int) error {
	bridgeAddress, err := c.GetTfchainBridgeAddress()
	if err != nil {
		return err
	}
	return c.AwaitTransactionWithMemo(bridgeAddress, memo, timeout)
}
