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
	"github.com/drakkan/sftpgo/v2/pkg/service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/threefoldtech/3bot/web3gw/server/pkg"
	atomicswap "github.com/threefoldtech/3bot/web3gw/server/pkg/atomic_swap"
	"github.com/threefoldtech/3bot/web3gw/server/pkg/btc"
	"github.com/threefoldtech/3bot/web3gw/server/pkg/eth"
	"github.com/threefoldtech/3bot/web3gw/server/pkg/ipfs"
	"github.com/threefoldtech/3bot/web3gw/server/pkg/nostr"
	"github.com/threefoldtech/3bot/web3gw/server/pkg/stellar"
	"github.com/threefoldtech/3bot/web3gw/server/pkg/tfchain"
	"github.com/threefoldtech/3bot/web3gw/server/pkg/tfgrid"
)

const (
	defaultSFTPConfigFile        = "sftpgo.json"
	defaultSFTPLogFile           = "sftpgo.log"
	defaultSFTPLogMaxSize        = 10
	defaultSFTPLogMaxBackup      = 5
	defaultSFTPLogMaxAge         = 28
	defaultSFTPLogCompress       = false
	defaultSFTPLogLevel          = "info"
	defaultSFTPLogUTCTime        = false
	defaultSFTPLoadDataFrom      = ""
	defaultSFTPLoadDataMode      = 1
	defaultSFTPLoadDataQuotaScan = 0
	defaultSFTPLoadDataClean     = false
)

func main() {
	var enableIpfs, debug bool
	var port, ipfsPort uint64
	var sftpConfigDir string

	flag.Uint64Var(&port, "port", 8080, "RPC Port to listen on")
	flag.Uint64Var(&ipfsPort, "ipfs-port", 4001, "IPFS Port to listen on")

	flag.BoolVar(&enableIpfs, "ipfs", false, "Enable IPFS")
	flag.BoolVar(&debug, "debug", false, "sets debug level log output")
	flag.StringVar(&sftpConfigDir, "sftp-config-dir", "", "directory that includes sftpgo config file and will host sftpgo generated files")

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

	if sftpConfigDir != "" {
		sftpLogLevel := defaultSFTPLogLevel
		if debug {
			sftpLogLevel = "debug"
		}
		log.Info().Msg("Starting SFTP server")
		go func() {
			service := service.Service{
				ConfigDir:         sftpConfigDir,
				ConfigFile:        defaultSFTPConfigFile,
				LogFilePath:       defaultSFTPLogFile,
				LogMaxSize:        defaultSFTPLogMaxSize,
				LogMaxBackups:     defaultSFTPLogMaxBackup,
				LogMaxAge:         defaultSFTPLogMaxAge,
				LogCompress:       defaultSFTPLogCompress,
				LogLevel:          sftpLogLevel,
				LogUTCTime:        defaultSFTPLogUTCTime,
				LoadDataFrom:      defaultSFTPLoadDataFrom,
				LoadDataMode:      defaultSFTPLoadDataMode,
				LoadDataQuotaScan: defaultSFTPLoadDataQuotaScan,
				LoadDataClean:     defaultSFTPLoadDataClean,
				Shutdown:          make(chan bool),
			}
			if err := service.Start(false); err == nil {
				service.Wait()
				if service.Error == nil {
					os.Exit(0)
				}
			}
			os.Exit(1)
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
