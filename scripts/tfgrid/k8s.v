module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

struct K8sClient {
	RpcWsClient
}

pub fn (mut client K8sClient) k8s_deploy(cluster K8sCluster) !K8sClusterResult {
	return client.send_json_rpc[[]K8sCluster, K8sClusterResult]('tfgrid.K8sDeploy', [
		cluster,
	], default_timeout)!
}

pub fn (mut client K8sClient) k8s_get(cluster_name string) !K8sClusterResult {
	return client.send_json_rpc[[]string, K8sClusterResult]('tfgrid.K8sGet', [
		cluster_name,
	], default_timeout)!
}

pub fn (mut client K8sClient) k8s_delete(cluster_name string) ! {
	_ := client.send_json_rpc[[]string, string]('tfgrid.K8sDelete', [cluster_name], default_timeout)!
}

// NOTE: not implemented
// pub fn k8s_add_node(mut client RpcWsClient, params AddK8sNode) !K8sClusterResult {
// 	return client.send_json_rpc[[]AddK8sNode, K8sClusterResult]('tfgrid.K8sNodeAdd', [params], default_timeout)!
// }

// pub fn k8s_remove_node(mut client RpcWsClient, params RemoveK8sNode) !K8sClusterResult {
// 	return client.send_json_rpc[[]RemoveK8sNode, K8sClusterResult]('tfgrid.K8sNodeRemove', [params], default_timeout)!
// }
