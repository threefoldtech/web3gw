package main

import (
	"net/http"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/threefoldtech/web3_proxy/server/pkg/eth"
)

func main() {
	// Register custom error codes
	errors := jsonrpc.NewErrors()
	errors.Register(-1001, &eth.ErrClientNotConnected{})

	rpcServer := jsonrpc.NewServer(jsonrpc.WithServerErrors(errors))
	rpcServer.Register("eth", eth.NewClient())

	http.HandleFunc("/", rpcServer.ServeHTTP)
	http.ListenAndServe(":8080", nil)

}
