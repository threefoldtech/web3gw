module tfgrid

pub struct K8sClusterResult {
	name                 string    // cluster name
	token                string    // cluster token, workers must have this token to join the cluster
	ssh_key              string    // public ssh key to access the instance in a later stage
	master               K8sNode   // master configs
	workers              []K8sNode // workers configs
	add_wireguard_access bool      // if true, adds a wireguard access point to the network
}

pub struct K8sNodeResult {
	name       string // name of the cluster node
	node_id    u32    // node id that this node was deployed on
	farm_id    u32    // farm id that this node was deployed on
	public_ip  bool   // whether or not a public ipv4 was added to this node
	public_ip6 bool   // whether or not a public ipv6 was added to this node
	planetary  bool   // whether or not a yggdrasil ip was added to this node
	flist      string // flist of this node
	cpu        u32    // number of vcpu cores
	memory     u32    // node memory in MBs
	disk_size  u32    // size of disk mounted on this node in GB
	// computed
	computed_ip4 string // public ipv4 attached to this node, if any
	computed_ip6 string // public ipv6 attached to this node, if any
	wg_ip        string // wireguard private ip of this node
	ygg_ip       string // ygg ip attached to this node, if any
}
