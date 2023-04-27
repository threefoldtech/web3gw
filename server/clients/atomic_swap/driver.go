package atomicswap

import (
	"github.com/pkg/errors"
	"github.com/threefoldtech/web3_proxy/server/clients/nostr"
)

type (
	// Driver for atomic swaps
	Driver struct {
		nostr *nostr.Client
		// TODO
	}
)

func initDriver(nostr *nostr.Client) *Driver {
	return &Driver{
		nostr: nostr,
	}
}

// Buy flow for the driver
func (d *Driver) Buy(sale nostr.Product, amount uint) error {
	// TODO
	return errors.New("TODO")
}

// OpenSale on the driver
func (d *Driver) OpenSale() error {
	// TODO
	return errors.New("TODO")
}
