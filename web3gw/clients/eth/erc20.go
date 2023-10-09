package goethclient

import (
	"context"

	"github.com/daoleno/uniswapv3-sdk/examples/helper"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/threefoldtech/3bot/web3gw/server/clients/eth/erc20"
)

func (c *Client) GetTokenBalance(contractAddress string) (string, error) {
	token, err := erc20.NewErc20(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return "", err
	}

	b, err := token.BalanceOf(&bind.CallOpts{}, c.Address)
	if err != nil {
		return "", err
	}

	return WeiToString(b), nil
}

func (c *Client) TransferTokens(ctx context.Context, contractAddress common.Address, target string, amount string) (string, error) {
	token, err := erc20.NewErc20(contractAddress, c.Eth)
	if err != nil {
		return "", err
	}

	opts, err := c.getDefaultTransactionOpts(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed to get default transaction opts")
	}

	amountIn := helper.FloatStringToBigInt(amount, EthDecimals)
	tx, err := token.Transfer(opts, common.HexToAddress(target), amountIn)
	if err != nil {
		return "", err
	}

	r, err := bind.WaitMined(ctx, c.Eth, tx)
	if err != nil {
		log.Err(err).Msg("failed to wait for tft approval")
		return "", err
	}

	log.Debug().Msgf("Approve spend tx mined: %s, block %d, gas: %d, status: %d", tx.Hash().Hex(), r.BlockNumber, r.GasUsed, r.Status)

	return tx.Hash().Hex(), nil
}

func (c *Client) ApproveTokenSpending(ctx context.Context, contractAddress, spender string, amount string) (string, error) {
	token, err := erc20.NewErc20(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return "", err
	}

	opts, err := c.getDefaultTransactionOpts(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed to get default transaction opts")
	}

	amountIn := helper.FloatStringToBigInt(amount, EthDecimals)
	tx, err := token.Approve(opts, common.HexToAddress(spender), amountIn)
	if err != nil {
		return "", err
	}

	r, err := bind.WaitMined(ctx, c.Eth, tx)
	if err != nil {
		log.Err(err).Msg("failed to wait for tft approval")
		return "", err
	}

	log.Debug().Msgf("Approve spend tx mined: %s, block %d, gas: %d, status: %d", tx.Hash().Hex(), r.BlockNumber, r.GasUsed, r.Status)

	return tx.Hash().Hex(), nil
}

func (c *Client) TransferFromTokens(ctx context.Context, contractAddress, from, to string, amount string) (string, error) {
	token, err := erc20.NewErc20(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return "", err
	}

	opts, err := c.getDefaultTransactionOpts(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed to get default transaction opts")
	}

	amountIn := helper.FloatStringToBigInt(amount, EthDecimals)
	tx, err := token.TransferFrom(opts, common.HexToAddress(from), common.HexToAddress(to), amountIn)
	if err != nil {
		return "", err
	}

	r, err := bind.WaitMined(ctx, c.Eth, tx)
	if err != nil {
		log.Err(err).Msg("failed to wait for tft approval")
		return "", err
	}

	log.Debug().Msgf("Approve spend tx mined: %s, block %d, gas: %d, status: %d", tx.Hash().Hex(), r.BlockNumber, r.GasUsed, r.Status)

	return tx.Hash().Hex(), nil
}
