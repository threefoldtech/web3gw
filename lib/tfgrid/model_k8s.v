module tfgrid

pub struct K8sClusterResult {
	name    string
	token   string
	ssh_key string
	master  K8sNode
	workers []K8sNode
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
