package eth

import (
	"context"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/threefoldtech/3bot/web3gw/server/pkg"
)

type (
	MultisigOwner struct {
		ContractAddress string `json:"contract_address"`
		Target          string `json:"target"`
		Threshold       int64  `json:"threshold"`
	}

	ApproveHash struct {
		ContractAddress string `json:"contract_address"`
		Hash            string `json:"hash"`
	}

	InitiateMultisigEthTransfer struct {
		ContractAddress string `json:"contract_address"`
		Destination     string `json:"destination"`
		Amount          string `json:"amount"`
	}

	InitiateMultisigTokenTransfer struct {
		ContractAddress string `json:"contract_address"`
		TokenAddress    string `json:"token_address"`
		Destination     string `json:"destination"`
		Amount          int64  `json:"amount"`
	}
)

// GetMultisigOwners fetches the owner addresses for a multisig contract
func (c *Client) GetMultisigOwners(ctx context.Context, conState jsonrpc.State, contractAddress string) ([]string, error) {
	state := State(conState)
	if state.Client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.Client.GetOwners(contractAddress)
}

// GetMultisigThreshold fetches the treshold for a multisig contract
func (c *Client) GetMultisigThreshold(ctx context.Context, conState jsonrpc.State, contractAddress string) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	threshold, err := state.Client.GetThreshold(contractAddress)
	if err != nil {
		return "", err
	}

	return threshold.String(), nil
}

// AddMultisigOwner adds an owner to a multisig contract
func (c *Client) AddMultisigOwner(ctx context.Context, conState jsonrpc.State, args MultisigOwner) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.AddOwner(args.ContractAddress, args.Target, args.Threshold)
}

// RemoveMultisigOwner adds an owner to a multisig contract
func (c *Client) RemoveMultisigOwner(ctx context.Context, conState jsonrpc.State, args MultisigOwner) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.RemoveOwner(args.ContractAddress, args.Target, args.Threshold)
}

// ApproveHash approves a transaction hash
func (c *Client) ApproveHash(ctx context.Context, conState jsonrpc.State, args ApproveHash) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.ApproveHash(args.ContractAddress, args.Hash)
}

// IsApproved approves a transaction hash
func (c *Client) IsApproved(ctx context.Context, conState jsonrpc.State, args ApproveHash) (bool, error) {
	state := State(conState)
	if state.Client == nil {
		return false, pkg.ErrClientNotConnected{}
	}

	return state.Client.IsApproved(args.ContractAddress, args.Hash)
}

// InitiateMultisigEthTransfer initiates a multisig eth transfer operation
func (c *Client) InitiateMultisigEthTransfer(ctx context.Context, conState jsonrpc.State, args InitiateMultisigEthTransfer) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.InitiateMultisigEthTransfer(args.ContractAddress, args.Destination, args.Amount)
}

// InitiateMultisigTokenTransfer initiates a multisig eth transfer operation
func (c *Client) InitiateMultisigTokenTransfer(ctx context.Context, conState jsonrpc.State, args InitiateMultisigTokenTransfer) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.InitiateMultisigTokenTransfer(args.ContractAddress, args.TokenAddress, args.Destination, args.Amount)
}
