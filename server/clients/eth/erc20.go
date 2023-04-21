package goethclient

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/threefoldtech/web3_proxy/server/clients/eth/erc20"
)

func (c *Client) GetTokenBalance(contractAddress, target string) (*big.Int, error) {
	token, err := erc20.NewErc20(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return nil, err
	}

	return token.BalanceOf(&bind.CallOpts{}, common.HexToAddress(target))
}

func (c *Client) TransferTokens(contractAddress, target string, amount int64) (string, error) {
	token, err := erc20.NewErc20(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return "", err
	}

	tx, err := token.Transfer(&bind.TransactOpts{}, common.HexToAddress(target), big.NewInt(amount))
	if err != nil {
		return "", err
	}

	return c.sendTransaction(tx)
}

func (c *Client) ApproveTokenSpending(contractAddress, spender string, amount int64) (string, error) {
	token, err := erc20.NewErc20(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return "", err
	}

	tx, err := token.Approve(&bind.TransactOpts{}, common.HexToAddress(spender), big.NewInt(amount))
	if err != nil {
		return "", err
	}

	return c.sendTransaction(tx)
}

func (c *Client) TransferFromTokens(contractAddress, from, to string, amount int64) (string, error) {
	token, err := erc20.NewErc20(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return "", err
	}

	tx, err := token.TransferFrom(&bind.TransactOpts{}, common.HexToAddress(from), common.HexToAddress(to), big.NewInt(amount))
	if err != nil {
		return "", err
	}

	return c.sendTransaction(tx)
}
