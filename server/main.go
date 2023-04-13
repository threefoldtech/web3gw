package main

import (
	"fmt"
	"net/http"

	"github.com/LeeSmet/go-jsonrpc"
	tfgridBase "github.com/threefoldtech/web3_proxy/server/clients/tfgrid"
	"github.com/threefoldtech/web3_proxy/server/pkg"
	"github.com/threefoldtech/web3_proxy/server/pkg/eth"
	"github.com/threefoldtech/web3_proxy/server/pkg/nostr"
	"github.com/threefoldtech/web3_proxy/server/pkg/stellar"
	"github.com/threefoldtech/web3_proxy/server/pkg/tfchain"
	"github.com/threefoldtech/web3_proxy/server/pkg/tfgrid"
)

func main() {
	// Register custom error codes
	errors := jsonrpc.NewErrors()
	errors.Register(-1001, &pkg.ErrClientNotConnected{})
	errors.Register(-2001, &stellar.ErrUnknownNetwork{})
	errors.Register(-4001, &tfgridBase.ErrUnknownNetwork)

	rpcServer := jsonrpc.NewServer(jsonrpc.WithServerErrors(errors))
	rpcServer.Register("eth", eth.NewClient())
	rpcServer.Register("stellar", stellar.NewClient())
	rpcServer.Register("tfchain", tfchain.NewClient())
	rpcServer.Register("tfgrid", tfgrid.NewClient())
	rpcServer.Register("nostr", nostr.NewClient())

	http.HandleFunc("/", rpcServer.ServeHTTP)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server listening on ws://localhost:8080 ")
}
