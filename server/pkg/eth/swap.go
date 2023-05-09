package eth

import (
	"context"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/threefoldtech/web3_proxy/server/pkg"
)

func (c *Client) QuoteTftEth(ctx context.Context, conState jsonrpc.State, amountIn string) (int64, error) {
	state := State(conState)
	if state.Client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.Client.QuoteTftEth(ctx, amountIn)
}

func (c *Client) SwapTftEth(ctx context.Context, conState jsonrpc.State, amountIn string) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.SwapTftEth(ctx, amountIn)
}
