package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/threefoldtech/web3_proxy/server/pkg"
	atomicswap "github.com/threefoldtech/web3_proxy/server/pkg/atomic_swap"
	"github.com/threefoldtech/web3_proxy/server/pkg/btc"
	"github.com/threefoldtech/web3_proxy/server/pkg/eth"
	"github.com/threefoldtech/web3_proxy/server/pkg/explorer"
	"github.com/threefoldtech/web3_proxy/server/pkg/ipfs"
	"github.com/threefoldtech/web3_proxy/server/pkg/nostr"
	"github.com/threefoldtech/web3_proxy/server/pkg/stellar"
	"github.com/threefoldtech/web3_proxy/server/pkg/tfchain"
	"github.com/threefoldtech/web3_proxy/server/pkg/tfgrid"
)

func main() {
	var enableIpfs, debug bool
	var port, ipfsPort uint64

	flag.Uint64Var(&port, "port", 8080, "RPC Port to listen on")
	flag.Uint64Var(&ipfsPort, "ipfs-port", 4001, "IPFS Port to listen on")

	flag.BoolVar(&enableIpfs, "ipfs", false, "Enable IPFS")
	flag.BoolVar(&debug, "debug", false, "sets debug level log output")

	flag.Parse()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Msg("debug mode enabled")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Register custom error codes
	errors := jsonrpc.NewErrors()
	errors.Register(-1001, &pkg.ErrClientNotConnected{})
	errors.Register(-2001, &stellar.ErrUnknownNetwork{})

	rpcServer := jsonrpc.NewServer(jsonrpc.WithServerErrors(errors))
	rpcServer.Register("btc", btc.NewClient())
	rpcServer.Register("eth", eth.NewClient())
	rpcServer.Register("stellar", stellar.NewClient())
	rpcServer.Register("tfchain", tfchain.NewClient())
	rpcServer.Register("tfgrid", tfgrid.NewClient())
	rpcServer.Register("nostr", nostr.NewClient())
	rpcServer.Register("explorer", explorer.NewClient())
	rpcServer.Register("atomicswap", atomicswap.NewClient())
	s := http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}

	if enableIpfs {
		log.Info().Msg("Starting IPFS server")
		go func() {
			lite, err := StartIpfsServer("0.0.0.0", ipfsPort, ctx)
			if err != nil {
				log.Error().Err(err).Msg("Failed to start IPFS server")
				panic(err)
			}
			rpcServer.Register("ipfs", ipfs.NewClient(lite))
		}()
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info().Msg("awaiting signal")
		<-sigs
		log.Info().Msg("shutting now")
		cancel()
		s.Shutdown(ctx)
	}()

	http.HandleFunc("/", rpcServer.ServeHTTP)
	log.Info().Msgf("RPC Server started on port %d", port)

	if err := s.ListenAndServe(); err != nil && err != context.Canceled {
		log.Fatal().Err(err).Msg("Failed to start HTTP server")
	}
}
