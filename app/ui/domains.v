module main

import os
import vweb
import time
import freeflowuniverse.crystallib.pathlib

struct Domain {
	name string
	description string
	category string
	playground_url string
	manual_url string
	image_url string
}

pub fn (mut app App) domains() vweb.Result {
	domains := [
		Domain {
			name: 'Bitcoin'
			description: 'Web3Proxy Client for Bitcoin'
			category: 'Crypto'
			playground_url: '/playground?schemaUrl=https://raw.githubusercontent.com/threefoldtech/web3_proxy/development_openrpcdoc_fix/server/pkg/btc/openrpc.json&uiSchema[appBar][ui:input]=false&uiSchema[appBar][ui:title]=Ethereum JSON-RPC API/'
			image_url: 'https://en.bitcoin.it/w/images/en/2/29/BC_Logo_.png'
		},
		Domain {
			name: 'Ethereum'
			description: 'Web3Proxy Client for Ethereum'
			category: 'Crypto'
			playground_url: '/playground?schemaUrl=https://raw.githubusercontent.com/threefoldtech/web3_proxy/development_openrpcdoc_fix/server/pkg/eth/openrpc.json&uiSchema[appBar][ui:input]=false&uiSchema[appBar][ui:title]=Web3Proxy JSON-RPC API/'
			image_url: 'https://cryptologos.cc/logos/ethereum-eth-logo.png?v=025'
		},
		Domain {
			name: 'Stellar'
			description: 'Web3Proxy Client for Stellar'
			category: 'Crypto'
			playground_url: '/playground?schemaUrl=https://raw.githubusercontent.com/threefoldtech/web3_proxy/development_openrpcdoc_fix/server/pkg/stellar/openrpc.json&uiSchema[appBar][ui:input]=false&uiSchema[appBar][ui:title]=Web3Proxy JSON-RPC API/'
			image_url: 'https://cryptologos.cc/logos/stellar-xlm-logo.png?v=025'
		},
		Domain {
			name: 'TFChain'
			description: 'Threefold Grid JSON-RPC API'
			category: 'Crypto'
			playground_url: '/playground?schemaUrl=https://raw.githubusercontent.com/threefoldtech/web3_proxy/development_openrpcdoc_fix/server/pkg/tfchain/openrpc.json&uiSchema[appBar][ui:input]=false&uiSchema[appBar][ui:title]=Web3Proxy JSON-RPC API/'
			image_url: 'https://www.threefold.io/people/threefold-community/threefold_community.png'
		},
		Domain {
			name: 'TFGrid'
			description: 'Threefold Chain JSON-RPC API'
			category: 'Crypto'
			playground_url: '/playground?schemaUrl=https://raw.githubusercontent.com/threefoldtech/web3_proxy/development_openrpcdoc_fix/server/pkg/tfgrid/openrpc.json&uiSchema[appBar][ui:input]=false&uiSchema[appBar][ui:title]=Web3Proxy JSON-RPC API/'
			manual_url: 'https://threefoldtech.github.io/web3_proxy/client/tfgrid.html'
			image_url: 'https://www.threefold.io/people/threefold-community/threefold_community.png'
		},
		Domain {
			name: 'IPFS'
			description: 'IPFS JSON-RPC API'
			category: 'Crypto'
			playground_url: '/playground?schemaUrl=https://raw.githubusercontent.com/threefoldtech/web3_proxy/development_openrpcdoc_fix/server/pkg/ipfs/openrpc.json&uiSchema[appBar][ui:input]=false&uiSchema[appBar][ui:title]=Web3Proxy JSON-RPC API/'
			image_url: 'https://upload.wikimedia.org/wikipedia/commons/1/18/Ipfs-logo-1024-ice-text.png'
		},
		Domain {
			name: 'IPFS'
			description: 'IPFS JSON-RPC API'
			category: 'Crypto'
			playground_url: '/playground?schemaUrl=https://raw.githubusercontent.com/threefoldtech/web3_proxy/development_openrpcdoc_fix/server/pkg/ipfs/openrpc.json&uiSchema[appBar][ui:input]=false&uiSchema[appBar][ui:title]=Web3Proxy JSON-RPC API/'
			image_url: 'https://upload.wikimedia.org/wikipedia/commons/1/18/Ipfs-logo-1024-ice-text.png'
		},
	]
	return $vweb.html()
}
