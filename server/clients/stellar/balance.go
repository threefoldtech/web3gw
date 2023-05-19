package stellargoclient

import (
	"math/big"

	"github.com/rs/zerolog/log"
	"github.com/stellar/go/clients/horizonclient"
)

func (c *Client) GetBalance(account string) (*big.Int, error) {
	// Get information about the account we just created
	log.Debug().Msgf("%s vs %s", account, c.kp.Address())
	accountRequest := horizonclient.AccountRequest{AccountID: c.kp.Address()}
	hAccount, err := c.horizon.AccountDetail(accountRequest)
	if err != nil {
		return nil, err
	}

	balance := new(big.Int)
	for _, b := range hAccount.Balances {
		log.Debug().Msgf("%s: %s", b.Asset.Code, b.Balance)

		if b.Asset == c.GetTftBaseAsset() {
			balance.SetString(b.Balance, 10)
		}
	}

	return balance, nil
}
