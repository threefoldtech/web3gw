package atomicswap

import (
	"context"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/pkg/errors"
	atomicswap "github.com/threefoldtech/3bot/web3gw/server/clients/atomic_swap"
	"github.com/threefoldtech/3bot/web3gw/server/pkg/eth"
	nostrpkg "github.com/threefoldtech/3bot/web3gw/server/pkg/nostr"
	"github.com/threefoldtech/3bot/web3gw/server/pkg/stellar"
)

type (
	Client struct {
	}

	AtomicSwapState struct {
		Client *atomicswap.Client
	}

	// SwapInfo used to start driving an atomic swap
	SwapInfo struct {
		// PaymentCurrency string of the currency used to buy TFT
		PaymentCurrency string `json:"paymentCurrency"`
		// Amount of TFT to sell (whole units)
		Amount uint64 `json:"amount"`
		// Price expressed as the smallest unit of payment currency for 1 TFT. As buyer, this is the
		// maximum price you are willing to pay for 1 TFT
		Price uint64 `json:"price"`
	}
)

const (
	// AtomicSwapID is the ID for state of an atomic swap client in the connection state.
	AtomicSwapID = "atomic_swap"
)

var (
	ErrAtomicSwapClientNotInitialized = errors.New("Atomic swap client is not initialized")
	ErrNoEthClient                    = errors.New("No ethereum client loaded")
	ErrNoStellarClient                = errors.New("No stellar client loaded")
	ErrNoNostrClient                  = errors.New("No nostr client loaded")
)

// NewClient creates a new Client ready for use
func NewClient() *Client {
	return &Client{}
}

// State from a connection. If no state is present, it is initialized
func State(conState jsonrpc.State) *AtomicSwapState {
	raw, exists := conState[AtomicSwapID]
	if !exists {
		ns := &AtomicSwapState{
			Client: nil,
		}
		conState[AtomicSwapID] = ns
		return ns
	}
	ns, ok := raw.(*AtomicSwapState)
	if !ok {
		// This means the invariant is violated, so panic here is ok
		panic("Invalid saved state for atomic swap")
	}
	return ns
}

// Close implements jsonrpc.Closer
func (s *AtomicSwapState) Close() {}

func (c *Client) Load(ctx context.Context, conState jsonrpc.State) error {
	nostrState := nostrpkg.State(conState)
	if nostrState.Client == nil {
		return ErrNoNostrClient
	}
	ethState := eth.State(conState)
	if ethState.Client == nil {
		return ErrNoEthClient
	}
	stellarState := stellar.State(conState)
	if stellarState.Client == nil {
		return ErrNoStellarClient
	}

	cl, err := atomicswap.NewClient(ctx, nostrState.Client, ethState.Client, stellarState.Client)
	if err != nil {
		return errors.Wrap(err, "could not create new atomic swap client")
	}
	state := State(conState)

	state.Client = cl

	return nil
}

func (c *Client) Sell(ctx context.Context, conState jsonrpc.State, si SwapInfo) error {
	state := State(conState)
	if state.Client == nil {
		return ErrAtomicSwapClientNotInitialized
	}

	// TODO: save driver so we can later reference it
	_, err := state.Client.PlaceSellOrder(ctx, uint(si.Amount), si.PaymentCurrency, uint(si.Price))
	if err != nil {
		return errors.Wrap(err, "could not place sell order")
	}

	return nil
}

func (c *Client) Buy(ctx context.Context, conState jsonrpc.State, si SwapInfo) error {
	state := State(conState)
	if state.Client == nil {
		return ErrAtomicSwapClientNotInitialized
	}

	// TODO: save driver so we can later reference it
	_, err := state.Client.AttemptBuy(ctx, uint(si.Amount), si.PaymentCurrency, uint(si.Price))
	if err != nil {
		return errors.Wrap(err, "could not start buying")
	}

	return nil
}
