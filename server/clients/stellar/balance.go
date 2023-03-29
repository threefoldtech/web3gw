package stellargoclient

import (
	"math/big"

	"github.com/stellar/go/clients/horizonclient"
)

func (c *Client) GetBalance(account string) (*big.Int, error) {
	// Get information about the account we just created
	accountRequest := horizonclient.AccountRequest{AccountID: c.kp.Address()}
	hAccount, err := c.horizon.AccountDetail(accountRequest)
	if err != nil {
		return nil, err
	}

	balance := new(big.Int)
	for _, b := range hAccount.Balances {
		if b.Asset == c.GetTftBaseAsset() {
			balance.SetString(b.Balance, 10)
		}
	}

	return balance, nil
}
