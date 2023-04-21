package eth

import (
	"context"
	"math/big"

	"github.com/threefoldtech/web3_proxy/server/pkg"
	"github.com/threefoldtech/web3_proxy/server/pkg/state"
)

// GetMultisigOwners fetches the owner addresses for a multisig contract
func (c *Client) GetMultisigOwners(ctx context.Context, contractAddress string) ([]string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetOwners(contractAddress)
}

// GetMultisigThreshold fetches the treshold for a multisig contract
func (c *Client) GetMultisigThreshold(ctx context.Context, contractAddress string) (*big.Int, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetThreshold(contractAddress)
}

// AddMultisigOwner adds an owner to a multisig contract
func (c *Client) AddMultisigOwner(ctx context.Context, contractAddress, target string, threshold int64) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.AddOwner(contractAddress, target, threshold)
}

// RemoveMultisigOwner adds an owner to a multisig contract
func (c *Client) RemoveMultisigOwner(ctx context.Context, contractAddress, target string, threshold int64) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.RemoveOwner(contractAddress, target, threshold)
}

// ApproveHash approves a transaction hash
func (c *Client) ApproveHash(ctx context.Context, contractAddress, hash string) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.ApproveHash(contractAddress, hash)
}

// IsApproved approves a transaction hash
func (c *Client) IsApproved(ctx context.Context, contractAddress, hash string) (bool, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return false, pkg.ErrClientNotConnected{}
	}

	return state.client.IsApproved(contractAddress, hash)
}

// InitiateMultisigEthTransfer initiates a multisig eth transfer operation
func (c *Client) InitiateMultisigEthTransfer(ctx context.Context, contractAddress, destination string, amount int64) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.InitiateMultisigEthTransfer(contractAddress, destination, amount)
}

// InitiateMultisigTokenTransfer initiates a multisig eth transfer operation
func (c *Client) InitiateMultisigTokenTransfer(ctx context.Context, contractAddress, tokenAddress, destination string, amount int64) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.InitiateMultisigTokenTransfer(contractAddress, tokenAddress, destination, amount)
}