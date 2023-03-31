module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }


pub fn k8s_deploy(mut client RpcWsClient, params K8sCluster) !K8sClusterResult {
	return client.send_json_rpc[K8sCluster, K8sClusterResult]('tfgrid.k8s.deploy', params, default_timeout)!
}

pub fn k8s_delete(mut client RpcWsClient, params string) ! {
	_ := client.send_json_rpc[string, string]('tfgrid.k8s.delete', params, default_timeout)!
}

pub fn k8s_get(mut client RpcWsClient, params string) !K8sClusterResult {
	return client.send_json_rpc[string, K8sClusterResult]('tfgrid.k8s.get', params, default_timeout)!
}

pub fn k8s_add_node(mut client RpcWsClient, params AddK8sNode) !K8sClusterResult {
	return client.send_json_rpc[AddK8sNode, K8sClusterResult]('tfgrid.k8s.node.add', params, default_timeout)!
}

pub fn k8s_remove_node(mut client RpcWsClient, params RemoveK8sNode) !K8sClusterResult {
	return client.send_json_rpc[RemoveK8sNode, K8sClusterResult]('tfgrid.k8s.node.remove', params, default_timeout)!
}
