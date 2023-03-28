package eth

import (
	"context"
	"errors"

	goethclient "github.com/threefoldtech/web3_proxy/server/clients/eth"
	"github.com/threefoldtech/web3_proxy/server/pkg/state"
)

var (
	// ClientNotConnected indicates an ethereum client is not yet connected to an ethereum node and or the client does not have a private key loaded yet.
	ClientNotConnected = errors.New("client not connected yet")
)

type ethState struct {
	client *goethclient.Client
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

func (c *Client) Balance(ctx context.Context, address string) (int64, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return 0, ClientNotConnected
	}

	balance, err := state.client.GetBalance(address)
	if err != nil {
		return 0, err
	}

	return balance.Int64(), nil
}

func (c *Client) Load(ctx context.Context, url string, secret string) error {
	cl, err := goethclient.NewClient(url, secret)
	if err != nil {
		return err
	}

	es := ethState{
		client: cl,
	}

	c.state.Set(state.IDFromContext(ctx), es)

	return nil
}
