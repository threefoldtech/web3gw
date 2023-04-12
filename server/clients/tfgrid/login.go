package tfgrid

import (
	"context"

	"github.com/pkg/errors"
	"github.com/threefoldtech/grid3-go/deployer"
)

type Credentials struct {
	Mnemonics string `json:"mnemonics"`
	Network   string `json:"network"`
}

func (c *Runner) Login(ctx context.Context, credentials Credentials) error {
	newClient, err := deployer.NewTFPluginClient(credentials.Mnemonics, "sr25519", credentials.Network, "", "", "", 10, true, false)
	if err != nil {
		return errors.Wrap(err, "failed to get tf plugin client")
	}

	c.client = NewTFGridClient(&newClient)

	return nil
}
