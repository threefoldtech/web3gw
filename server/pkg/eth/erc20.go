package eth

import (
	"context"
	"math/big"

	"github.com/threefoldtech/web3_proxy/server/pkg"
	"github.com/threefoldtech/web3_proxy/server/pkg/state"
)

// GetTokenBalance fetches the balance for an erc20 compatible contract
func (c *Client) GetTokenBalance(ctx context.Context, contractAddress, target string) (*big.Int, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetTokenBalance(contractAddress, target)
}

// TransferToken transfer an erc20 compatible token to a destination
func (c *Client) TransferToken(ctx context.Context, contractAddress, destination string, amount int64) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.TransferTokens(contractAddress, destination, amount)
}

// TransferFromTokens transfer tokens from an account to another account (can be executed by anyone that is approved to spend)
func (c *Client) TransferFromTokens(ctx context.Context, contractAddress, from, to string, amount int64) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.TransferFromTokens(contractAddress, from, to, amount)
}

// ApproveTokenSpending approved spending of a token contract with a limit
func (c *Client) ApproveTokenSpending(ctx context.Context, contractAddress, target string, amount int64) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.ApproveTokenSpending(contractAddress, target, amount)
}
