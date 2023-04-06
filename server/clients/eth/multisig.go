package goethclient

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/threefoldtech/web3_proxy/server/clients/eth/gnosis"
)

func (c *Client) GetMultisigVersion(address string) (string, error) {
	ms, err := gnosis.NewGnosis(common.HexToAddress(address), c.Eth)
	if err != nil {
		return "", err
	}

	return ms.VERSION(&bind.CallOpts{})
}

func (c *Client) GetThreshold(address string) (*big.Int, error) {
	ms, err := gnosis.NewGnosis(common.HexToAddress(address), c.Eth)
	if err != nil {
		return nil, err
	}

	return ms.GetThreshold(&bind.CallOpts{})
}
