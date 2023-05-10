module main

import os
import vweb
import time
// import threefoldtech.chat
import freeflowuniverse.crystallib.pathlib

pub struct App {
	vweb.Context
}

pub fn (mut app App) playground() vweb.Result {
	return $vweb.html()
}

pub fn (mut app App) before_request() {
	referer := app.get_header('Referer')
	hx_request := app.get_header('Hx-Request') == 'true'
	if !(hx_request) && !app.req.url.contains('static') && !app.req.url.contains('/playground'){
		app.index()
	}
}

struct NavItem {
	label string
	url string
}

pub fn (mut app App) index() vweb.Result {

	route := if app.req.url == '/' {
		'/dashboard'
	} else {
		app.req.url
	}
	sidebar := [
		NavItem{
			label: 'Dashboard'
			url: '/dashboard'
		},
		NavItem{
			label: 'Domains'
			url: '/domains'
		},
		NavItem{
			label: 'Manual'
			url: 'https://threefoldtech.github.io/web3_proxy/index.html'
		},
		NavItem{
			label: 'Playground'
			url: '/playground2'
		}
	]

	server_ip := '127.0.0.1:8000'
	last_deployment := time.now().format()
	server_status := if app.server_is_running() {
		'Connected'
	} else {
		'Disconnected'
	}
	return $vweb.html()
}

pub fn (mut app App) dashboard() vweb.Result {
	sidebar := [
		NavItem{
			label: 'Dashboard'
			url: '/'
		},
		NavItem{
			label: 'Domains'
			url: '/domains'
		},
		NavItem{
			label: 'Manual'
			url: 'https://threefoldtech.github.io/web3_proxy/index.html'

		},
		NavItem{
			label: 'Docs'
			url: '/Docs'
		},
		NavItem{
			label: 'Playground'
			url: '/playground?schemaUrl=https://raw.githubusercontent.com/threefoldtech/web3_proxy/development_openrpcdoc_fix/server/openrpc.json&uiSchema[appBar][ui:input]=false&uiSchema[appBar][ui:title]=Web3Proxy JSON-RPC API/'
		}
	]

	server_ip := '127.0.0.1:8000'
	last_deployment := time.now().format()
	server_status := if app.server_is_running() {
		'Connected'
	} else {
		'Disconnected'
	}
	return $vweb.html()
}

fn (mut app App) server_is_running() bool {
	return true
}

// pub fn (mut app App) chat() vweb.Result {
// 	mut chaty := chat.App{}
// 	return app.html(chaty.html())
// }

pub fn (mut app App) playground2() vweb.Result {
	return $vweb.html()
}

pub fn (mut app App) chat() vweb.Result {
	return $vweb.html()
}

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

pub fn main() {

	mut app := App{}
	// app.mount_static_folder_at('../build/static', '/static')
	// playground_
	// os.execute('ln -s ')
	// mount_playground()!
	app.mount_static_folder_at('static', '/static')
	vweb.run[App](app, 8081)
}

pub fn mount_playground() ! {
	os.chdir(os.dir(@FILE))!
	playground_path := pathlib.get('../playground').path
	static_path := pathlib.get('static').path
	os.execute('ln -s $playground_path/css/* $static_path/css/')
	os.execute('ln -s $static_path/js/* static/js/')
	os.execute('cp ../playground/index.html templates')
}