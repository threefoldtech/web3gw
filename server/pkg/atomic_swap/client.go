package atomicswap

import (
	"context"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/pkg/errors"
	atomicswap "github.com/threefoldtech/web3_proxy/server/clients/atomic_swap"
	"github.com/threefoldtech/web3_proxy/server/pkg/eth"
	nostrpkg "github.com/threefoldtech/web3_proxy/server/pkg/nostr"
	"github.com/threefoldtech/web3_proxy/server/pkg/stellar"
)

type (
	Client struct {
	}

	AtomicSwapState struct {
		Client *atomicswap.Client
	}
)

const (
	// AtomicSwapID is the ID for state of an atomic swap client in the connection state.
	AtomicSwapID = "atomic_swap"
)

var (
	ErrNoEthClient     = errors.New("No ethereum client loaded")
	ErrNoStellarClient = errors.New("No stellar client loaded")
	ErrNoNostrClient   = errors.New("No nostr client loaded")
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
