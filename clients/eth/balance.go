package goethclient

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) GetBalance(addr string) (*big.Int, error) {
	address := common.HexToAddress(addr)

	currentBlock, err := c.Eth.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	return c.Eth.BalanceAt(context.Background(), address, big.NewInt(int64(currentBlock)))
}
