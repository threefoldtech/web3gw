package goethclient

import (
	"context"
	"math/big"
	"time"

	"github.com/daoleno/uniswapv3-sdk/examples/helper"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	tft "github.com/threefoldfoundation/tft/bridge/stellar/contracts/tokenv1"
)

var (
	MainnetEthTftContractAddress = common.HexToAddress("0x395E925834996e558bdeC77CD648435d620AfB5b")
	GoerliEthTftContractAddress  = common.HexToAddress("0xDa38782ce31Fc9861087320ABffBdee64Ed60515")
)

func (c *Client) TransferEthTft(ctx context.Context, destination string, amount int64) (string, error) {
	return c.TransferTokens(ctx, MainnetEthTftContractAddress, destination, amount)
}

func (c *Client) WithdrawEthTftToStellar(ctx context.Context, destination string, amount int64) (string, error) {
	tft, err := tft.NewToken(MainnetEthTftContractAddress, c.Eth)
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

	r, err := bind.WaitMined(ctx, c.Eth, tx)
	if err != nil {
		log.Err(err).Msg("failed to wait for tft approval")
		return "", err
	}

	log.Debug().Msgf("Withdraw tx mined: %s, block %d, gas: %d, status: %d", tx.Hash().Hex(), r.BlockNumber, r.GasUsed, r.Status)

	return tx.Hash().Hex(), nil
}

func (c *Client) GetEthTftBalance(ctx context.Context) (*big.Int, error) {
	tftC, err := c.GetTftTokenContract()
	if err != nil {
		return nil, err
	}

	tft, err := tft.NewToken(tftC.Address, c.Eth)
	if err != nil {
		return nil, err
	}

	ctxWithCancel, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	return tft.BalanceOf(&bind.CallOpts{
		Context: ctxWithCancel,
	}, c.Address)
}

func (c *Client) ApproveEthTftSpending(ctx context.Context, input string) (string, error) {
	tftC, err := c.GetTftTokenContract()
	if err != nil {
		return "", err
	}

	tft, err := tft.NewToken(tftC.Address, c.Eth)
	if err != nil {
		return "", err
	}

	ctxWithCancel, cancel := context.WithTimeout(ctx, time.Minute*1)
	defer cancel()

	opts, err := c.getDefaultTransactionOpts(ctxWithCancel)
	if err != nil {
		return "", err
	}

	amount := helper.FloatStringToBigInt(input, int(tftC.Decimals()))
	tx, err := tft.Approve(opts, SwapRouter, amount)
	if err != nil {
		log.Err(err).Msg("failed to approve tft spending")
		return "", err
	}

	r, err := bind.WaitMined(ctxWithCancel, c.Eth, tx)
	if err != nil {
		log.Err(err).Msg("failed to wait for tft approval")
		return "", err
	}

	log.Debug().Msgf("Approve spend tx mined: %s, block %d, gas: %d, status: %d", tx.Hash().Hex(), r.BlockNumber, r.GasUsed, r.Status)

	return tx.Hash().Hex(), nil
}
