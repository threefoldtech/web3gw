module tfgrid

import freeflowuniverse.crystallib.actionsparser { Action }
import threefoldtech.threebot.tfgrid { AddK8sWorker, GetK8sParams, RemoveK8sWorker, K8sNode , K8sCluster}
import rand

fn (mut t TFGridHandler) k8s(action Action) ! {
	match action.name {
		'create' {
			name := action.params.get('name')!
			farm_id := action.params.get_int_default('farm_id', 0)!
			capacity := action.params.get_default('capacity', 'small')!
			replica := action.params.get_int_default('replica', 1)!
			public_ip := action.params.get_default_false('add_public_ips')
			ssh_key_name := action.params.get_default('sshkey', 'default')!
			ssh_key := t.get_ssh_key(ssh_key_name)!

			cpu, memory, disk_size := get_k8s_capacity(capacity)!

			mut node := K8sNode{
				name: 'master'
				cpu: cpu
				memory: memory
				disk_size: disk_size
				public_ip: public_ip
			}

			mut workers := []tfgrid.K8sNode{}
			for _ in 0 .. replica {
				mut worker := K8sNode{
					name: 'wr' + rand.string(6)
					cpu: cpu
					memory: memory
					disk_size: disk_size
					public_ip: public_ip
				}

				workers << worker
			}

			cluster := K8sCluster{
				name: name
				token: rand.string(6)
				ssh_key: ssh_key
				master: node
				workers: workers
			}

			deploy_res := t.tfclient.k8s_deploy(cluster)!

			t.logger.info('${deploy_res}')
		}
		'get' {
			name := action.params.get('name')!

			get_res := t.tfclient.k8s_get(GetK8sParams{
				cluster_name: name
				master_name: 'master'
			})!

			t.logger.info('${get_res}')
		}
		'add' {
			name := action.params.get('name')!
			farm_id := action.params.get_int_default('farm_id', 0)!
			capacity := action.params.get_default('capacity', 'small')!
			public_ip := action.params.get_default_false('add_public_ips')
			ssh_key_name := action.params.get_default('sshkey', 'default')!
			ssh_key := t.get_ssh_key(ssh_key_name)!

			cpu, memory, disk_size := get_k8s_capacity(capacity)!

			mut worker := K8sNode{
				name: 'wr' + rand.string(6)
				cpu: cpu
				memory: memory
				disk_size: disk_size
				public_ip: public_ip
			}

			add_res := t.tfclient.k8s_add_worker(AddK8sWorker{
				cluster_name: name
				master_name: 'master'
				worker: worker
			})!

			t.logger.info('${add_res}')
		}
		'remove' {
			name := action.params.get('name')!
			worker_name := action.params.get('worker_name')!

			remove_res := t.tfclient.k8s_remove_worker(RemoveK8sWorker{
				cluster_name: name
				master_name: 'master'
				worker_name: worker_name
			})!
			t.logger.info('${remove_res}')
		}
		'delete' {
			name := action.params.get('name')!

			t.tfclient.k8s_delete(name) or { return error('failed to delete k8s cluster: ${err}') }
		}
		else {
			return error('operation ${action.name} is not supported on k8s')
		}
	}
}

fn get_k8s_capacity(capacity string) !(u32, u32, u32) {
	match capacity {
		'small' {
			return 1, 2048, 10
		}
		'medium' {
			return 2, 4096, 20
		}
		'large' {
			return 8, 8192, 40
		}
		'extra-large' {
			return 8, 16384, 100
		}
		else {
			return error('invalid capacity ${capacity}')
		}
	}
}
