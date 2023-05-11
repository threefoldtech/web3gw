package goethclient

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	tft "github.com/threefoldfoundation/tft/bridge/stellar/contracts/tokenv1"
)

const (
	PublicEthTftContractAddress        = "0x395E925834996e558bdeC77CD648435d620AfB5b"
	GoerliTestnetEthTftContractAddress = "0xDa38782ce31Fc9861087320ABffBdee64Ed60515"
)

func (c *Client) TransferTftEth(ctx context.Context, destination string, amount int64) (string, error) {
	tftContractAddress, err := c.GetTftContractAddress(ctx)
	if err != nil {
		return "", err
	}
	return c.TransferTokens(ctx, tftContractAddress, destination, amount)
}

func (c *Client) WithdrawEthTftToStellar(ctx context.Context, destination string, amount int64) (string, error) {
	tftContractAddress, err := c.GetTftContractAddress(ctx)
	if err != nil {
		return "", err
	}
	tft, err := tft.NewToken(tftContractAddress, c.Eth)
	if err != nil {
		return "", err
	}

	opts, err := c.getDefaultTransactionOpts(ctx)
	if err != nil {
		return "", err
	}

	tx, err := tft.Withdraw(opts, big.NewInt(amount), destination, "stellar")
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (c *Client) GetTftBalance(ctx context.Context) (*big.Int, error) {
	tftContractAddress, err := c.GetTftContractAddress(ctx)
	if err != nil {
		return nil, err
	}

	tft, err := tft.NewToken(tftContractAddress, c.Eth)
	if err != nil {
		return nil, err
	}

	ctxWithCancel, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	return tft.BalanceOf(&bind.CallOpts{
		Context: ctxWithCancel,
	}, c.Address)
}

func (c *Client) GetTftContractAddress(ctx context.Context) (common.Address, error) {
	chainID, err := c.Eth.NetworkID(ctx)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "failed to get chainID")
	}

	if chainID.Cmp(big.NewInt(1)) == 0 {
		return common.HexToAddress(PublicEthTftContractAddress), nil
	} else if chainID.Cmp(big.NewInt(5)) == 0 {
		return common.HexToAddress(GoerliTestnetEthTftContractAddress), nil
	} else {
		return common.Address{}, errors.Errorf("unsupported chainID: %d", chainID)
	}
}
