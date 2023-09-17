module tfgrid

import net.http
import json
import x.json2

pub struct GraphQl {
	url string
}

pub struct Contract {
pub:
	contract_id     string [json: contractID]
	deployment_data string [json: deploymentData]
	state           string
	node_id         u32    [json: nodeID]
}

pub struct Contracts {
pub:
	name_contracts []Contract [json: nameContracts]
	node_contracts []Contract [json: nodeContracts]
	rent_contracts []Contract [json: rentContracts]
}

// contractsList, err := c.ListContractsByTwinID([]string{"Created, GracePeriod"})
pub fn (g GraphQl) list_twin_contracts(twin_id u32, states []string) !Contracts {
	state := states.join(', ')
	options := '(where: {twinID_eq: ${twin_id}, state_in: ${state}}, orderBy: twinID_ASC)'
	name_contracts_count := g.get_item_total_count('nameContracts', options)!
	node_contracts_count := g.get_item_total_count('nodeContracts', options)!
	rent_contracts_count := g.get_item_total_count('rentContracts', options)!

	contracts_data := g.query('query getContracts(\$nameContractsCount: Int!, \$nodeContractsCount: Int!, \$rentContractsCount: Int!){
            nameContracts(where: {twinID_eq: ${twin_id}, state_in: ${state}}, limit:  \$nameContractsCount) {
              contractID
              state
              name
            }
            nodeContracts(where: {twinID_eq: ${twin_id}, state_in: ${state}}, limit: \$nodeContractsCount) {
              contractID
              deploymentData
              state
              nodeID
            }
            rentContracts(where: {twinID_eq: ${twin_id}, state_in: ${state}}, limit: \$rentContractsCount) {
              contractID
              state
              nodeID
            }
          }', // map[string]u32{}
	 {
		'nodeContractsCount': node_contracts_count
		'nameContractsCount': name_contracts_count
		'rentContractsCount': rent_contracts_count
	})!

	return json.decode(Contracts, contracts_data.str())!
}

// GetItemTotalCount return count of items
fn (g GraphQl) get_item_total_count(item_name string, options string) !u32 {
	count_body := 'query { items: ${item_name}Connection${options} { count: totalCount } }'
	request_body := {
		'query': count_body
	}
	json_body := json.encode(request_body)

	resp := http.post_json(g.url, json_body)!
	query_data := json2.raw_decode(resp.body)!
	query_map := query_data.as_map()
	data := query_map['data']! as map[string]json2.Any
	items := data['items']! as map[string]json2.Any
	count := u32(items['count']!.int())
	return count
}

struct QueryRequest {
	query     string
	variables map[string]u32
}

// Query queries graphql
fn (g GraphQl) query(body string, variables map[string]u32) !map[string]json2.Any {
	mut request_body := QueryRequest{
		query: body
		variables: variables
	}
	json_body := json.encode(request_body)
	resp := http.post_json(g.url, json_body)!

	query_data := json2.raw_decode(resp.body)!
	data_map := query_data.as_map()
	result := data_map['data']!.as_map()
	return result
}
