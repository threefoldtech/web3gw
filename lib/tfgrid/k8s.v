module tfgrid

[params]
pub struct K8sCluster {
	name    string    [required]
	token   string    [required]
	ssh_key string    [required]
	master  K8sNode   [required]
	workers []K8sNode
}

[params]
pub struct K8sNode {
	name       string [required]
	node_id    u32
	farm_id    u32
	public_ip  bool
	public_ip6 bool
	planetary  bool   = true
	flist      string = 'https://hub.grid.tf/tf-official-apps/threefoldtech-k3s-latest.flist'
	cpu        u32    [required] // number of vcpu cores.
	memory     u32    [required] // in MBs
	disk_size  u32 = 10 // in GB, monted in /mydisk
}

// GetK8sParams defines the params needed to get a k8s cluster
[params]
pub struct GetK8sParams {
pub:
	cluster_name string // cluster name
	master_name  string // master node's name
}

// AddK8sWorker defines the params needed to add a new worker to an existing cluster
[params]
pub struct AddK8sWorker {
pub:
	worker       K8sNode // the new worker
	cluster_name string  // cluster name
	master_name  string  // master node's name
}

// RemoveK8sWorker defines the params needed to remove a worker from an existing cluster
[params]
pub struct RemoveK8sWorker {
pub:
	cluster_name string // cluster name
	worker_name  string // worker name
	master_name  string // master node's name
}

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
pub fn (mut t TFGridClient) k8s_get(get_info GetK8sParams) !K8sClusterResult {
	return t.client.send_json_rpc[[]GetK8sParams, K8sClusterResult]('tfgrid.K8sGet', [
		get_info,
	], default_timeout)!
}

// Deletes a deployed kubernetes cluster given a name. The call returns an error if it fails to do so.
pub fn (mut t TFGridClient) k8s_delete(cluster_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.K8sDelete', [cluster_name],
		default_timeout)!
}

// adds a worker to a deployed kubernetes cluster
pub fn (mut t TFGridClient) k8s_add_worker(params AddK8sWorker) !K8sClusterResult {
	return t.client.send_json_rpc[[]AddK8sWorker, K8sClusterResult]('tfgrid.AddK8sWorker',
		[params], default_timeout)!
}

// remove a worker from a deployed kubernetes cluster
pub fn (mut t TFGridClient) k8s_remove_worker(params RemoveK8sWorker) !K8sClusterResult {
	return t.client.send_json_rpc[[]RemoveK8sWorker, K8sClusterResult]('tfgrid.RemoveK8sWorker',
		[params], default_timeout)!
}
