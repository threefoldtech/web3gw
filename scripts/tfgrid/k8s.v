module tfgrid

// Deploys a kubernetes cluster given the cluster configuration. The cluster object is returned with extra
// data if the call succeeds. 
pub fn (mut t TFGridClient) k8s_deploy(cluster K8sCluster) !K8sClusterResult {
	return t.client.send_json_rpc[[]K8sCluster, K8sClusterResult]('tfgrid.K8sDeploy', [
		cluster,
	], default_timeout)!
}

// Gets a deployed kubernetes cluster data given its name. Returns an error if no cluster can be found 
// with the provided name. 
pub fn (mut t TFGridClient) k8s_get(cluster_name string) !K8sClusterResult {
	return t.client.send_json_rpc[[]string, K8sClusterResult]('tfgrid.K8sGet', [
		cluster_name,
	], default_timeout)!
}

// Deletes a deployed kubernetes cluster given a name. The call returns an error if it fails to do so. 
pub fn (mut t TFGridClient) k8s_delete(cluster_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.K8sDelete', [cluster_name], default_timeout)!
}

// NOTE: not implemented
// pub fn k8s_add_node(mut client RpcWsClient, params AddK8sNode) !K8sClusterResult {
// 	return t.client.send_json_rpc[[]AddK8sNode, K8sClusterResult]('tfgrid.K8sNodeAdd', [params], default_timeout)!
// }

// pub fn k8s_remove_node(mut client RpcWsClient, params RemoveK8sNode) !K8sClusterResult {
// 	return t.client.send_json_rpc[[]RemoveK8sNode, K8sClusterResult]('tfgrid.K8sNodeRemove', [params], default_timeout)!
// }
