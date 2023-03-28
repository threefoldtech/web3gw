package main

import (
	"net/http"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/threefoldtech/web3_proxy/server/pkg/eth"
)

func main() {

	rpcServer := jsonrpc.NewServer()
	rpcServer.Register("eth", eth.NewClient())

	http.HandleFunc("/", rpcServer.ServeHTTP)
	http.ListenAndServe(":8080", nil)

}
