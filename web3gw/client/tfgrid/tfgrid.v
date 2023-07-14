module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import threefoldtech.web3gw.tfgrid.applications

// TFGridClient is a client containig an RpcWsClient instance, and implements all tfgrid functionality
[openrpc: exclude]
[noinit]
pub struct TFGridClient {
mut:
	client  &RpcWsClient
	timeout int
}

[openrpc: exclude]
pub fn new(mut client RpcWsClient) TFGridClient {
	return TFGridClient{
		client: &client
		timeout: 500000
	}
}

// todo
pub fn (mut t TFGridClient) load(args Load) ! {
	_ := t.client.send_json_rpc[[]Load, string]('tfgrid.Load', [args], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) deploy_vm(args DeployVM) !VMDeployment {
	return t.client.send_json_rpc[[]DeployVM, VMDeployment]('tfgrid.DeployVM', [args], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) get_vm_deployment(name string) !VMDeployment {
	return t.client.send_json_rpc[[]string, VMDeployment]('tfgrid.GetVMDeployment', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) cancel_vm_deployment(name string) ! {
	_ :=  t.client.send_json_rpc[[]string, VMResult]('tfgrid.CancelVMDeployment', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) deploy_network(args NetworkDeployment) !NetworkDeployment {
	return t.client.send_json_rpc[[]NetworkDeployment, NetworkDeployment]('tfgrid.DeployNetwork',
		[args], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) get_network_deployment(name string) !NetworkDeployment {
	return t.client.send_json_rpc[[]string, NetworkDeployment]('tfgrid.GetNetworkDeployment',
		[name], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) cancel_network_deployment(name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.CancelNetworkDeployment', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) add_vm_to_network_deployment(args AddVMToNetworkDeployment) !NetworkDeployment {
	return t.client.send_json_rpc[[]AddVMToNetworkDeployment, NetworkDeployment]('tfgrid.AddVMToNetworkDeployment',
		[args], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) remove_vm_from_network_deployment(args RemoveVMFromNetworkDeployment) !NetworkDeployment {
	return t.client.send_json_rpc[[]RemoveVMFromNetworkDeployment, NetworkDeployment]('tfgrid.RemoveVMFromNetworkDeployment',
		[args], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) deploy_k8s_cluster(args K8sCluster) !K8sCluster {
	return t.client.send_json_rpc[[]K8sCluster, K8sCluster]('tfgrid.DeployK8sCluster',
		[args], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) get_k8s_cluster(name string) !K8sCluster {
	return t.client.send_json_rpc[[]string, K8sCluster]('tfgrid.GetK8sCluster', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) cancel_k8s_cluster(name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.CancelK8SCluster', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) add_worker_to_k8s_cluster(args AddWorkerToK8sCluster) !K8sCluster {
	return t.client.send_json_rpc[[]AddWorkerToK8sCluster, K8sCluster]('tfgrid.AddWorkerToK8sCluster', [
		args,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) remove_worker_from_k8s_cluster(args RemoveWorkerFromK8sCluster) !K8sCluster {
	return t.client.send_json_rpc[[]RemoveWorkerFromK8sCluster, K8sCluster]('tfgrid.RemoveWorkerFromK8sCluster', [
		args,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) deploy_zdb(args ZDBDeployment) !ZDBDeployment {
	return t.client.send_json_rpc[[]ZDBDeployment, ZDBDeployment]('tfgrid.DeployZDB', [
		args,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) get_zdb_deployment(name string) !ZDBDeployment {
	return t.client.send_json_rpc[[]string, ZDBDeployment]('tfgrid.GetZDBDeployment', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) cancel_zdb_deployment(name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.CancelZDBDeployment', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) deploy_gateway_name(args GatewayName) !GatewayName {
	return t.client.send_json_rpc[[]GatewayName, GatewayName]('tfgrid.DeployGatewayName', [
		args,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) get_gateway_name(name string) !GatewayName {
	return t.client.send_json_rpc[[]string, GatewayName]('tfgrid.GetGatewayName', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) cancel_gateway_name(name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.CancelGatewayName', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) deploy_gateway_fqdn(args GatewayFQDN) !GatewayFQDN {
	return t.client.send_json_rpc[[]GatewayFQDN, GatewayFQDN]('tfgrid.DeployGatewayFQDN', [
		args,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) get_gateway_fqdn(name string) !GatewayFQDN {
	return t.client.send_json_rpc[[]string, GatewayFQDN]('tfgrid.GetGatewayFQDN', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) cancel_gateway_fqdn(name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.CancelGatewayFQDN', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) find_nodes(args FindNodes) !NodesResult {
	return t.client.send_json_rpc[[]FindNodes, NodesResult]('tfgrid.FindNodes', [
		args,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) find_farms(args FindFarms) !FarmsResult {
	return t.client.send_json_rpc[[]FindFarms, FarmsResult]('tfgrid.FindFarms', [
		args,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) find_contracts(args FindContracts) !ContractsResult {
	return t.client.send_json_rpc[[]FindContracts, ContractsResult]('tfgrid.FindContracts', [
		args,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) find_twins(args FindTwins) !TwinsResult {
	return t.client.send_json_rpc[[]FindTwins, TwinsResult]('tfgrid.FindTwins', [
		args,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) statistics(args GetStatistics) !Statistics {
	return t.client.send_json_rpc[[]GetStatistics, Statistics]('tfgrid.Statistics', [
		args,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) deploy_discourse(args DeployDiscourse) !DiscourseDeployment {
	return t.client.send_json_rpc[[]DeployDiscourse, DiscourseDeployment]('tfgrid.DeployDiscourse', [
		args,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) get_discourse_deployment(name string) !DiscourseDeployment {
	return t.client.send_json_rpc[[]string, DiscourseDeployment]('tfgrid.GetDiscourseDeployment', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) cancel_discourse_deployment(name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.CancelDiscourceDeployment', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) deploy_funkwhale(args DeployFunkwhale) !FunkwhaleDeployment {
	return t.client.send_json_rpc[[]DeployFunkwhale, FunkwhaleDeployment]('tfgrid.DeployFunkwhale', [
		args,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) get_funkwhale_deployment(name string) !FunkwhaleDeployment {
	return t.client.send_json_rpc[[]string, FunkwhaleDeployment]('tfgrid.GetFunkwhaleDeployment', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) cancel_funkwhale_deployment(name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.CancelFunkwhaleDeployment', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) deploy_peertube(args DeployPeertube) !PeertubeDeployment {
	return t.client.send_json_rpc[[]DeployPeertube, PeertubeDeployment]('tfgrid.DeployPeertube', [
		args,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) get_peertube_deployment(name string) !PeertubeDeployment {
	return t.client.send_json_rpc[[]string, PeertubeDeployment]('tfgrid.GetPeertubeDeployment', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) cancel_peertube_deployment(name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.CancelPeertubeDeployment', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) deploy_presearch(args DeployPresearch) !PresearchDeployment {
	return t.client.send_json_rpc[[]DeployPresearch, PresearchDeployment]('tfgrid.DeployPresearch', [
		args,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) get_presearch_deployment(name string) !PresearchDeployment {
	return t.client.send_json_rpc[[]string, PresearchDeployment]('tfgrid.GetPresearchDeployment', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) cancel_presearch_deployment(name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.CancelPresearchDeployment', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) deploy_taiga(args DeployTaiga) !TaigaDeployment {
	return t.client.send_json_rpc[[]DeployTaiga, TaigaDeployment]('tfgrid.DeployTaiga', [
		args,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) get_taiga_deployment(name string) !TaigaDeployment {
	return t.client.send_json_rpc[[]string, TaigaDeployment]('tfgrid.GetTaigaDeployment', [
		name,
	], t.timeout)!
}

// todo
pub fn (mut t TFGridClient) cancel_taiga_deployment(name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.CancelTaigaDeployment', [
		name,
	], t.timeout)!
}