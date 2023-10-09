package eth

import (
	"context"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/threefoldtech/3bot/web3gw/server/pkg"
)

func (c *Client) QuoteEthForTft(ctx context.Context, conState jsonrpc.State, amountIn string) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.QuoteEthForTft(ctx, amountIn)
}

func (c *Client) SwapEthForTft(ctx context.Context, conState jsonrpc.State, amountIn string) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.SwapEthForTft(ctx, amountIn)
}

func (c *Client) QuoteTftForEth(ctx context.Context, conState jsonrpc.State, amountIn string) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.QuoteTftForEth(ctx, amountIn)
}

func (c *Client) SwapTftForEth(ctx context.Context, conState jsonrpc.State, amountIn string) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.SwapTftForEth(ctx, amountIn)
}
