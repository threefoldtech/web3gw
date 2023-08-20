# module threefoldtech.vgrid.gridproxy




## Contents
- [get](#get)
- [TFGridNet](#TFGridNet)
- [GridProxyClient](#GridProxyClient)
  - [get_contracts](#get_contracts)
  - [get_contracts_by_node_id](#get_contracts_by_node_id)
  - [get_contracts_by_twin_id](#get_contracts_by_twin_id)
  - [get_contracts_iterator](#get_contracts_iterator)
  - [get_farm_by_id](#get_farm_by_id)
  - [get_farm_by_name](#get_farm_by_name)
  - [get_farms](#get_farms)
  - [get_farms_by_twin_id](#get_farms_by_twin_id)
  - [get_farms_iterator](#get_farms_iterator)
  - [get_gateway_by_id](#get_gateway_by_id)
  - [get_gateways](#get_gateways)
  - [get_gateways_iterator](#get_gateways_iterator)
  - [get_node_by_id](#get_node_by_id)
  - [get_node_stats_by_id](#get_node_stats_by_id)
  - [get_nodes](#get_nodes)
  - [get_nodes_has_resources](#get_nodes_has_resources)
  - [get_nodes_iterator](#get_nodes_iterator)
  - [get_stats](#get_stats)
  - [get_twin_by_account](#get_twin_by_account)
  - [get_twin_by_id](#get_twin_by_id)
  - [get_twins](#get_twins)
  - [get_twins_iterator](#get_twins_iterator)
  - [is_pingable](#is_pingable)

## get
```v
fn get(net TFGridNet, use_redis_cache bool) !&GridProxyClient
```

get returns a gridproxy client for the given net.  

* `net` (enum): the net to get the gridproxy client for (one of .main, .test, .dev, .qa).  
* `use_redis_cache` (bool): if true, the gridproxy client will use a redis cache and redis should be running on the host. otherwise, the gridproxy client will not use cache.  

returns: `&GridProxyClient`.  

[[Return to contents]](#Contents)

## TFGridNet
```v
enum TFGridNet {
	main
	test
	dev
	qa
}
```


[[Return to contents]](#Contents)

## GridProxyClient
```v
struct GridProxyClient {
pub mut:
	http_client httpconnection.HTTPConnection
}
```


[[Return to contents]](#Contents)

## get_contracts
```v
fn (mut c GridProxyClient) get_contracts(params ContractFilter) ![]Contract
```

get_contracts fetchs contracts information with pagination.  

* `page` (u64): Page number. [optional].  
* `size` (u64): Max result per page. [optional].  
* `ret_count` (bool): Set farms' count on headers based on filter. [optional].  
* `contract_id` (u64): contract id. [optional].  
* `twin_id` (u64): twin id. [optional].  
* `node_id` (u64): node id which contract is deployed on in case of ('rent' or 'node' contracts). [optional].  
* `name` (string): contract name in case of 'name' contracts. [optional].  
* `type` (string): contract type 'node', 'name', or 'rent'. [optional].  
* `state` (string): contract state 'Created', or 'Deleted'. [optional].  
* `deployment_data` (string): contract deployment data in case of 'node' contracts. [optional].  
* `deployment_hash` (string): contract deployment hash in case of 'node' contracts. [optional].  
* `number_of_public_ips` (u64): Min number of public ips in the 'node' contract. [optional].  

* returns: `[]Contract` or `Error`.  

[[Return to contents]](#Contents)

## get_contracts_by_node_id
```v
fn (mut c GridProxyClient) get_contracts_by_node_id(node_id u64) ContractIterator
```

get_contracts_by_node_id returns iterator over all contracts deployed on specific node.  

* `node_id`: node id.  

returns: `ContractIterator`.  

[[Return to contents]](#Contents)

## get_contracts_by_twin_id
```v
fn (mut c GridProxyClient) get_contracts_by_twin_id(twin_id u64) ContractIterator
```

get_contracts_by_twin_id returns iterator over all contracts owned by specific twin.  

* `twin_id`: twin id.  

returns: `ContractIterator`.  

[[Return to contents]](#Contents)

## get_contracts_iterator
```v
fn (mut c GridProxyClient) get_contracts_iterator(filter ContractFilter) ContractIterator
```

get_contracts_iterator creates an iterator through contracts pages with custom filter

[[Return to contents]](#Contents)

## get_farm_by_id
```v
fn (mut c GridProxyClient) get_farm_by_id(farm_id u64) ?Farm
```

fetch specific farm information by id.  

* `farm_id`: farm id.  

returns: `Farm` or `Error`.  

[[Return to contents]](#Contents)

## get_farm_by_name
```v
fn (mut c GridProxyClient) get_farm_by_name(farm_name string) ?Farm
```

fetch specific farm information by farm name.  

* `farm_name`: farm name.  

returns: `Farm` or `Error`.  

[[Return to contents]](#Contents)

## get_farms
```v
fn (mut c GridProxyClient) get_farms(params FarmFilter) ![]Farm
```

get_farms fetchs farms information and public ips.  

* `page` (u64): Page number. [optional].  
* `size` (u64): Max result per page. [optional].  
* `ret_count` (bool): Set farms' count on headers based on filter. [optional].  
* `free_ips` (u64): Min number of free ips in the farm. [optional].  
* `total_ips` (u64): Min number of total ips in the farm. [optional].  
* `pricing_policy_id` (u64): Pricing policy id. [optional].  
* `version` (u64): farm version. [optional].  
* `farm_id` (u64): farm id. [optional].  
* `twin_id` (u64): twin id associated with the farm. [optional].  
* `name` (string): farm name. [optional].  
* `name_contains` (string): farm name contains. [optional].  
* `certification_type` (string): certificate type DIY or Certified. [optional].  
* `dedicated` (bool): farm is dedicated. [optional].  
* `stellar_address` (string): farm stellar_address. [optional].  

returns: `[]Farm` or `Error`.  

[[Return to contents]](#Contents)

## get_farms_by_twin_id
```v
fn (mut c GridProxyClient) get_farms_by_twin_id(twin_id u64) FarmIterator
```

get_farms_by_twin_id returns iterator over all farms information associated with specific twin.  

* `twin_id`: twin id.  

returns: `FarmIterator`.  

[[Return to contents]](#Contents)

## get_farms_iterator
```v
fn (mut c GridProxyClient) get_farms_iterator(filter FarmFilter) FarmIterator
```

get_farms_iterator creates an iterator through farms pages with custom filter

[[Return to contents]](#Contents)

## get_gateway_by_id
```v
fn (mut c GridProxyClient) get_gateway_by_id(node_id u64) !Node
```

get_gateway_by_id fetchs specific gateway information by node id.  

* `node_id` (u64): node id.  

returns: `Node` or `Error`.  

[[Return to contents]](#Contents)

## get_gateways
```v
fn (mut c GridProxyClient) get_gateways(params NodeFilter) ![]Node
```

get_gateways fetchs gateways information and public configurations and domains with pagination.  

* `page` (u64): Page number. [optional].  
* `size` (u64): Max result per page. [optional].  
* `ret_count` (bool): Set nodes' count on headers based on filter. [optional].  
* `free_mru` (u64): Min free reservable mru in bytes. [optional].  
* `free_hru` (u64): Min free reservable hru in bytes. [optional].  
* `free_sru` (u64): Min free reservable sru in bytes. [optional].  
* `free_ips` (u64): Min number of free ips in the farm of the node. [optional].  
* `status` (string): Node status filter, set to 'up' to get online nodes only.. [optional].  
* `city` (string): Node city filter. [optional].  
* `country` (string): Node country filter. [optional].  
* `farm_name` (string): Get nodes for specific farm. [optional].  
* `ipv4` (string): Set to true to filter nodes with ipv4. [optional].  
* `ipv6` (string): Set to true to filter nodes with ipv6. [optional].  
* `domain` (string): Set to true to filter nodes with domain. [optional].  
* `dedicated` (bool): Set to true to get the dedicated nodes only. [optional].  
* `rentable` (bool): Set to true to filter the available nodes for renting. [optional].  
* `rented_by` (u64): rented by twin id. [optional].  
* `available_for` (u64): available for twin id. [optional].  
* `farm_ids` ([]u64): List of farm ids. [optional].  

returns: `[]Node` or `Error`.  

[[Return to contents]](#Contents)

## get_gateways_iterator
```v
fn (mut c GridProxyClient) get_gateways_iterator(filter NodeFilter) NodeIterator
```

get_gateways_iterator creates an iterator through gateway pages with custom filter

[[Return to contents]](#Contents)

## get_node_by_id
```v
fn (mut c GridProxyClient) get_node_by_id(node_id u64) !Node
```

get_node_by_id fetchs specific node information by node id.  

* `node_id` (u64): node id.  

returns: `Node` or `Error`.  

[[Return to contents]](#Contents)

## get_node_stats_by_id
```v
fn (mut c GridProxyClient) get_node_stats_by_id(node_id u64) !NodeStats
```

get_node_stats_by_id fetchs specific node statistics by node id.  

* `node_id` (u64): node id.  
statistics
returns: `Node_stats` or `Error`.  

[[Return to contents]](#Contents)

## get_nodes
```v
fn (mut c GridProxyClient) get_nodes(params NodeFilter) ![]Node
```

get_nodes fetchs nodes information and public configurations with pagination.  

* `page` (u64): Page number. [optional].  
* `size` (u64): Max result per page. [optional].  
* `ret_count` (bool): Set nodes' count on headers based on filter. [optional].  
* `free_mru` (u64): Min free reservable mru in bytes. [optional].  
* `free_hru` (u64): Min free reservable hru in bytes. [optional].  
* `free_sru` (u64): Min free reservable sru in bytes. [optional].  
* `free_ips` (u64): Min number of free ips in the farm of the node. [optional].  
* `status` (string): Node status filter, set to 'up' to get online nodes only. [optional].  
* `city` (string): Node city filter. [optional].  
* `country` (string): Node country filter. [optional].  
* `farm_name` (string): Get nodes for specific farm. [optional].  
* `ipv4` (string): Set to true to filter nodes with ipv4. [optional].  
* `ipv6` (string): Set to true to filter nodes with ipv6. [optional].  
* `domain` (string): Set to true to filter nodes with domain. [optional].  
* `dedicated` (bool): Set to true to get the dedicated nodes only. [optional].  
* `rentable` (bool): Set to true to filter the available nodes for renting. [optional].  
* `rented_by` (u64): rented by twin id. [optional].  
* `available_for` (u64): available for twin id. [optional].  
* `farm_ids` ([]u64): List of farm ids. [optional].  

returns: `[]Node` or `Error`.  

[[Return to contents]](#Contents)

## get_nodes_has_resources
```v
fn (mut c GridProxyClient) get_nodes_has_resources(filter ResourceFilter) NodeIterator
```

get_nodes_has_resources returns iterator over all nodes with specific minimum free reservable resources.  

* `free_ips` (u64): minimum free ips. [optional].  
* `free_mru_gb` (u64): minimum free mru in GB. [optional].  
* `free_sru_gb` (u64): minimum free sru in GB. [optional].  
* `free_hru_gb` (u64): minimum free hru in GB. [optional].  

returns: `NodeIterator`.  

[[Return to contents]](#Contents)

## get_nodes_iterator
```v
fn (mut c GridProxyClient) get_nodes_iterator(filter NodeFilter) NodeIterator
```

get_nodes_iterator creates an iterator through node pages with custom filter

[[Return to contents]](#Contents)

## get_stats
```v
fn (mut c GridProxyClient) get_stats(filter StatFilter) !GridStat
```

get_stats fetchs stats about the grid.  

* `status` (string): Node status filter, set to 'up' to get online nodes only.. [optional].  

returns: `GridStat` or `Error`.  

[[Return to contents]](#Contents)

## get_twin_by_account
```v
fn (mut c GridProxyClient) get_twin_by_account(account_id string) ?Twin
```

fetch specific twin information by account.  

* `account_id`: account id.  

returns: `Twin` or `Error`.  

[[Return to contents]](#Contents)

## get_twin_by_id
```v
fn (mut c GridProxyClient) get_twin_by_id(twin_id u64) ?Twin
```

fetch specific twin information by twin id.  

* `twin_id`: twin id.  

returns: `Twin` or `Error`.  

[[Return to contents]](#Contents)

## get_twins
```v
fn (mut c GridProxyClient) get_twins(params TwinFilter) ![]Twin
```

get_twins fetchs twins information with pagaination.  

* `page` (u64): Page number. [optional].  
* `size` (u64): Max result per page. [optional].  
* `ret_count` (bool): Set farms' count on headers based on filter. [optional].  
* `twin_id` (u64): twin id. [optional].  
* `account_id` (string): account address. [optional].  

returns: `[]Twin` or `Error`.  

[[Return to contents]](#Contents)

## get_twins_iterator
```v
fn (mut c GridProxyClient) get_twins_iterator(filter TwinFilter) TwinIterator
```

get_twins_iterator creates an iterator through twin pages with custom filter

[[Return to contents]](#Contents)

## is_pingable
```v
fn (mut c GridProxyClient) is_pingable() !bool
```

is_pingable checks if API server is reachable and responding.  

returns: bool, `true` if API server is reachable and responding, `false` otherwise

[[Return to contents]](#Contents)

#### Powered by vdoc. Generated on: 20 Aug 2023 17:40:51
