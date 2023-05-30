package tfgrid

import (
	"context"

	"github.com/pkg/errors"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/deployer"
)

type Credentials struct {
	Mnemonics string `json:"mnemonics"`
	Network   string `json:"network"`
}

func (c *Client) Login(ctx context.Context, credentials Credentials) error {
	newClient, err := deployer.NewTFPluginClient(credentials.Mnemonics, "sr25519", credentials.Network, "", "", "", 10, true)
	if err != nil {
		return errors.Wrap(err, "failed to get tf plugin client")
	}

	c.GridClient = NewTFGridClient(&newClient)
	c.TwinID = newClient.TwinID
	c.Identity = newClient.Identity

	return nil
}

func (c *Client) Logout(ctx context.Context) {
	c.GridClient.Close()
}
