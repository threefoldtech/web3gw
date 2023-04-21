package eth

import (
	"context"
	"math/big"

	"github.com/threefoldtech/web3_proxy/server/pkg"
	"github.com/threefoldtech/web3_proxy/server/pkg/state"
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
func (c *Client) GetTokenBalance(ctx context.Context, args GetTokenBalance) (*big.Int, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetTokenBalance(args.ContractAddress, args.Target)
}

// TransferToken transfer an erc20 compatible token to a destination
func (c *Client) TransferTokens(ctx context.Context, args TokenTransfer) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.TransferTokens(args.ContractAddress, args.Destination, args.Amount)
}

// TransferFromTokens transfer tokens from an account to another account (can be executed by anyone that is approved to spend)
func (c *Client) TransferFromTokens(ctx context.Context, args TokenTransferFrom) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.TransferFromTokens(args.ContractAddress, args.From, args.Destination, args.Amount)
}

// ApproveTokenSpending approved spending of a token contract with a limit
func (c *Client) ApproveTokenSpending(ctx context.Context, args ApproveTokenSpending) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.ApproveTokenSpending(args.ContractAddress, args.Target, args.Amount)
}
