module main

import threefoldtech.threebot.tfgrid
import log

fn test_k8s_ops(mut client tfgrid.TFGridClient, mut logger log.Logger) ! {
	cluster_name := 'testK8sOps'

	mut res := client.k8s_deploy(tfgrid.K8sCluster{
		name: cluster_name
		token: 'token6'
		ssh_key: 'SSH-Key'
		master: tfgrid.K8sNode{
			name: 'master'
			node_id: 2
			cpu: 1
			memory: 1024
		}
		workers: [
			tfgrid.K8sNode{
				name: 'w1'
				node_id: 2
				cpu: 1
				memory: 1024
			},
		]
	})!
	logger.info('${res}')

	defer {
		client.k8s_delete(cluster_name) or { logger.error('failed to delete cluster: ${err}') }
	}

	res = client.k8s_get(tfgrid.GetK8sParams{
		cluster_name: cluster_name
		master_name: 'master'
	})!
	logger.info('${res}')

	res = client.k8s_add_worker(tfgrid.AddK8sWorker{
		cluster_name: cluster_name
		worker: tfgrid.K8sNode{
			name: 'w3'
			node_id: 3
			cpu: 1
			memory: 1024
		}
		master_name: 'master'
	})!
	logger.info('${res}')

	res = client.k8s_remove_worker(tfgrid.RemoveK8sWorker{
		cluster_name: cluster_name
		worker_name: 'w1'
		master_name: 'master'
	})!
	logger.info('${res}')
}

fn main() {
	mut logger := log.Log{
		level: .info
	}

	mut tfgrid_client, _ := tfgrid.cli(mut logger) or {
		logger.error('failed to initialize tfgrid client: ${err}')
		exit(1)
	}

	test_k8s_ops(mut tfgrid_client, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
