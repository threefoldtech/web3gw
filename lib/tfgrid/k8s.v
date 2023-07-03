module tfgrid

[params]
pub struct K8sCluster {
	name          string    [required] // name of the cluster, must be unique
	token         string    [required] // cluster token, workers must have this token to join the cluster
	ssh_key       string    [required] // public ssh key to access the instance in a later stage
	master        K8sNode   [required] // master node configs
	workers       []K8sNode // workers configs
	add_wg_access bool      // if true, adds a wireguard access point to the network
}

[params]
pub struct K8sNode {
	name       string [required] // name of the cluster node
	node_id    u32    // node id to deploy on, if 0, a random eligible node will be selected
	farm_id    u32    // farm id to deploy on, if 0, a random eligible farm will be selected
	public_ip  bool   // if true, a public ipv4 will be added to the node
	public_ip6 bool   // if true, a public ipv6 will be added to the node
	planetary  bool   = true // if true, a yggdrasil ip will be added to the node
	flist      string = 'https://hub.grid.tf/tf-official-apps/threefoldtech-k3s-latest.flist' // flist for kubernetes
	cpu        u32    [required] // number of vcpu cores.
	memory     u32    [required] // node memory in MBs
	disk_size  u32 = 10 // size of disk mounted on the node in GB, monted in /mydisk
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
	], t.timeout)!
}

// Gets a deployed kubernetes cluster data given its name. Returns an error if no cluster can be found
// with the provided name.
pub fn (mut t TFGridClient) k8s_get(get_info GetK8sParams) !K8sClusterResult {
	return t.client.send_json_rpc[[]GetK8sParams, K8sClusterResult]('tfgrid.K8sGet', [
		get_info,
	], t.timeout)!
}

// Deletes a deployed kubernetes cluster given a name. The call returns an error if it fails to do so.
pub fn (mut t TFGridClient) k8s_delete(cluster_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.K8sDelete', [cluster_name],
		t.timeout)!
}

// adds a worker to a deployed kubernetes cluster
pub fn (mut t TFGridClient) k8s_add_worker(params AddK8sWorker) !K8sClusterResult {
	return t.client.send_json_rpc[[]AddK8sWorker, K8sClusterResult]('tfgrid.AddK8sWorker',
		[params], t.timeout)!
}

// remove a worker from a deployed kubernetes cluster
pub fn (mut t TFGridClient) k8s_remove_worker(params RemoveK8sWorker) !K8sClusterResult {
	return t.client.send_json_rpc[[]RemoveK8sWorker, K8sClusterResult]('tfgrid.RemoveK8sWorker',
		[params], t.timeout)!
}
