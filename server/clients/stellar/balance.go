package stellargoclient

import (
	"github.com/stellar/go/clients/horizonclient"
)

func (c *Client) GetBalance(account string) (string, error) {
	// Get information about the account we just created
	accountRequest := horizonclient.AccountRequest{AccountID: c.kp.Address()}
	hAccount, err := c.horizon.AccountDetail(accountRequest)
	if err != nil {
		return "", err
	}

	for _, b := range hAccount.Balances {
		if b.Asset == c.GetTftBaseAsset() {
			return b.Balance, nil
		}
	}

	return "", nil
}
