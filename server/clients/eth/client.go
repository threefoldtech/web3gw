package goethclient

import (
	"context"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type Client struct {
	Url string
	Eth *ethclient.Client
	Key *ecdsa.PrivateKey
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

	return &cl, nil
}
