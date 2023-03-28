package eth

import (
	"context"
	"crypto/ecdsa"

	goethclient "github.com/threefoldtech/web3_proxy/server/clients/eth"
	"github.com/threefoldtech/web3_proxy/server/pkg/state"
)

type ethState struct {
	keypair *ecdsa.PrivateKey
}

// NewClient creates a new Client ready for use
func NewClient() *Client {
	return &Client{
		state: state.NewStateManager[ethState](),
	}
}

// Client exposes ethereum related functionality
type Client struct {
	state *state.StateManager[ethState]
}

func (c *Client) Load(ctx context.Context, secret string) error {
	key, err := goethclient.KeyFromSecret(secret)
	if err != nil {
		return err
	}

	es := ethState{
		keypair: key,
	}

	c.state.Set(state.IDFromContext(ctx), es)

	return nil
}
