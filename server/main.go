package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/threefoldtech/web3_proxy/server/pkg/eth"
)

func main() {
	fmt.Println("Hello world!")

	rpcServer := jsonrpc.NewServer()
	serverHandler := &Client{}
	rpcServer.Register("test", serverHandler)
	rpcServer.Register("eth", eth.NewClient())

	http.HandleFunc("/", rpcServer.ServeHTTP)
	http.ListenAndServe(":8080", nil)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

type Client struct {
	c int
}

func (c *Client) Add(ctx context.Context) {
	fmt.Printf("%+v/n", ctx)
	count, ok := ctx.Value("count").(int)
	if !ok {
		fmt.Println("No value found yet, initializing new")
	}
	fmt.Println("c was", c.c)
	c.c += 1
	fmt.Println("c is", c.c)
	fmt.Println("Count was", count)
	count += 1
	fmt.Println("Count is now", count)
	ctx = context.WithValue(ctx, "count", count)
}
