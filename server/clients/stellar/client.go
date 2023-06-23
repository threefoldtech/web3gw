package stellargoclient

import (
	"github.com/rs/zerolog/log"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
)

type Client struct {
	stellarNetwork string
	horizon        *horizonclient.Client
	kp             *keypair.Full
}

// NewClient creates a new client
// stellarNetwork can be "testnet" or "public"
// if stellarNetwork is not "testnet" or "public" it will default to "testnet"
func NewClient(stellarNetwork string) *Client {
	log.Debug().Msgf("Creating stellar client for the %s network", stellarNetwork)

	return &Client{
		stellarNetwork: stellarNetwork,
		horizon:        GetHorizonClient(stellarNetwork),
		kp:             nil,
	}
}

// Address of the loaded keypair
func (c *Client) Address() string {
	return c.kp.Address()
}

// KeyPair loaded in the client
func (c *Client) KeyPair() keypair.Full {
	return *c.kp
}
