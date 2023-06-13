package goethclient

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/threefoldfoundation/tft/accountactivation/eth/contract"
	stellargoclient "github.com/threefoldtech/web3_proxy/server/clients/stellar"
)

const (
	contractAddress      = "0x42aa90c3c5f3B3428b68A22d1C85bB68b2a7E3D7"
	gasLimit             = 210000
	timeoutCreateAccount = 300
)

func GenerateKeypair() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}

func KeyFromSecret(secret string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(strings.TrimPrefix(secret, "0x"))
}

func (c *Client) AddressFromKey() common.Address {
	publicKey := c.Key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA)
}

func (c *Client) GetHexSeed() string {
	return hex.EncodeToString(crypto.FromECDSA(c.Key))
}

func (c *Client) CreateAndActivateStellarAccount(ctx context.Context, network string) (string, error) {
	kp, err := keypair.Random()
	if err != nil {
		return "", errors.Wrap(err, "failed to generate keypair")
	}

	// Fetch the price for activating an account on the Stellar network
	contractCaller, err := contract.NewAccountActivationCaller(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return "", errors.Wrap(err, "failed to create account activation caller")
	}
	cost, err := contractCaller.NetworkCost(&bind.CallOpts{
		Context: ctx,
	}, "stellar")
	if err != nil {
		return "", errors.Wrap(err, "failed to get network cost")
	}

	// Call the ActivateAccount function
	contractTransactor, err := contract.NewAccountActivationTransactor(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return "", errors.Wrap(err, "failed to create account activation transactor")
	}

	chainID, err := c.Eth.ChainID(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed getting the chain id")
	}

	pubKey, _ := c.Key.Public().(*ecdsa.PublicKey)
	ethereumAddress := crypto.PubkeyToAddress(*pubKey)
	_, err = contractTransactor.ActivateAccount(&bind.TransactOpts{
		Context: ctx,
		From:    ethereumAddress,
		Signer: func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(t, types.LatestSignerForChainID(chainID), c.Key)
		},
		Value:    cost,
		GasLimit: gasLimit,
	}, "stellar", kp.Address())
	if err != nil {
		return "", errors.Wrap(err, "failed to activate account")
	}

	client := stellargoclient.GetHorizonClient(network)
	for i := 0; i < int(timeoutCreateAccount); i++ {
		select {
		case <-time.After(1 * time.Second):
			accountRequest := horizonclient.AccountRequest{AccountID: kp.Address()}
			_, err := client.AccountDetail(accountRequest)
			if err == nil {
				return kp.Seed(), nil
			}
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}

	return "", errors.New("failed to wait on creation of stellar account")
}
