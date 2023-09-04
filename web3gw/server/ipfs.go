package main

import (
	"context"
	"fmt"

	ipfslite "github.com/hsanjuan/ipfs-lite"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/multiformats/go-multiaddr"
)

func StartIpfsServer(host string, port uint64, ctx context.Context) (*ipfslite.Peer, error) {
	ds := ipfslite.NewInMemoryDatastore()
	priv, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	if err != nil {
		return nil, err
	}

	listen, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%d", host, port))
	if err != nil {
		return nil, err
	}

	h, dht, err := ipfslite.SetupLibp2p(
		ctx,
		priv,
		nil,
		[]multiaddr.Multiaddr{listen},
		ds,
		ipfslite.Libp2pOptionsExtra...,
	)

	if err != nil {
		return nil, err
	}

	lite, err := ipfslite.New(ctx, ds, nil, h, dht, nil)
	if err != nil {
		return nil, err
	}

	lite.Bootstrap(ipfslite.DefaultBootstrapPeers())

	return lite, nil
}
