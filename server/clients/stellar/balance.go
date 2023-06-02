package stellargoclient

import (
	"math/big"

	"github.com/stellar/go/clients/horizonclient"
)

func (c *Client) GetBalance(account string) (*big.Int, error) {
	if account == "" {
		account = c.kp.Address()
	}
	accountRequest := horizonclient.AccountRequest{AccountID: account}
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
