package eth

import (
	"context"
	"math/big"

	"github.com/threefoldtech/web3_proxy/server/pkg"
	"github.com/threefoldtech/web3_proxy/server/pkg/state"
)

type (
	GetFungibleBalance struct {
		ContractAddress string `json:"contract_address"`
		Target          string `json:"target"`
	}

	OwnerOfFungible struct {
		ContractAddress string `json:"contract_address"`
		TokenID         int64  `json:"token_id"`
	}

	TransferFungible struct {
		ContractAddress string `json:"contract_address"`
		From            string `json:"from"`
		To              string `json:"to"`
		TokenID         int64  `json:"token_id"`
	}

	SetFungibleApproval struct {
		ContractAddress string `json:"contract_address"`
		From            string `json:"from"`
		To              string `json:"to"`
		Amount          int64  `json:"amount"`
	}

	SetFungibleApprovalForAll struct {
		ContractAddress string `json:"contract_address"`
		From            string `json:"from"`
		To              string `json:"to"`
		Approved        bool   `json:"bool"`
	}

	ApprovalForFungible struct {
		ContractAddress string `json:"contract_address"`
		Owner           string `json:"owner"`
		Operator        string `json:"operator"`
	}
)

// GetFungibleBalance returns the balance of the given address for the given fungible token contract
func (c *Client) GetFungibleBalance(ctx context.Context, args GetFungibleBalance) (*big.Int, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetFungibleBalance(args.ContractAddress, args.Target)
}

// OwnerOfFungible returns the owner of the given fungible token
func (c *Client) OwnerOfFungible(ctx context.Context, args OwnerOfFungible) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.OwnerOfFungible(args.ContractAddress, args.TokenID)
}

// SafeTransferFungible transfers a fungible token from the given address to the given target address
func (c *Client) SafeTransferFungible(ctx context.Context, args TransferFungible) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.SafeTransferFungible(args.ContractAddress, args.From, args.To, args.TokenID)
}

// TransferFungible transfers the given fungible token from the given address to the given target address
func (c *Client) TransferFungible(ctx context.Context, args TransferFungible) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.TransferFungible(args.ContractAddress, args.From, args.To, args.TokenID)
}

// SetFungibleApproval approves the given address to spend the given tokenId of the given fungible token
func (c *Client) SetFungibleApproval(ctx context.Context, args SetFungibleApproval) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.SetFungibleApproval(args.ContractAddress, args.From, args.To, args.Amount)
}

// SetFungibleApprovalForAll approves the given address to spend all the given fungible tokens
func (c *Client) SetFungibleApprovalForAll(ctx context.Context, args SetFungibleApprovalForAll) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.SetFungibleApprovalForAll(args.ContractAddress, args.From, args.To, args.Approved)
}

// GetApprovalForFungible returns whether the given address is approved to spend the given tokenId of the given fungible token
func (c *Client) GetApprovalForFungible(ctx context.Context, args ApprovalForFungible) (bool, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return false, pkg.ErrClientNotConnected{}
	}

	return state.client.GetApprovalForFungible(args.ContractAddress, args.Owner, args.Operator)
}

// GetApprovalForAllFungible returns whether the given address is approved to spend all the given fungible tokens
func (c *Client) GetApprovalForAllFungible(ctx context.Context, args ApprovalForFungible) (bool, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return false, pkg.ErrClientNotConnected{}
	}

	return state.client.GetApprovalForAllFungible(args.ContractAddress, args.Owner, args.Operator)
}
