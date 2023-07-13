package goethclient

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (c *Client) sendTransaction(ctx context.Context, tx *types.Transaction) (string, error) {
	chainID, err := c.Eth.NetworkID(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed to get chainID")
	}

	log.Debug().Msg("signing tx")
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), c.Key)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign tx")
	}

	err = c.Eth.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", errors.Wrap(err, "failed to send transaction")
	}

	res, err := bind.WaitMined(ctx, c.Eth, signedTx)
	if err != nil {
		return "", err
	}

	log.Debug().Msgf("Swap tx mined: %s, block %d, gas: %d, status: %d", tx.Hash().Hex(), res.BlockNumber, res.GasUsed, res.Status)

	return tx.Hash().Hex(), nil
}
