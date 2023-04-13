module github.com/threefoldtech/web3_proxy/server

go 1.20

require (
	github.com/LeeSmet/go-jsonrpc v0.0.0-20230328142836-3e61d560b1c7
	github.com/ethereum/go-ethereum v1.11.5
	github.com/golang/mock v1.6.0
	github.com/google/uuid v1.3.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.29.0
	github.com/stellar/go v0.0.0-20230316102104-335848c1cd8e
	github.com/stretchr/testify v1.8.2
	github.com/threefoldtech/substrate-client v0.1.5
	github.com/threefoldtech/tfgrid-sdk-go/grid-client v0.1.0
	github.com/threefoldtech/tfgrid-sdk-go/grid-proxy v0.0.0-20230413122927-238e88d57822
	github.com/threefoldtech/tfgrid-sdk-go/rmb-sdk-go v0.0.0-20230413122927-238e88d57822
	github.com/threefoldtech/zos v0.5.6-0.20230321103809-44426c1a69c7
	golang.org/x/net v0.9.0
	golang.zx2c4.com/wireguard/wgctrl v0.0.0-20200609130330-bd2cb7843e1b
)

require (
	github.com/ChainSafe/go-schnorrkel v1.0.0 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.2 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cenkalti/backoff/v3 v3.2.2 // indirect
	github.com/centrifuge/go-substrate-rpc-client/v4 v4.0.5 // indirect
	github.com/cosmos/go-bip39 v1.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/deckarep/golang-set/v2 v2.1.0 // indirect
	github.com/decred/base58 v1.0.4 // indirect
	github.com/decred/dcrd/crypto/blake256 v1.0.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/filecoin-project/go-jsonrpc v0.2.3 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-chi/chi v4.0.3+incompatible // indirect
	github.com/go-errors/errors v0.0.0-20150906023321-a41850380601 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/gorilla/schema v1.1.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/gtank/merlin v0.1.1 // indirect
	github.com/gtank/ristretto255 v0.1.2 // indirect
	github.com/holiman/uint256 v1.2.0 // indirect
	github.com/ipfs/go-log/v2 v2.0.8 // indirect
	github.com/jbenet/go-base58 v0.0.0-20150317085156-6237cf65f3a6 // indirect
	github.com/manucorporat/sse v0.0.0-20160126180136-ee05b128a739 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/mimoo/StrobeGo v0.0.0-20220103164710-9a04d6ca976b // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pierrec/xxHash v0.1.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/cors v1.8.3 // indirect
	github.com/segmentio/go-loggly v0.5.1-0.20171222203950-eb91657e62b2 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/stellar/go-xdr v0.0.0-20211103144802-8017fc4bdfee // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	github.com/tklauser/numcpus v0.3.0 // indirect
	github.com/tyler-smith/go-bip39 v1.1.0 // indirect
	github.com/vedhavyas/go-subkey v1.0.3 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.uber.org/atomic v1.6.0 // indirect
	go.uber.org/multierr v1.5.0 // indirect
	go.uber.org/zap v1.14.1 // indirect
	golang.org/x/crypto v0.8.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/xerrors v0.0.0-20220517211312-f3a8303e98df // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/centrifuge/go-substrate-rpc-client/v4 v4.0.5 => github.com/threefoldtech/go-substrate-rpc-client/v4 v4.0.6-0.20230102154731-7c633b7d3c71
