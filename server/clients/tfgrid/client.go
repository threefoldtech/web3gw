package tfgrid

import (
	substrate "github.com/threefoldtech/tfchain/clients/tfchain-client-go"
)

// Client holds a tfgrid client instance.
type Client struct {
	client   TFGridClient
	TwinID   uint32
	Identity substrate.Identity
	// holds each Projects deployments
	Projects map[string]ProjectState
}
