package eth

import (
	"context"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/threefoldtech/3bot/web3gw/server/pkg"
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
func (c *Client) GetFungibleBalance(ctx context.Context, conState jsonrpc.State, args GetFungibleBalance) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	balance, err := state.Client.GetFungibleBalance(args.ContractAddress, args.Target)
	if err != nil {
		return "", err
	}

	return balance.String(), nil
}

// OwnerOfFungible returns the owner of the given fungible token
func (c *Client) OwnerOfFungible(ctx context.Context, conState jsonrpc.State, args OwnerOfFungible) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.OwnerOfFungible(args.ContractAddress, args.TokenID)
}

// SafeTransferFungible transfers a fungible token from the given address to the given target address
func (c *Client) SafeTransferFungible(ctx context.Context, conState jsonrpc.State, args TransferFungible) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.SafeTransferFungible(ctx, args.ContractAddress, args.From, args.To, args.TokenID)
}

// TransferFungible transfers the given fungible token from the given address to the given target address
func (c *Client) TransferFungible(ctx context.Context, conState jsonrpc.State, args TransferFungible) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.TransferFungible(ctx, args.ContractAddress, args.From, args.To, args.TokenID)
}

// SetFungibleApproval approves the given address to spend the given tokenId of the given fungible token
func (c *Client) SetFungibleApproval(ctx context.Context, conState jsonrpc.State, args SetFungibleApproval) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.SetFungibleApproval(ctx, args.ContractAddress, args.From, args.To, args.Amount)
}

// SetFungibleApprovalForAll approves the given address to spend all the given fungible tokens
func (c *Client) SetFungibleApprovalForAll(ctx context.Context, conState jsonrpc.State, args SetFungibleApprovalForAll) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.SetFungibleApprovalForAll(ctx, args.ContractAddress, args.From, args.To, args.Approved)
}

// GetApprovalForFungible returns whether the given address is approved to spend the given tokenId of the given fungible token
func (c *Client) GetApprovalForFungible(ctx context.Context, conState jsonrpc.State, args ApprovalForFungible) (bool, error) {
	state := State(conState)
	if state.Client == nil {
		return false, pkg.ErrClientNotConnected{}
	}

	return state.Client.GetApprovalForFungible(args.ContractAddress, args.Owner, args.Operator)
}

// GetApprovalForAllFungible returns whether the given address is approved to spend all the given fungible tokens
func (c *Client) GetApprovalForAllFungible(ctx context.Context, conState jsonrpc.State, args ApprovalForFungible) (bool, error) {
	state := State(conState)
	if state.Client == nil {
		return false, pkg.ErrClientNotConnected{}
	}

	return state.Client.GetApprovalForAllFungible(args.ContractAddress, args.Owner, args.Operator)
}
