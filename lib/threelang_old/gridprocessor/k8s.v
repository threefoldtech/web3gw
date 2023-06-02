module gridprocessor

import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { Capacity, CapacityPackage, SolutionHandler, get_capacity }
import rand

const (
	// TODO: fix to be enum
	k8s_packages = {
		'small':       CapacityPackage{
			cpu: 1
			memory: 2048
			size: 10
		}
		'medium':      CapacityPackage{
			cpu: 2
			memory: 4096
			size: 20
		}
		'large':       CapacityPackage{
			cpu: 4
			memory: 8192
			size: 40
		}
		'extra_large': CapacityPackage{
			cpu: 8
			memory: 16384
			size: 100
		}
	}
)

struct K8sClusterCreate {
	name                      string
	farm_id                   u32
	replica                   u32
	capacity                  string
	ssh_key                   string
	add_wireguard_access      bool
	add_public_ip_to_master   bool
	add_public_ips_to_workers bool
}

fn (k8s K8sClusterCreate) execute(mut s SolutionHandler) !string {
	capacity := gridprocessor.k8s_packages[k8s.capacity]!
	mut public_ip := false
	if k8s.add_public_ip_to_master {
		public_ip = true
	}

	mut node := tfgrid.K8sNode{
		name: 'master'
		cpu: capacity.cpu
		memory: capacity.memory
		disk_size: capacity.size
		public_ip: public_ip
	}

	mut workers := []tfgrid.K8sNode{}
	for _ in 0 .. k8s.replica {
		if k8s.add_public_ips_to_workers {
			public_ip = true
		}

		mut worker := tfgrid.K8sNode{
			name: 'wr' + rand.string(6)
			cpu: capacity.cpu
			memory: capacity.memory
			disk_size: capacity.size
			public_ip: public_ip
		}

		workers << worker
	}

	cluster := tfgrid.K8sCluster{
		name: k8s.name
		token: rand.string(6)
		ssh_key: k8s.ssh_key
		master: node
		workers: workers
	}

	// if k8s.add_wireguard_access {
	// 	cluster.wireguard_access = true
	// }

	println('-------------------------here')
	res := s.tfclient.k8s_deploy(cluster)!
	println('res: '+ res.str())
	return res.str()
}

struct K8sClusterGet {
	name string
}

fn (k8s K8sClusterGet) execute(mut s SolutionHandler) !string {
	res := s.tfclient.k8s_get(tfgrid.GetK8sParams{
		cluster_name: k8s.name
		master_name: 'master'
	})!
	return res.str()
}

struct K8sClusterAddNode {
	name      string
	farm_id   u32
	capacity  string
	ssh_key   string
	public_ip bool
}

fn (k8s K8sClusterAddNode) execute(mut s SolutionHandler) !string {
	capacity := gridprocessor.k8s_packages[k8s.capacity]!
	mut public_ip := false
	if k8s.public_ip {
		public_ip = true
	}

	mut worker := tfgrid.K8sNode{
		name: 'wr' + rand.string(6)
		cpu: capacity.cpu
		memory: capacity.memory
		disk_size: capacity.size
		public_ip: public_ip
	}

	res := s.tfclient.k8s_add_worker(tfgrid.AddK8sWorker{
		cluster_name: k8s.name
		master_name: 'master'
		worker: worker
	})!
	return res.str()
}

struct K8sClusterRemoveNode {
	name        string
	worker_name string
}

fn (k8s K8sClusterRemoveNode) execute(mut s SolutionHandler) !string {
	res := s.tfclient.k8s_remove_worker(tfgrid.RemoveK8sWorker{
		cluster_name: k8s.name
		master_name: 'master'
		worker_name: k8s.worker_name
	})!
	return res.str()
}

struct K8sClusterDelete {
	name string
}

fn (k8s K8sClusterDelete) execute(mut s SolutionHandler) !string {
	s.tfclient.k8s_delete(k8s.name)!
	return 'Cluster has been deleted'
}

fn k8s_deploy(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	name := param_map['name'] or { return error('name is required') }
	farm_id := param_map['farm_id'] or { '0' }
	replica := param_map['replica'] or { '0' }
	capacity := param_map['capacity'] or { return error('capacity is required') }
	ssh_key := param_map['ssh_key'] or { return error('ssh_key is required') }

	add_wireguard_access := args_set['add_wireguard_access']
	add_public_ip_to_master := args_set['add_public_ip_to_master']
	add_public_ips_to_workers := args_set['add_public_ips_to_workers']

	k8s := K8sClusterCreate{
		name: name
		farm_id: farm_id.u32()
		replica: replica.u32()
		capacity: capacity
		ssh_key: ssh_key
		add_wireguard_access: add_wireguard_access
		add_public_ip_to_master: add_public_ip_to_master
		add_public_ips_to_workers: add_public_ips_to_workers
	}

	return name, k8s
}

fn k8s_get(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	name := param_map['name'] or { return error('name is required') }

	k8s := K8sClusterGet{
		name: name
	}

	return name, k8s
}

fn k8s_add_worker(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	name := param_map['name'] or { return error('name is required') }
	farm_id := param_map['farm_id'] or { '0' }
	capacity := param_map['capacity'] or { return error('capacity is required') }
	ssh_key := param_map['ssh_key'] or { return error('ssh_key is required') }
	public_ip := args_set['public_ip']

	k8s := K8sClusterAddNode{
		name: name
		farm_id: farm_id.u32()
		capacity: capacity
		ssh_key: ssh_key
		public_ip: public_ip
	}

	return name, k8s
}

fn k8s_remove_worker(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	name := param_map['name'] or { return error('name is required') }
	worker_name := param_map['worker_name'] or { return error('worker_name is required') }

	k8s := K8sClusterRemoveNode{
		name: name
		worker_name: worker_name
	}

	return name, k8s
}

fn k8s_delete(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	name := param_map['name'] or { return error('name is required') }

	k8s := K8sClusterDelete{
		name: name
	}

	return name, k8s
}

fn build_k8s_process(op GridOp, param_map map[string]string, args_set map[string]bool) !(string, Process) {
	match op {
		.create {
			return k8s_deploy(param_map, args_set)!
		}
		.get {
			return k8s_get(param_map, args_set)!
		}
		.add {
			return k8s_add_worker(param_map, args_set)!
		}
		.remove {
			return k8s_remove_worker(param_map, args_set)!
		}
		.delete {
			return k8s_delete(param_map, args_set)!
		}
		else {
			return error('Invalid operation')
		}
	}
}

