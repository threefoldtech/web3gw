package stellargoclient

func (c *Client) GetBalance(account string) (string, error) {
	if account == "" {
		account = c.kp.Address()
	}
	hAccount, err := c.AccountData(account)
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
