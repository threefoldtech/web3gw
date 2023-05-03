package eth

import (
	"context"
	"math/big"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/threefoldtech/web3_proxy/server/pkg"
)

type (
	GetTokenBalance struct {
		ContractAddress string `json:"contract_address"`
		Target          string `json:"target"`
	}

	TokenTransfer struct {
		ContractAddress string `json:"contract_address"`
		Destination     string `json:"destination"`
		Amount          int64  `json:"amount"`
	}

	TokenTransferFrom struct {
		ContractAddress string `json:"contract_address"`
		From            string `json:"from"`
		Destination     string `json:"destination"`
		Amount          int64  `json:"amount"`
	}

	ApproveTokenSpending struct {
		ContractAddress string `json:"contract_address"`
		Target          string `json:"target"`
		Amount          int64  `json:"amount"`
	}
)

// GetTokenBalance fetches the balance for an erc20 compatible contract
func (c *Client) GetTokenBalance(ctx context.Context, conState jsonrpc.State, args GetTokenBalance) (*big.Int, error) {
	state := State(conState)
	if state.Client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.Client.GetTokenBalance(args.ContractAddress, args.Target)
}

// TransferToken transfer an erc20 compatible token to a destination
func (c *Client) TransferTokens(ctx context.Context, conState jsonrpc.State, args TokenTransfer) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.TransferTokens(args.ContractAddress, args.Destination, args.Amount)
}

// TransferFromTokens transfer tokens from an account to another account (can be executed by anyone that is approved to spend)
func (c *Client) TransferFromTokens(ctx context.Context, conState jsonrpc.State, args TokenTransferFrom) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.TransferFromTokens(args.ContractAddress, args.From, args.Destination, args.Amount)
}

// ApproveTokenSpending approved spending of a token contract with a limit
func (c *Client) ApproveTokenSpending(ctx context.Context, conState jsonrpc.State, args ApproveTokenSpending) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.ApproveTokenSpending(args.ContractAddress, args.Target, args.Amount)
}
