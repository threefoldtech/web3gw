package goethclient

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type Client struct {
	Url     string
	Eth     *ethclient.Client
	Key     *ecdsa.PrivateKey
	Address common.Address
}

func NewClient(url, secret string) (*Client, error) {
	eth, err := ethclient.DialContext(context.Background(), url)
	if err != nil {
		return nil, err
	}

	cl := Client{
		Url: url,
		Eth: eth,
	}

	if secret == "" {
		kp, err := GenerateKeypair()
		if err != nil {
			return nil, err
		}
		cl.Key = kp
	} else {
		kp, err := KeyFromSecret(secret)
		if err != nil {
			return nil, errors.Wrap(err, "failed to import key")
		}
		cl.Key = kp
	}

	cl.Address = crypto.PubkeyToAddress(cl.Key.PublicKey)

	return &cl, nil
}

func (c *Client) getDefaultTransactionOpts(ctx context.Context) (*bind.TransactOpts, error) {
	nonce, err := c.Eth.PendingNonceAt(context.Background(), c.Address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get nonce")
	}

	gasLimit := uint64(21000)
	gasPrice, err := c.Eth.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to suggest gas price")
	}

	ctxWithCancel, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	return &bind.TransactOpts{
		GasPrice: gasPrice,
		GasLimit: gasLimit,
		Nonce:    big.NewInt(int64(nonce)),
		Context:  ctxWithCancel,
	}, nil
}
