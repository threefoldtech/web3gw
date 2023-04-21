package tfgrid

import "github.com/threefoldtech/substrate-client"

// Client holds a tfgrid client instance.
type Client struct {
	client   TFGridClient
	TwinID   uint32
	Identity substrate.Identity
}
