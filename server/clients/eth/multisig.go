package goethclient

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/threefoldtech/web3_proxy/server/clients/eth/gnosis"
)

func (c *Client) GetMultisigVersion(contractAddress string) (string, error) {
	ms, err := gnosis.NewGnosis(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return "", err
	}

	return ms.VERSION(&bind.CallOpts{})
}

func (c *Client) GetOwners(contractAddress string) ([]string, error) {
	log.Info("getting owners for contract", contractAddress)
	ms, err := gnosis.NewGnosis(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return nil, err
	}

	owners, err := ms.GetOwners(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	ownersAsHexStr := make([]string, 0)
	for _, owner := range owners {
		ownersAsHexStr = append(ownersAsHexStr, owner.Hex())
	}
	return ownersAsHexStr, nil
}

func (c *Client) IsOwner(contractAddress string, address string) (bool, error) {
	ms, err := gnosis.NewGnosis(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return false, err
	}

	return ms.IsOwner(&bind.CallOpts{}, common.HexToAddress(address))
}

func (c *Client) GetThreshold(contractAddress string) (*big.Int, error) {
	ms, err := gnosis.NewGnosis(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return nil, err
	}

	return ms.GetThreshold(&bind.CallOpts{})
}

func (c *Client) AddOwner(contractAddress, target string, treshold int64) (string, error) {
	ms, err := gnosis.NewGnosis(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return "", err
	}

	isOwner, err := ms.IsOwner(&bind.CallOpts{}, common.HexToAddress(target))
	if err != nil {
		return "", err
	}

	if isOwner {
		return "", errors.New("target is already owner")
	}

	tx, err := ms.AddOwnerWithThreshold(&bind.TransactOpts{}, common.HexToAddress(target), big.NewInt(treshold))
	if err != nil {
		return "", err
	}

	return c.sendTransaction(tx)
}

func (c *Client) RemoveOwner(contractAddress, target string, treshold int64) (string, error) {
	ms, err := gnosis.NewGnosis(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return "", err
	}

	isOwner, err := ms.IsOwner(&bind.CallOpts{}, common.HexToAddress(target))
	if err != nil {
		return "", err
	}

	if !isOwner {
		return "", errors.New("target is not an owner")
	}

	tx, err := ms.RemoveOwner(&bind.TransactOpts{}, common.HexToAddress(target), common.HexToAddress(target), big.NewInt(treshold))
	if err != nil {
		return "", err
	}

	return c.sendTransaction(tx)
}

func (c *Client) ApproveHash(contractAddress, hex string) (string, error) {
	ms, err := gnosis.NewGnosis(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return "", err
	}

	tx, err := ms.ApproveHash(&bind.TransactOpts{}, common.HexToHash(hex))
	if err != nil {
		return "", err
	}

	return c.sendTransaction(tx)
}
