package main

import (
	"net/http"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/threefoldtech/web3_proxy/server/pkg"
	"github.com/threefoldtech/web3_proxy/server/pkg/eth"
	"github.com/threefoldtech/web3_proxy/server/pkg/stellar"
	"github.com/threefoldtech/web3_proxy/server/pkg/tfchain"
)

func main() {
	// Register custom error codes
	errors := jsonrpc.NewErrors()
	errors.Register(-1001, &pkg.ErrClientNotConnected{})
	errors.Register(-2001, &stellar.ErrUnknownNetwork{})

	rpcServer := jsonrpc.NewServer(jsonrpc.WithServerErrors(errors))
	rpcServer.Register("eth", eth.NewClient())
	rpcServer.Register("stellar", stellar.NewClient())
	rpcServer.Register("tfchain", tfchain.NewClient())

	http.HandleFunc("/", rpcServer.ServeHTTP)
	http.ListenAndServe(":8080", nil)

}
