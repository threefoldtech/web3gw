package goethclient

import (
	"context"
	"crypto/ecdsa"

	"github.com/daoleno/uniswapv3-sdk/examples/helper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

const EthDecimals = 18

func (c *Client) TransferEth(ctx context.Context, amount string, destination string) (string, error) {
	tx, err := c.createTransferTransaction(amount, destination)
	if err != nil {
		return "", errors.Wrap(err, "failed to create transfer transaction")
	}

	return c.sendTransaction(ctx, tx)
}

func (c *Client) createTransferTransaction(amount string, destination string) (*types.Transaction, error) {
	publicKey := c.Key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := c.Eth.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get nonce")
	}

	gasPrice, err := c.Eth.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to suggest gas price")
	}

	toAddress := common.HexToAddress(destination)

	amountIn := helper.FloatStringToBigInt(amount, EthDecimals)
	tx := types.NewTransaction(nonce, toAddress, amountIn, GasLimit, gasPrice, nil)

	return tx, nil
}
