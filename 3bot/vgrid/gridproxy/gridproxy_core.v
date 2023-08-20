module gridproxy

// client library for threefold gridproxy API.
import json
import model { Contract, ContractFilter, ContractIterator, Farm, FarmFilter, FarmIterator, GridStat, Node, NodeFilter, NodeIterator, Node_, StatFilter, Twin, TwinFilter, TwinIterator }

/*
all errors returned by the gridproxy API or the client are wrapped in a standard `Error` object with two fields.
{
	msg string
	code int // could be API call error code or client error code
}

`code` is an error code that can be used to identify the error.
in API call errors, `code` represents the HTTP status code. (100..599)

Client errors codes are represented by numbers in the range of 1..99
currently, the following client error codes are used:
id not found error code: 4
json parsing error code: 10
http client error code: 11
invalid response from server (e.g. empty response) error code: 24
*/
const (
	// clinet error codes
	err_not_found    = 4
	err_json_parse   = 10
	err_http_client  = 11
	err_invalid_resp = 24
)

// get_node_by_id fetchs specific node information by node id.
//
// * `node_id` (u64): node id.
//
// returns: `Node` or `Error`.
pub fn (mut c GridProxyClient) get_node_by_id(node_id u64) !Node {
	// needed to allow to use threads
	mut http_client := c.http_client.clone()!

	res := http_client.send(prefix: 'nodes/', id: '${node_id}') or {
		return error_with_code('http client error: ${err.msg()}', gridproxy.err_http_client)
	}

	if !res.is_ok() {
		return error_with_code(res.data, res.code)
	}

	if res.data == '' {
		return error_with_code('empty response', gridproxy.err_invalid_resp)
	}

	node := json.decode(Node, res.data) or {
		return error_with_code('error to get jsonstr for node data, json decode: node id: ${node_id}, data: ${res.data}',
			gridproxy.err_json_parse)
	}
	return node
}

// get_gateway_by_id fetchs specific gateway information by node id.
//
// * `node_id` (u64): node id.
//
// returns: `Node` or `Error`.
pub fn (mut c GridProxyClient) get_gateway_by_id(node_id u64) !Node {
	// needed to allow to use threads
	mut http_client := c.http_client.clone()!

	res := http_client.send(prefix: 'gateways/', id: '${node_id}') or {
		return error_with_code('http client error: ${err.msg()}', gridproxy.err_http_client)
	}

	if !res.is_ok() {
		return error_with_code(res.data, res.code)
	}

	if res.data == '' {
		return error_with_code('empty response', gridproxy.err_invalid_resp)
	}

	node := json.decode(Node, res.data) or {
		return error_with_code('error to get jsonstr for gateway data, json decode: gateway id: ${node_id}, data: ${res.data}',
			gridproxy.err_json_parse)
	}
	return node
}

// get_nodes fetchs nodes information and public configurations with pagination.
//
// * `page` (u64): Page number. [optional].
// * `size` (u64): Max result per page. [optional].
// * `ret_count` (bool): Set nodes' count on headers based on filter. [optional].
// * `free_mru` (u64): Min free reservable mru in bytes. [optional].
// * `free_hru` (u64): Min free reservable hru in bytes. [optional].
// * `free_sru` (u64): Min free reservable sru in bytes. [optional].
// * `free_ips` (u64): Min number of free ips in the farm of the node. [optional].
// * `status` (string): Node status filter, set to 'up' to get online nodes only. [optional].
// * `city` (string): Node city filter. [optional].
// * `country` (string): Node country filter. [optional].
// * `farm_name` (string): Get nodes for specific farm. [optional].
// * `ipv4` (string): Set to true to filter nodes with ipv4. [optional].
// * `ipv6` (string): Set to true to filter nodes with ipv6. [optional].
// * `domain` (string): Set to true to filter nodes with domain. [optional].
// * `dedicated` (bool): Set to true to get the dedicated nodes only. [optional].
// * `rentable` (bool): Set to true to filter the available nodes for renting. [optional].
// * `rented_by` (u64): rented by twin id. [optional].
// * `available_for` (u64): available for twin id. [optional].
// * `farm_ids` ([]u64): List of farm ids. [optional].
//
// returns: `[]Node` or `Error`.
pub fn (mut c GridProxyClient) get_nodes(params NodeFilter) ![]Node {
	// needed to allow to use threads
	mut http_client := c.http_client.clone()!
	params_map := params.to_map()
	res := http_client.send(prefix: 'nodes/', params: params_map) or {
		return error_with_code('http client error: ${err.msg()}', gridproxy.err_http_client)
	}

	if !res.is_ok() {
		return error_with_code(res.data, res.code)
	}

	if res.data == '' {
		return error_with_code('empty response', gridproxy.err_invalid_resp)
	}

	nodes_ := json.decode([]Node_, res.data) or {
		return error_with_code('error to get jsonstr for node list data, json decode: node filter: ${params_map}, data: ${res.data}',
			gridproxy.err_json_parse)
	}
	nodes := nodes_.map(it.with_nested_capacity())
	return nodes
}

// get_gateways fetchs gateways information and public configurations and domains with pagination.
//
// * `page` (u64): Page number. [optional].
// * `size` (u64): Max result per page. [optional].
// * `ret_count` (bool): Set nodes' count on headers based on filter. [optional].
// * `free_mru` (u64): Min free reservable mru in bytes. [optional].
// * `free_hru` (u64): Min free reservable hru in bytes. [optional].
// * `free_sru` (u64): Min free reservable sru in bytes. [optional].
// * `free_ips` (u64): Min number of free ips in the farm of the node. [optional].
// * `status` (string): Node status filter, set to 'up' to get online nodes only.. [optional].
// * `city` (string): Node city filter. [optional].
// * `country` (string): Node country filter. [optional].
// * `farm_name` (string): Get nodes for specific farm. [optional].
// * `ipv4` (string): Set to true to filter nodes with ipv4. [optional].
// * `ipv6` (string): Set to true to filter nodes with ipv6. [optional].
// * `domain` (string): Set to true to filter nodes with domain. [optional].
// * `dedicated` (bool): Set to true to get the dedicated nodes only. [optional].
// * `rentable` (bool): Set to true to filter the available nodes for renting. [optional].
// * `rented_by` (u64): rented by twin id. [optional].
// * `available_for` (u64): available for twin id. [optional].
// * `farm_ids` ([]u64): List of farm ids. [optional].
//
// returns: `[]Node` or `Error`.
pub fn (mut c GridProxyClient) get_gateways(params NodeFilter) ![]Node {
	// needed to allow to use threads
	mut http_client := c.http_client.clone()!
	params_map := params.to_map()
	res := http_client.send(prefix: 'gateways/', params: params_map) or {
		return error_with_code('http client error: ${err.msg()}', gridproxy.err_http_client)
	}

	if !res.is_ok() {
		return error_with_code(res.data, res.code)
	}

	if res.data == '' {
		return error_with_code('empty response', gridproxy.err_invalid_resp)
	}

	nodes_ := json.decode([]Node_, res.data) or {
		return error_with_code('error to get jsonstr for gateways list data, json decode: gateway filter: ${params_map}, data: ${res.data}',
			gridproxy.err_json_parse)
	}
	nodes := nodes_.map(it.with_nested_capacity())
	return nodes
}

// get_stats fetchs stats about the grid.
//
// * `status` (string): Node status filter, set to 'up' to get online nodes only.. [optional].
//
// returns: `GridStat` or `Error`.
pub fn (mut c GridProxyClient) get_stats(filter StatFilter) !GridStat {
	// needed to allow to use threads
	mut http_client := c.http_client.clone()!
	mut params_map := map[string]string{}
	params_map['status'] = match filter.status {
		.all { '' }
		.online { 'up' }
	}

	res := http_client.send(prefix: 'stats/', params: params_map) or {
		return error_with_code('http client error: ${err.msg()}', gridproxy.err_http_client)
	}

	if !res.is_ok() {
		return error_with_code(res.data, res.code)
	}

	if res.data == '' {
		return error_with_code('empty response', gridproxy.err_invalid_resp)
	}

	stats := json.decode(GridStat, res.data) or {
		return error_with_code('error to get jsonstr for grid stats data, json decode: stats filter: ${params_map}, data: ${res.data}',
			gridproxy.err_json_parse)
	}
	return stats
}

// get_twins fetchs twins information with pagaination.
//
// * `page` (u64): Page number. [optional].
// * `size` (u64): Max result per page. [optional].
// * `ret_count` (bool): Set farms' count on headers based on filter. [optional].
// * `twin_id` (u64): twin id. [optional].
// * `account_id` (string): account address. [optional].
//
// returns: `[]Twin` or `Error`.
pub fn (mut c GridProxyClient) get_twins(params TwinFilter) ![]Twin {
	// needed to allow to use threads
	mut http_client := c.http_client.clone()!
	params_map := params.to_map()
	res := http_client.send(prefix: 'twins/', params: params_map) or {
		return error_with_code('http client error: ${err.msg()}', gridproxy.err_http_client)
	}

	if !res.is_ok() {
		return error_with_code(res.data, res.code)
	}

	if res.data == '' {
		return error_with_code('empty response', gridproxy.err_invalid_resp)
	}

	twins := json.decode([]Twin, res.data) or {
		return error_with_code('error to get jsonstr for twin list data, json decode: twin filter: ${params_map}, data: ${res.data}',
			gridproxy.err_json_parse)
	}
	return twins
}

// get_contracts fetchs contracts information with pagination.
//
// * `page` (u64): Page number. [optional].
// * `size` (u64): Max result per page. [optional].
// * `ret_count` (bool): Set farms' count on headers based on filter. [optional].
// * `contract_id` (u64): contract id. [optional].
// * `twin_id` (u64): twin id. [optional].
// * `node_id` (u64): node id which contract is deployed on in case of ('rent' or 'node' contracts). [optional].
// * `name` (string): contract name in case of 'name' contracts. [optional].
// * `type` (string): contract type 'node', 'name', or 'rent'. [optional].
// * `state` (string): contract state 'Created', or 'Deleted'. [optional].
// * `deployment_data` (string): contract deployment data in case of 'node' contracts. [optional].
// * `deployment_hash` (string): contract deployment hash in case of 'node' contracts. [optional].
// * `number_of_public_ips` (u64): Min number of public ips in the 'node' contract. [optional].
//
// * returns: `[]Contract` or `Error`.
pub fn (mut c GridProxyClient) get_contracts(params ContractFilter) ![]Contract {
	// needed to allow to use threads
	mut http_client := c.http_client.clone()!
	params_map := params.to_map()
	res := http_client.send(prefix: 'contracts/', params: params_map) or {
		return error_with_code('http client error: ${err.msg()}', gridproxy.err_http_client)
	}

	if !res.is_ok() {
		return error_with_code(res.data, res.code)
	}

	if res.data == '' {
		return error_with_code('empty response', gridproxy.err_invalid_resp)
	}

	contracts := json.decode([]Contract, res.data) or {
		return error_with_code('error to get jsonstr for contract list data, json decode: contract filter: ${params_map}, data: ${res.data}',
			gridproxy.err_json_parse)
	}
	return contracts
}

// get_farms fetchs farms information and public ips.
//
// * `page` (u64): Page number. [optional].
// * `size` (u64): Max result per page. [optional].
// * `ret_count` (bool): Set farms' count on headers based on filter. [optional].
// * `free_ips` (u64): Min number of free ips in the farm. [optional].
// * `total_ips` (u64): Min number of total ips in the farm. [optional].
// * `pricing_policy_id` (u64): Pricing policy id. [optional].
// * `version` (u64): farm version. [optional].
// * `farm_id` (u64): farm id. [optional].
// * `twin_id` (u64): twin id associated with the farm. [optional].
// * `name` (string): farm name. [optional].
// * `name_contains` (string): farm name contains. [optional].
// * `certification_type` (string): certificate type DIY or Certified. [optional].
// * `dedicated` (bool): farm is dedicated. [optional].
// * `stellar_address` (string): farm stellar_address. [optional].
//
// returns: `[]Farm` or `Error`.
pub fn (mut c GridProxyClient) get_farms(params FarmFilter) ![]Farm {
	// needed to allow to use threads
	mut http_client := c.http_client.clone()!
	params_map := params.to_map()
	res := http_client.send(prefix: 'farms/', params: params_map) or {
		return error_with_code('http client error: ${err.msg()}', gridproxy.err_http_client)
	}

	if !res.is_ok() {
		return error_with_code(res.data, res.code)
	}

	if res.data == '' {
		return error_with_code('empty response', gridproxy.err_invalid_resp)
	}

	farms := json.decode([]Farm, res.data) or {
		return error_with_code('error to get jsonstr for farm list data, json decode: farm filter: ${params_map}, data: ${res.data}',
			gridproxy.err_json_parse)
	}
	return farms
}

// is_pingable checks if API server is reachable and responding.
//
// returns: bool, `true` if API server is reachable and responding, `false` otherwise
pub fn (mut c GridProxyClient) is_pingable() !bool {
	mut http_client := c.http_client.clone()!
	res := http_client.send(prefix: 'ping/') or { return false }
	if !res.is_ok() {
		return false
	}
	health_map := json.decode(map[string]string, res.data) or { return false }

	if health_map['ping'] != 'pong' {
		return false
	}

	return true
}

// Iterators have the next() method, which returns the next page of the objects.
// to be used in a loop to get all available results, or to lazely traverse pages till a specific condition is met.

// get_nodes_iterator creates an iterator through node pages with custom filter
pub fn (mut c GridProxyClient) get_nodes_iterator(filter NodeFilter) NodeIterator {
	return NodeIterator{filter, c.get_nodes}
}

// get_gateways_iterator creates an iterator through gateway pages with custom filter
pub fn (mut c GridProxyClient) get_gateways_iterator(filter NodeFilter) NodeIterator {
	return NodeIterator{filter, c.get_gateways}
}

// get_farms_iterator creates an iterator through farms pages with custom filter
pub fn (mut c GridProxyClient) get_farms_iterator(filter FarmFilter) FarmIterator {
	return FarmIterator{filter, c.get_farms}
}

// get_twins_iterator creates an iterator through twin pages with custom filter
pub fn (mut c GridProxyClient) get_twins_iterator(filter TwinFilter) TwinIterator {
	return TwinIterator{filter, c.get_twins}
}

// get_contracts_iterator creates an iterator through contracts pages with custom filter
pub fn (mut c GridProxyClient) get_contracts_iterator(filter ContractFilter) ContractIterator {
	return ContractIterator{filter, c.get_contracts}
}
