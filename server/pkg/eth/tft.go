package eth

import (
	"context"
	"math/big"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/threefoldtech/web3_proxy/server/pkg"
)

type (
	TftEthTransfer struct {
		Destination string `json:"destination"`
		Amount      int64  `json:"amount"`
	}
)

func (c *Client) TransferTftEth(ctx context.Context, conState jsonrpc.State, args TftEthTransfer) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.TransferTftEth(ctx, args.Destination, args.Amount)
}

func (c *Client) WithdrawEthTftToStellar(ctx context.Context, conState jsonrpc.State, args TftEthTransfer) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.WithdrawEthTftToStellar(ctx, args.Destination, args.Amount)
}

func (c *Client) GetTftBalance(ctx context.Context, conState jsonrpc.State) (*big.Int, error) {
	state := State(conState)
	if state.Client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.Client.GetTftBalance(ctx)
}
