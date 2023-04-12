package goethclient

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

func (c *Client) TransferEth(amount int64, destination string) (string, error) {
	publicKey := c.Key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := c.Eth.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", errors.Wrap(err, "failed to get nonce")
	}

	value := big.NewInt(amount)
	gasLimit := uint64(21000)
	gasPrice, err := c.Eth.SuggestGasPrice(context.Background())
	if err != nil {
		return "", errors.Wrap(err, "failed to suggest gas price")
	}

	toAddress := common.HexToAddress(destination)

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID, err := c.Eth.NetworkID(context.Background())
	if err != nil {
		return "", errors.Wrap(err, "failed to get chainID")
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), c.Key)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign tx")
	}

	err = c.Eth.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", errors.Wrap(err, "failed to send transaction")
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	return signedTx.Hash().Hex(), nil
}