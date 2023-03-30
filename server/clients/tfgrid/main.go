package tfgrid

import (
	"errors"

	"github.com/threefoldtech/grid3-go/deployer"
)

type (
	Client struct {
		cl *deployer.TFPluginClient
	}
)

const (
	// keyType for the TF grid
	keyType = "sr25519"

	// NetworkMain is the TF grid mainnet
	NetworkMain = "main"
	// NetworkTest is the TF grid testnet
	NetworkTest = "test"
	// NetworkQa is the TF grid qanet
	NetworkQA = "qa"
	// NetworkDev is the TF grid devnet
	NetworkDev = "dev"

	// DeployerTimeoutSeconds is the amount of seconds before deployment operations time out
	DeployerTimeoutSeconds = 600 // 10 minutes
)

var (
	// KnownNets are the known tfgrid networks
	KnownNets = map[string]struct{}{
		NetworkMain: {},
		NetworkTest: {},
		NetworkQA:   {},
		NetworkDev:  {},
	}
	// ErrUnknownNetwork indicates the network type is not known
	ErrUnknownNetwork = errors.New("unknown network type, valid networks are main,test,qa,dev")
)

// NewClient creates a new tf grid deployment client
func NewClient(mnemonic string, network string) (*Client, error) {
	if _, ok := KnownNets[network]; !ok {
		return nil, ErrUnknownNetwork
	}
	// TODO: network check
	cl, err := deployer.NewTFPluginClient(mnemonic, keyType, network, deployer.SubstrateURLs[network], deployer.RelayURLS[network], deployer.RMBProxyURLs[network], DeployerTimeoutSeconds, true, false)
	if err != nil {
		return nil, err
	}

	return &Client{cl: &cl}, nil
}
