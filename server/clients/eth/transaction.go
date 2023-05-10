package goethclient

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (c *Client) sendTransaction(tx *types.Transaction) (string, error) {
	chainID, err := c.Eth.NetworkID(context.Background())
	if err != nil {
		return "", errors.Wrap(err, "failed to get chainID")
	}

	log.Debug().Msg("signing tx")
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), c.Key)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign tx")
	}

	err = c.Eth.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", errors.Wrap(err, "failed to send transaction")
	}

	return signedTx.Hash().Hex(), nil
}
