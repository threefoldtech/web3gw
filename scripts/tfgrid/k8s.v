module tfgrid

// k8s_deploy deploys kubernetes cluster
// - cluster: the kubernetes cluster model
// returns the whole cluster model with the computed files from the grid
pub fn (mut client TFGridClient) k8s_deploy(cluster K8sCluster) !K8sClusterResult {
	return client.send_json_rpc[[]K8sCluster, K8sClusterResult]('tfgrid.K8sDeploy', [
		cluster,
	], default_timeout)!
}

// k8s_get get a deployed kubernetes cluster info
// - cluster_name: name of the cluster
// returns the cluster model info
pub fn (mut client TFGridClient) k8s_get(cluster_name string) !K8sClusterResult {
	return client.send_json_rpc[[]string, K8sClusterResult]('tfgrid.K8sGet', [
		cluster_name,
	], default_timeout)!
}

// k8s_delete delete a deployed kubernetes cluster
// - cluster_name: name of the cluster
pub fn (mut client TFGridClient) k8s_delete(cluster_name string) ! {
	_ := client.send_json_rpc[[]string, string]('tfgrid.K8sDelete', [cluster_name], default_timeout)!
}

// NOTE: not implemented
// pub fn k8s_add_node(mut client RpcWsClient, params AddK8sNode) !K8sClusterResult {
// 	return client.send_json_rpc[[]AddK8sNode, K8sClusterResult]('tfgrid.K8sNodeAdd', [params], default_timeout)!
// }

// pub fn k8s_remove_node(mut client RpcWsClient, params RemoveK8sNode) !K8sClusterResult {
// 	return client.send_json_rpc[[]RemoveK8sNode, K8sClusterResult]('tfgrid.K8sNodeRemove', [params], default_timeout)!
// }
