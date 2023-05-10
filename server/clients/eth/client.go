package goethclient

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

const (
	EthMainnetId = 1
	EthGoerliId  = 5
)

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
	nonce, err := c.Eth.PendingNonceAt(ctx, c.Address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get nonce")
	}

	gasPrice, err := c.Eth.SuggestGasPrice(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to suggest gas price")
	}

	chainID, err := c.Eth.NetworkID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chainID")
	}

	signerFn, addr, err := newSigner(c.Key, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create signer")
	}

	return &bind.TransactOpts{
		From:     addr,
		GasPrice: gasPrice,
		Signer:   signerFn,
		GasLimit: GasLimit,
		Nonce:    big.NewInt(int64(nonce)),
		Context:  ctx,
	}, nil
}

// newSigner creates a signer func using the flag-passed
// private credentials of the sender
func newSigner(privKey *ecdsa.PrivateKey, chainID *big.Int) (bind.SignerFn, common.Address, error) {
	keyAddr := crypto.PubkeyToAddress(privKey.PublicKey)
	return func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if address != keyAddr {
			return nil, errors.New("not authorized to sign this account")
		}
		s := types.NewEIP155Signer(chainID)
		signature, err := crypto.Sign(s.Hash(tx).Bytes(), privKey)
		if err != nil {
			return nil, err
		}
		return tx.WithSignature(s, signature)
	}, keyAddr, nil
}
