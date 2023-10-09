package goethclient

import "context"

func (c *Client) GetCurrentHeight() (uint64, error) {
	return c.Eth.BlockNumber(context.Background())
}
