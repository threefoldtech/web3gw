package stellargoclient

import (
	"github.com/pkg/errors"
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
func NewClient(secret, stellarNetwork string) (*Client, error) {

	log.Debug().Msgf("Creating stellar client for the %s network", stellarNetwork)

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

		// check if account has trustline, if not add it
		accountRequest := horizonclient.AccountRequest{AccountID: k.Address()}
		hAccount, err := cl.horizon.AccountDetail(accountRequest)
		if err != nil {
			return nil, errors.Wrap(err, "account does not exist")
		}

		if !hasTrustline(hAccount, cl.GetTftBaseAsset()) {
			log.Debug().Msgf("Adding trustline for account %s", k.Address())
			cl.SetTrustLine(k.Address())
		}
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
