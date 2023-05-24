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

func (r *Client) Login(ctx context.Context, credentials Credentials) error {
	newClient, err := deployer.NewTFPluginClient(credentials.Mnemonics, "sr25519", credentials.Network, "", "", "", 10, true)
	if err != nil {
		return errors.Wrap(err, "failed to get tf plugin client")
	}

	r.client = NewTFGridClient(&newClient)
	r.TwinID = newClient.TwinID
	r.Identity = newClient.Identity

	return nil
}
