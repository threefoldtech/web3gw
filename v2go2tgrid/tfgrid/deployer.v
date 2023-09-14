module tfgrid

import os
import strconv
import json
import time

pub struct Deployer {
pub:
	mnemonics     string
	substrate_url string
	twin_id       u32
	relay_url     string
}

pub enum ChainNetwork {
	dev
	qa
	test
	main
}

const substrate_url = {
	ChainNetwork.dev:  'wss://tfchain.dev.grid.tf/ws'
	ChainNetwork.qa:   'wss://tfchain.qa.grid.tf/ws'
	ChainNetwork.test: 'wss://tfchain.test.grid.tf/ws'
	ChainNetwork.main: 'wss://tfchain.grid.tf/ws'
}

const relay_url = {
	ChainNetwork.dev:  'wss://relay.dev.grid.tf'
	ChainNetwork.qa:   'wss://relay.qa.grid.tf'
	ChainNetwork.test: 'wss://relay.test.grid.tf'
	ChainNetwork.main: 'wss://relay.grid.tf'
}

pub fn get_mnemonics() !string {
	mnemonics := os.getenv('MNEMONICS')
	if mnemonics == '' {
		return error('failed to get mnemonics, run `export MNEMONICS=....`')
	
	}
	return mnemonics
}

pub fn new_deployer(mnemonics string, chain_network ChainNetwork) !Deployer {
	twin_id := get_user_twin(mnemonics, tfgrid.substrate_url[chain_network])!

	return Deployer{
		mnemonics: mnemonics
		substrate_url: tfgrid.substrate_url[chain_network]
		twin_id: twin_id
		relay_url: tfgrid.relay_url[chain_network]
	}
}

pub fn (mut d Deployer) deploy(node_id u32, mut dl Deployment, body string, solution_provider u64) !u64 {
	hash_hex := dl.challenge_hash().hex()
	public_ips := dl.count_public_ips()

	contract_id := d.create_node_contract(node_id, body, hash_hex, public_ips, solution_provider)!
	println('ContractID: ${contract_id}')
	dl.contract_id = contract_id
	signature := d.sign_deployment(hash_hex)!
	dl.add_signature(d.twin_id, signature)
	payload := dl.json_encode()

	node_twin_id := d.get_node_twin(node_id)!
	d.rmb_deployment_deploy(node_twin_id, payload)!
	workload_versions := d.assign_versions(dl)
	d.wait_deployment(node_id, contract_id, workload_versions) or { 
		println("Rolling back...")
		println("deleting contract id: ${contract_id}")
		d.cancel_contract(contract_id) or { return err }
		return err
	 }
	return contract_id
}

pub fn (mut d Deployer) assign_versions(dl Deployment) map[string]u32 {
	mut workload_versions := map[string]u32{}
	for wl in dl.workloads {
		workload_versions[wl.name] = wl.version
	}
	return workload_versions
}

pub fn (mut d Deployer) wait_deployment(node_id u32, contract_id u64, workload_versions map[string]u32) ! {
	start := time.now()
	num_workloads := workload_versions.len
	for {
		mut state_ok := 0
		changes := d.deployment_changes(node_id, contract_id)!
		println('got ${changes.len} workloads')
		for wl in changes {
			println('Workload: ${wl.name}, State: ${wl.result.state}, version: ${wl.version}')
			if wl.version == workload_versions[wl.name] && wl.result.state == result_states.ok {
				state_ok++
			} else if wl.version == workload_versions[wl.name] && wl.result.state == result_states.error {
				return error("failed to deploy deployment due error: ${wl.result.message}")
				
			}
		}
		if state_ok == num_workloads {
			return
		}
		if (time.now() - start).minutes() > 5 {
			return error('failed to deploy deployment: contractID: ${contract_id}, some workloads are not ready after wating 5 minutes')
		} else {
			println('Waiting for deployment to become ready')
			time.sleep(1 * time.second)
		}
	}
}

pub fn (mut d Deployer) get_deployment(contract_id u64, node_id u32) !Deployment {
	twin_id := d.get_node_twin(node_id)!
	payload := {
		'contract_id': contract_id
	}
	res := d.rmb_deployment_get(twin_id, json.encode(payload))!
	return json.decode(Deployment, res)
}

pub fn (mut d Deployer) rmb_deployment_deploy(dst u32, data string) !string {
		res := os.execute("grid-cli rmb-dl-deploy --substrate ${d.substrate_url} --mnemonics \"${d.mnemonics}\" --relay ${d.relay_url} --dst ${dst} --data '${data}'")
	if res.exit_code != 0 {
		return error(res.output)
	}

	return res.output
}

pub fn (mut d Deployer) deployment_changes(node_id u32, contract_id u64) ![]Workload {
	twin_id := d.get_node_twin(node_id)!

	res := d.rmb_deployment_changes(twin_id, contract_id)!
	return json.decode([]Workload, res)
}

pub fn (mut d Deployer) rmb_deployment_changes(dst u32, contract_id u64) !string {
	res := os.execute("grid-cli rmb-dl-changes --substrate ${d.substrate_url} --mnemonics \"${d.mnemonics}\" --relay ${d.relay_url} --dst ${dst} --contract_id '${contract_id}'")
	if res.exit_code != 0 {
		return error(res.output)
	}

	return res.output
}

pub fn (mut d Deployer) rmb_deployment_get(dst u32, data string) !string {
	res := os.execute("grid-cli rmb-dl-get --substrate ${d.substrate_url} --mnemonics \"${d.mnemonics}\" --relay ${d.relay_url} --dst ${dst} --data '${data}'")
	if res.exit_code != 0 {
		return error(res.output)
	}

	return res.output
}

pub fn (mut d Deployer) get_node_twin(node_id u64) !u32 {
	res := os.execute('grid-cli node-twin --substrate ${d.substrate_url}  --node_id ${node_id}')
	if res.exit_code != 0 {
		return error(res.output)
	}

	return u32(strconv.parse_uint(res.output, 10, 32)!)
}

pub fn (mut d Deployer) create_node_contract(node_id u32, body string, hash string, public_ips u32, solution_provider u64) !u64 {
	res := os.execute("grid-cli new-node-cn --substrate ${d.substrate_url} --mnemonics \"${d.mnemonics}\" --node_id ${node_id} --hash \"${hash}\" --public_ips ${public_ips} --solution_provider ${solution_provider}")
	if res.exit_code != 0 {
		return error(res.output)
	}

	return strconv.parse_uint(res.output, 10, 64)!
}

pub fn (mut d Deployer) create_name_contract(name string) !u64 {
	res := os.execute("grid-cli new-name-cn --substrate ${d.substrate_url} --mnemonics \"${d.mnemonics}\" --name ${name}")
	if res.exit_code != 0 {
		return error(res.output)
	}

	return strconv.parse_uint(res.output, 10, 64)!
}

pub fn (mut d Deployer) update_node_contract(contract_id u64, body string, hash string) ! {
	res := os.execute("grid-cli update-cn --substrate ${d.substrate_url} --mnemonics \"${d.mnemonics}\" --contract_id ${contract_id} --body \"${body}\" --hash \"${hash}\"")
	if res.exit_code != 0 {
		return error(res.output)
	}
}

pub fn (mut d Deployer) cancel_contract(contract_id u64) ! {
	res := os.execute("grid-cli cancel-cn --substrate ${d.substrate_url} --mnemonics \"${d.mnemonics}\" --contract_id ${contract_id}")
	if res.exit_code != 0 {
		return error(res.output)
	}
}

pub fn (mut d Deployer) sign_deployment(hash string) !string {
	res := os.execute("grid-cli sign --mnemonics \"${d.mnemonics}\" --hash ${hash}")
	if res.exit_code != 0 {
		return error(res.output)
	}

	return res.output
}

pub fn get_user_twin(mnemonics string, substrate_url string) !u32 {
	res := os.execute("grid-cli user-twin --mnemonics \"${mnemonics}\" --substrate \"${substrate_url}\"")
	if res.exit_code != 0 {
		return error(res.output)
	}

	return u32(strconv.parse_uint(res.output, 10, 32)!)
}

pub fn (mut d Deployer) assign_wg_port(node_id u32) !u16 {
	node_twin := d.get_node_twin(node_id)!
	res := os.execute("grid-cli rmb-taken-ports --substrate ${d.substrate_url} --mnemonics \"${d.mnemonics}\" --relay ${d.relay_url} --dst ${node_twin} ")
	if res.exit_code != 0 {
		return error(res.output)
	}

	taken_ports := json.decode([]u16,res.output) or {
		return error("can't parse node taken ports: ${err}")
	}
	port := rand_port(taken_ports) or { 
		return error("can't assign wireguard port: ${err}")
	 }

	return port
}

pub fn(mut d Deployer) get_node_pub_config(node_id u32) !PublicConfig {
	node_twin :=  d.get_node_twin(node_id)!
	res := os.execute("grid-cli rmb-node-pubConfig --substrate ${d.substrate_url} --mnemonics \"${d.mnemonics}\" --relay ${d.relay_url} --dst ${node_twin} ")
	if res.exit_code != 0 {
		return error(res.output)
	}

	public_config := json.decode(PublicConfig,res.output) or {
		return err
	}


	return public_config
}