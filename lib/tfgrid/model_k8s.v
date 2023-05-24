module tfgrid

[params]
pub struct K8sCluster {
	name    string    [required]
	token   string    [required]
	ssh_key string    [required]
	master  K8sNode   [required]
	workers []K8sNode
}

pub struct K8sClusterResult {
	name    string
	token   string
	ssh_key string
	master  K8sNode
	workers []K8sNode
}

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

pub struct K8sNodeResult {
	name       string
	node_id    u32
	farm_id    u32
	public_ip  bool
	public_ip6 bool
	planetary  bool
	flist      string
	cpu        u32
	memory     u32
	disk_size  u32
	// computed
	computed_ip4 string
	computed_ip6 string
	wg_ip        string
	ygg_ip       string
}

// GetK8sParams defines the params needed to get a k8s cluster
pub struct GetK8sParams {
pub:
	cluster_name string // cluster name
	master_name  string // master node's name
}

// AddK8sWorker defines the params needed to add a new worker to an existing cluster
pub struct AddK8sWorker {
pub:
	worker       K8sNode // the new worker
	cluster_name string  // cluster name
	master_name  string  // master node's name
}

// RemoveK8sWorker defines the params needed to remove a worker from an existing cluster
pub struct RemoveK8sWorker {
pub:
	cluster_name string // cluster name
	worker_name  string // worker name
	master_name  string // master node's name
}
