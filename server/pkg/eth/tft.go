package eth

import (
	"context"

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

func (c *Client) BridgeToStellar(ctx context.Context, conState jsonrpc.State, args TftEthTransfer) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.BridgeToStellar(ctx, args.Destination, args.Amount)
}

func (c *Client) GetEthTftBalance(ctx context.Context, conState jsonrpc.State) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	tftBalance, err := state.Client.GetEthTftBalance(ctx)
	if err != nil {
		return "", err
	}

	return tftBalance.String(), nil
}

func (c *Client) ApproveEthTftSpending(ctx context.Context, conState jsonrpc.State, amount string) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.ApproveEthTftSpending(ctx, amount)
}

func (c *Client) EthTftSpendingAllowance(ctx context.Context, conState jsonrpc.State) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.EthTftSpendingAllowance(ctx)
}
