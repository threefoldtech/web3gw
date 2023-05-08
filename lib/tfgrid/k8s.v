module tfgrid

// Deploys a kubernetes cluster given the cluster configuration. The cluster object is returned with extra
// data if the call succeeds.
pub fn (mut t TFGridClient) k8s_deploy(cluster K8sCluster) !K8sClusterResult {
	return t.client.send_json_rpc[[]K8sCluster, K8sClusterResult]('tfgrid.K8sDeploy',
		[
		cluster,
	], default_timeout)!
}

// Gets a deployed kubernetes cluster data given its name. Returns an error if no cluster can be found
// with the provided name.
pub fn (mut t TFGridClient) k8s_get(get_info GetK8sInfo) !K8sClusterResult {
	return t.client.send_json_rpc[[]GetK8sInfo, K8sClusterResult]('tfgrid.K8sGet', [
		get_info,
	], default_timeout)!
}

// Deletes a deployed kubernetes cluster given a name. The call returns an error if it fails to do so.
pub fn (mut t TFGridClient) k8s_delete(cluster_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.K8sDelete', [cluster_name],
		default_timeout)!
}

// adds a worker to a deployed kubernetes cluster
pub fn (mut t TFGridClient)k8s_add_worker(params AddK8sWorker) ! {
	_ := t.client.send_json_rpc[[]AddK8sWorker, string]('tfgrid.AddK8sWorker', [params], default_timeout)!
}

// remove a worker from a deployed kubernetes cluster
pub fn (mut t TFGridClient)k8s_remove_worker(params RemoveK8sWorker) ! {
	_ := t.client.send_json_rpc[[]RemoveK8sWorker, string]('tfgrid.RemoveK8sWorker', [params], default_timeout)!
}
