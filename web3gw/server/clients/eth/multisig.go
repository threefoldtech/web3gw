package goethclient

import (
	"math/big"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/threefoldtech/web3_proxy/server/clients/eth/erc20"
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

	return tx.Hash().Hex(), nil
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

	return tx.Hash().Hex(), nil
}

func (c *Client) ApproveHash(contractAddress, hash string) (string, error) {
	ms, err := gnosis.NewGnosis(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return "", err
	}

	tx, err := ms.ApproveHash(&bind.TransactOpts{}, common.HexToHash(hash))
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (c *Client) IsApproved(contractAddress string, hash string) (bool, error) {
	ms, err := gnosis.NewGnosis(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return false, err
	}

	h, err := ms.ApprovedHashes(&bind.CallOpts{}, c.AddressFromKey(), common.HexToHash(hash))
	if err != nil {
		return false, err
	}

	return h.Int64() == 1, nil
}

func (c *Client) InitiateMultisigEthTransfer(safeContractAddress, destination string, amount string) (string, error) {
	ms, err := gnosis.NewGnosis(common.HexToAddress(safeContractAddress), c.Eth)
	if err != nil {
		return "", err
	}

	tx, err := c.createTransferTransaction(amount, destination)
	if err != nil {
		return "", errors.Wrap(err, "failed to create transfer transaction")
	}

	msTx, err := ms.ExecTransaction(
		&bind.TransactOpts{},
		// toAddress,
		*tx.To(),
		// value
		tx.Value(),
		// data
		tx.Data(),
		// operation (optional)
		0,
		// safeTxGas (optional)
		nil,
		// baseGas (optional)
		tx.GasPrice(),
		// gasPrice (optional)
		tx.GasPrice(),
		// gasToken (optional)
		common.HexToAddress("0x0000000000000000000000000000000000000000"),
		// refundReceiver (optional)
		c.AddressFromKey(),
		// signatures
		[]byte{},
	)

	if err != nil {
		return "", err
	}

	return msTx.Hash().Hex(), nil
}

func (c *Client) InitiateMultisigTokenTransfer(safeContractAddress, tokenAddress, destination string, amount int64) (string, error) {
	ms, err := gnosis.NewGnosis(common.HexToAddress(safeContractAddress), c.Eth)
	if err != nil {
		return "", err
	}

	token, err := erc20.NewErc20(common.HexToAddress(tokenAddress), c.Eth)
	if err != nil {
		return "", err
	}

	tx, err := token.Transfer(&bind.TransactOpts{}, common.HexToAddress(destination), big.NewInt(amount))
	if err != nil {
		return "", err
	}

	msTx, err := ms.ExecTransaction(
		&bind.TransactOpts{},
		// toAddress,
		*tx.To(),
		// value
		tx.Value(),
		// data
		tx.Data(),
		// operation (optional)
		0,
		// safeTxGas (optional)
		nil,
		// baseGas (optional)
		tx.GasPrice(),
		// gasPrice (optional)
		tx.GasPrice(),
		// gasToken (optional)
		common.HexToAddress("0x0000000000000000000000000000000000000000"),
		// refundReceiver (optional)
		c.AddressFromKey(),
		// signatures
		[]byte{},
	)

	if err != nil {
		return "", err
	}

	return msTx.Hash().Hex(), nil
}
