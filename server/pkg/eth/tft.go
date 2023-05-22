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
		Amount      string `json:"amount"`
	}
)

func (c *Client) TransferEthTft(ctx context.Context, conState jsonrpc.State, args TftEthTransfer) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.TransferEthTft(ctx, args.Destination, args.Amount)
}

func (c *Client) WithdrawEthTftToStellar(ctx context.Context, conState jsonrpc.State, args TftEthTransfer) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.WithdrawEthTftToStellar(ctx, args.Destination, args.Amount)
}

func (c *Client) GetEthTftBalance(ctx context.Context, conState jsonrpc.State) (*big.Int, error) {
	state := State(conState)
	if state.Client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.Client.GetEthTftBalance(ctx)
}

func (c *Client) ApproveEthTftSpending(ctx context.Context, conState jsonrpc.State, amount string) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.ApproveEthTftSpending(ctx, amount)
}
