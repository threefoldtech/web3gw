package stellargoclient

import (
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
func NewClient(secret, stellarNetwork string) (*Client, error) {
	cl := &Client{
		stellarNetwork: stellarNetwork,
		horizon:        GetHorizonClient(stellarNetwork),
		kp:             nil,
	}

	if secret != "" {
		k, err := GetKeypairFromSeed(secret)
		if err != nil {
			return nil, err
		}
		cl.kp = k
	} else {
		k, err := cl.GenerateAccount()
		if err != nil {
			return nil, err
		}
		cl.kp = k
	}

	return cl, nil
}

// Address of the loaded keypair
func (c *Client) Address() string {
	return c.kp.Address()
}

// KeyPair loaded in the client
func (c *Client) KeyPair() keypair.Full {
	return *c.kp
}
