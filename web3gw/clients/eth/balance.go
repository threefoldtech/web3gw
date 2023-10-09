package goethclient

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) GetBalance(addr string) (string, error) {
	address := common.HexToAddress(addr)

	currentBlock, err := c.Eth.BlockNumber(context.Background())
	if err != nil {
		return "", err
	}

	b, err := c.Eth.BalanceAt(context.Background(), address, big.NewInt(int64(currentBlock)))
	if err != nil {
		return "", err
	}

	return WeiToString(b), nil
}
