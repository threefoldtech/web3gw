package eth

import (
	"context"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/ethereum/go-ethereum/common"
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
		Amount          string `json:"amount"`
	}

	TokenTransferFrom struct {
		ContractAddress string `json:"contract_address"`
		From            string `json:"from"`
		Destination     string `json:"destination"`
		Amount          string `json:"amount"`
	}

	ApproveTokenSpending struct {
		ContractAddress string `json:"contract_address"`
		Spender         string `json:"spender"`
		Amount          string `json:"amount"`
	}
)

// GetTokenBalance fetches the balance for an erc20 compatible contract
func (c *Client) GetTokenBalance(ctx context.Context, conState jsonrpc.State, contractAddress string) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	balance, err := state.Client.GetTokenBalance(contractAddress)
	if err != nil {
		return "", err
	}
	return balance, nil
}

// TransferToken transfer an erc20 compatible token to a destination
func (c *Client) TransferTokens(ctx context.Context, conState jsonrpc.State, args TokenTransfer) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.TransferTokens(ctx, common.HexToAddress(args.ContractAddress), args.Destination, args.Amount)
}

// TransferFromTokens transfer tokens from an account to another account (can be executed by anyone that is approved to spend)
func (c *Client) TransferFromTokens(ctx context.Context, conState jsonrpc.State, args TokenTransferFrom) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.TransferFromTokens(ctx, args.ContractAddress, args.From, args.Destination, args.Amount)
}

// ApproveTokenSpending approves spending from a token contract with a limit
func (c *Client) ApproveTokenSpending(ctx context.Context, conState jsonrpc.State, args ApproveTokenSpending) (string, error) {
	state := State(conState)
	if state.Client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.Client.ApproveTokenSpending(ctx, args.ContractAddress, args.Spender, args.Amount)
}
