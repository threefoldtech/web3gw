package eth

import (
	"context"
	"math/big"

	"github.com/threefoldtech/web3_proxy/server/pkg"
	"github.com/threefoldtech/web3_proxy/server/pkg/state"
)

// GetFungibleBalance returns the balance of the given address for the given fungible token contract
func (c *Client) GetFungibleBalance(ctx context.Context, contractAddress, target string) (*big.Int, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetFungibleBalance(contractAddress, target)
}

// OwnerOfFungible returns the owner of the given fungible token
func (c *Client) OwnerOfFungible(ctx context.Context, contractAddress string, token int64) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.OwnerOfFungible(contractAddress, token)
}

// SafeTransferFungible transfers a fungible token from the given address to the given target address
func (c *Client) SafeTransferFungible(ctx context.Context, contractAddress, from, to string, tokenId int64) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.SafeTransferFungible(contractAddress, from, to, tokenId)
}

// TransferFungible transfers the given fungible token from the given address to the given target address
func (c *Client) TransferFungible(ctx context.Context, contractAddress, from, to string, tokenId int64) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.TransferFungible(contractAddress, from, to, tokenId)
}

// SetFungibleApproval approves the given address to spend the given tokenId of the given fungible token
func (c *Client) SetFungibleApproval(ctx context.Context, contractAddress, from, to string, tokenId int64) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.SetFungibleApproval(contractAddress, from, to, tokenId)
}

// SetFungibleApprovalForAll approves the given address to spend all the given fungible tokens
func (c *Client) SetFungibleApprovalForAll(ctx context.Context, contractAddress, from, to string, approved bool) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.SetFungibleApprovalForAll(contractAddress, from, to, approved)
}

// GetApprovalForFungible returns whether the given address is approved to spend the given tokenId of the given fungible token
func (c *Client) GetApprovalForFungible(ctx context.Context, contractAddress, owner, operator string) (bool, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return false, pkg.ErrClientNotConnected{}
	}

	return state.client.GetApprovalForFungible(contractAddress, owner, operator)
}

// GetApprovalForAllFungible returns whether the given address is approved to spend all the given fungible tokens
func (c *Client) GetApprovalForAllFungible(ctx context.Context, contractAddress, owner, operator string) (bool, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return false, pkg.ErrClientNotConnected{}
	}

	return state.client.GetApprovalForAllFungible(contractAddress, owner, operator)
}
