module gridprocessor

import threefoldtech.threebot.tfgrid { ZDB }
import threefoldtech.threebot.tfgrid.solution { SolutionHandler }
import strconv
import encoding.utf8
import rand

struct ZDBCreateParams {
	ZDB
}

fn (zdb_create ZDBCreateParams) execute(mut s SolutionHandler) !string {
	ret := s.tfclient.zdb_deploy(ZDB{
		node_id: zdb_create.node_id
		name: zdb_create.name
		password: zdb_create.password
		public: zdb_create.public
		size: zdb_create.size
		mode: zdb_create.mode
	})!

	return ret.str()
}

struct ZDBGetParams {
	name string
}

fn (zdb_get ZDBGetParams) execute(mut s SolutionHandler) !string {
	ret := s.tfclient.zdb_get(zdb_get.name)!
	return ret.str()
}

struct ZDBDeleteParams {
	name string
}

fn (zdb_delete ZDBDeleteParams) execute(mut s SolutionHandler) !string {
	s.tfclient.zdb_delete(zdb_delete.name)!
	return 'zdb ${zdb_delete.name} is deleted'
}

fn build_zdb_process(op GridOp, param_map map[string]string, args_set map[string]bool) !(string, Process) {
	match op {
		.create {
			return zdb_create(param_map, args_set)!
		}
		.get {
			return zdb_read(param_map, args_set)!
		}
		.delete {
			return zdb_delete(param_map, args_set)!
		}
		else {
			return error('zdbs do not support updates')
		}
	}
}

fn zdb_create(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	name := param_map['name'] or { utf8.to_lower(rand.string(10)) }

	node_id_str := param_map['node_id'] or { '0' }
	node_id := strconv.parse_uint(node_id_str, 10, 32)!

	password := param_map['password'] or { return error('password is required for zdb ${name}') }

	public := args_set['public']

	size_str := param_map['size'] or { return error('size is required in zdb ${name}') }
	size := strconv.parse_uint(size_str, 10, 32)!

	mode := param_map['mode'] or { 'user' }

	mut zdb := ZDBCreateParams{
		name: name
		node_id: u32(node_id)
		password: password
		public: public
		size: u32(size)
		mode: mode
	}

	return name, zdb
}

fn zdb_read(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	name := param_map['name'] or { return error('zdb name is missing') }

	zdb := ZDBGetParams{
		name: name
	}

	return name, zdb
}

fn zdb_delete(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	name := param_map['name'] or { return error('zdb is missing') }

	zdb := ZDBDeleteParams{
		name: name
	}

	return name, zdb
}
