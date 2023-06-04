module solution

import threefoldtech.threebot.tfgrid 
import rand

pub struct K8s {
pub mut:
	name string
	farm_id int
	capacity string
	replica int
	wg bool
	public_ip bool
	ssh_key string
}

const (
		k8s_cap = {
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

pub fn (mut s SolutionHandler) create_k8s(k8s K8s) !tfgrid.K8sClusterResult {
	capacity := k8s_cap[k8s.capacity]!

	mut node := tfgrid.K8sNode{
		name: 'master'
		cpu: capacity.cpu
		memory: capacity.memory
		disk_size: capacity.size
		public_ip: k8s.public_ip
	}

	mut workers := []tfgrid.K8sNode{}
	for _ in 0 .. k8s.replica {
		// if k8s.add_public_ips_to_workers {
		// 	public_ip = true
		// }

		mut worker := tfgrid.K8sNode{
			name: 'wr' + rand.string(6)
			cpu: capacity.cpu
			memory: capacity.memory
			disk_size: capacity.size
			public_ip: k8s.public_ip
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

	res := s.tfclient.k8s_deploy(cluster)!
	return res

}

pub fn (mut s SolutionHandler) add_k8s_worker(k8s K8s) !tfgrid.K8sClusterResult {
	capacity := k8s_cap[k8s.capacity]!

	mut worker := tfgrid.K8sNode{
		name: 'wr' + rand.string(6)
		cpu: capacity.cpu
		memory: capacity.memory
		disk_size: capacity.size
	}

	res := s.tfclient.k8s_add_worker(tfgrid.AddK8sWorker{
		cluster_name: k8s.name
		master_name: 'master'
		worker: worker
	})!
	return res
}

pub fn (mut s SolutionHandler) remove_k8s_worker(cluster_name string, worker_name string) !tfgrid.K8sClusterResult {
	res := s.tfclient.k8s_remove_worker(tfgrid.RemoveK8sWorker{
		cluster_name: cluster_name
		master_name: 'master'
		worker_name: worker_name
	})!
	return res
}

pub fn (mut s SolutionHandler) get_k8s(cluster_name string) !tfgrid.K8sClusterResult {
	return s.tfclient.k8s_get(tfgrid.GetK8sParams{
		cluster_name: cluster_name
		master_name: 'master'
	})
}

pub fn (mut s SolutionHandler) delete_k8s(cluster_name string) ! {
	return s.tfclient.k8s_delete(cluster_name)
}