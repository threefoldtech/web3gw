module zos

import os
import strconv
import freeflowuniverse.crystallib.threefold.rmb
import json
import encoding.base64

pub struct Deployer {
pub:
	mnemonics     string
	substrate_url string
	twin_id       u32
pub mut:
	rmb_cl rmb.RMBClient
}

pub fn (mut d Deployer) deploy(node_id u32, mut dl Deployment, body string, solution_provider u64) !u64 {
	hash_hex := dl.challenge_hash().hex()
	public_ips := dl.count_public_ips()

	contract_id := d.create_node_contract(node_id, body, hash_hex, public_ips, solution_provider)!
	dl.contract_id = contract_id

	signature := d.sign_deployment(hash_hex)!
	dl.add_signature(d.twin_id, signature)
	payload := dl.json_encode()

	node_twin_id := d.get_node_twin(node_id)!

	res := d.rmb_cl.rmb_request('zos.deployment.deploy', node_twin_id, payload)!
	if res.err.code != 0 {
		return error('an error occured while trying to deploy to the node: ${res.err.message}')
	}
	return contract_id
}

pub fn (mut d Deployer) get_deployment(contract_id u64, node_id u32) !Deployment {
	twin_id := d.get_node_twin(node_id)!
	payload := {
		'contract_id': contract_id
	}

	res := d.rmb_cl.rmb_request('zos.deployment.get', twin_id, json.encode(payload))!
	if res.err.code != 0 {
		return error('an error occured while trying to deploy to the node: ${res.err.message}')
	}

	decoded_data := base64.decode_str(res.dat)

	return json.decode(Deployment, decoded_data)
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
