package tfgrid

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/workloads"
	"github.com/threefoldtech/zos/pkg/gridtypes"
)

// K8sCluster struct for k8s cluster
type K8sCluster struct {
	Name        string    `json:"name"`
	Master      *K8sNode  `json:"master"`
	Workers     []K8sNode `json:"workers"`
	Token       string    `json:"token"`
	NetworkName string    `json:"network_name"`
	SSHKey      string    `json:"ssh_key"`
	AddWGAccess bool      `json:"add_wg_access"`
}

// K8sNode kubernetes data
type K8sNode struct {
	Name      string `json:"name"`
	NodeID    uint32 `json:"node_id"`
	FarmID    uint32 `json:"farm_id"`
	DiskSize  int    `json:"disk_size"`
	PublicIP  bool   `json:"public_ip"`
	PublicIP6 bool   `json:"public_ip6"`
	Planetary bool   `json:"planetary"`
	Flist     string `json:"flist"`
	CPU       int    `json:"cpu"`
	Memory    int    `json:"memory"`

	// computed
	ComputedIP4 string `json:"computed_ip4"`
	ComputedIP6 string `json:"computed_ip6"`
	WGIP        string `json:"wg_ip"`
	YggIP       string `json:"ygg_ip"`
}

type GetClusterParams struct {
	ClusterName string `json:"cluster_name"`
	MasterName  string `json:"master_name"`
}

type AddWorkerParams struct {
	ClusterName string  `json:"cluster_name"`
	MasterName  string  `json:"master_name"`
	Worker      K8sNode `json:"worker"`
}

type RemoveWorkerParams struct {
	ClusterName string `json:"cluster_name"`
	MasterName  string `json:"master_name"`
	WorkerName  string `json:"worker_name"`
}

// K8sDeploy deploys a kubernetes cluster
func (r *Client) K8sDeploy(ctx context.Context, cluster K8sCluster) (K8sCluster, error) {
	projectName := generateProjectName(cluster.Name)

	// validate project name is unique
	if err := r.validateProjectName(ctx, projectName); err != nil {
		return K8sCluster{}, err
	}

	if err := r.assignNodesIDsForCluster(ctx, &cluster); err != nil {
		return K8sCluster{}, errors.Wrapf(err, "Couldn't find node for all cluster nodes")
	}

	// deploy network
	nodes := []uint32{cluster.Master.NodeID}
	for _, worker := range cluster.Workers {
		nodes = append(nodes, worker.NodeID)
	}

	znet, err := r.deployNetwork(ctx, cluster.Name, nodes, "10.1.0.0/16", false, projectName)
	if err != nil {
		return K8sCluster{}, errors.Wrap(err, "failed to deploy network")
	}

	cluster.NetworkName = znet.Name

	// map to workloads.k8sCluster
	k8s := newK8sClusterFromModel(cluster)

	// Deploy workload
	if err := r.client.DeployK8sCluster(ctx, &k8s); err != nil {
		return K8sCluster{}, errors.Wrapf(err, "Failed to deploy K8s Cluster")
	}

	cluster.Master.assignComputedNodeValues(*k8s.Master)
	for idx := range k8s.Workers {
		cluster.Workers[idx].assignComputedNodeValues(k8s.Workers[idx])
	}

	return cluster, nil
}

// K8sDelete deletes a kubernetes cluster specified by the cluster name
func (r *Client) K8sDelete(ctx context.Context, clusterName string) error {
	projectName := generateProjectName(clusterName)

	err := r.client.CancelProject(ctx, projectName)
	if err != nil {
		return errors.Wrapf(err, "failed to cancel project: %s", projectName)
	}

	return nil
}

// K8sGet retreives a kubernetes cluster specified by the cluster name
func (c *Client) K8sGet(ctx context.Context, params GetClusterParams) (K8sCluster, error) {
	clusterContracts, err := c.loadModelContracts(ctx, params.ClusterName)
	if err != nil {
		return K8sCluster{}, errors.Wrapf(err, "failed to get cluster %s contracts", params.ClusterName)
	}

	if len(clusterContracts.nodeContarcts) == 0 {
		return K8sCluster{}, fmt.Errorf("found 0 contracts for cluster %s", params.ClusterName)
	}

	cluster, err := c.loadK8s(params.MasterName, clusterContracts.nodeContarcts)
	if err != nil {
		return K8sCluster{}, errors.Wrapf(err, "failed to load kubernetes cluster %s", params.ClusterName)
	}

	nodeFarms, err := getNodeFarmsIDs(c.client, &cluster)
	if err != nil {
		return K8sCluster{}, errors.Wrapf(err, "failed to get node farms ids")
	}

	ret := k8sClusterToModel(cluster, params.ClusterName, nodeFarms)

	return ret, nil
}

func getNodeFarmsIDs(c TFGridClient, cluster *workloads.K8sCluster) (map[uint32]uint32, error) {
	nodeFarms := map[uint32]uint32{}

	farm, err := c.GetNodeFarm(cluster.Master.Node)
	if err != nil {
		return nil, err
	}

	nodeFarms[cluster.Master.Node] = farm

	for _, w := range cluster.Workers {
		farm, err := c.GetNodeFarm(w.Node)
		if err != nil {
			return nil, err
		}

		nodeFarms[w.Node] = farm
	}

	return nodeFarms, nil
}

func k8sClusterToModel(cluster workloads.K8sCluster, clusterName string, nodeFarms map[uint32]uint32) K8sCluster {
	master := k8sNodeToModel(*cluster.Master, nodeFarms)
	workers := []K8sNode{}
	for _, worker := range cluster.Workers {
		workers = append(workers, k8sNodeToModel(worker, nodeFarms))
	}

	return K8sCluster{
		Name:        clusterName,
		Master:      &master,
		Workers:     workers,
		Token:       cluster.Token,
		NetworkName: cluster.NetworkName,
		SSHKey:      cluster.SSHKey,
	}
}

func k8sNodeToModel(node workloads.K8sNode, nodeFarms map[uint32]uint32) K8sNode {
	return K8sNode{
		Name:        node.Name,
		NodeID:      node.Node,
		FarmID:      nodeFarms[node.Node],
		DiskSize:    node.DiskSize,
		PublicIP:    node.PublicIP,
		PublicIP6:   node.PublicIP6,
		Planetary:   node.Planetary,
		Flist:       node.Flist,
		CPU:         node.CPU,
		Memory:      node.Memory,
		ComputedIP4: node.ComputedIP,
		ComputedIP6: node.ComputedIP6,
		WGIP:        node.IP,
		YggIP:       node.YggIP,
	}
}

func newClientK8sNodeFromK8sNode(k8sNode K8sNode) workloads.K8sNode {
	return workloads.K8sNode{
		Name:      k8sNode.Name,
		Node:      k8sNode.NodeID,
		DiskSize:  k8sNode.DiskSize,
		PublicIP:  k8sNode.PublicIP,
		PublicIP6: k8sNode.PublicIP6,
		Planetary: k8sNode.Planetary,
		Flist:     k8sNode.Flist,
		CPU:       k8sNode.CPU,
		Memory:    k8sNode.Memory,
	}
}

func (k *K8sNode) assignComputedNodeValues(node workloads.K8sNode) {
	k.ComputedIP4 = node.ComputedIP
	k.ComputedIP6 = node.ComputedIP6
	k.WGIP = node.IP
	k.YggIP = node.YggIP
}

func newK8sClusterFromModel(model K8sCluster) workloads.K8sCluster {
	master := newK8sNodeFromModel(*model.Master)
	workers := []workloads.K8sNode{}
	for _, w := range model.Workers {
		workers = append(workers, newK8sNodeFromModel(w))
	}

	return workloads.K8sCluster{
		Master:       &master,
		Workers:      workers,
		Token:        model.Token,
		NetworkName:  model.NetworkName,
		SolutionType: generateProjectName(model.Name),
		SSHKey:       model.SSHKey,
	}
}

func newK8sNodeFromModel(model K8sNode) workloads.K8sNode {
	return workloads.K8sNode{
		Name:        model.Name,
		Node:        model.NodeID,
		DiskSize:    model.DiskSize,
		PublicIP:    model.PublicIP,
		PublicIP6:   model.PublicIP6,
		Planetary:   model.Planetary,
		Flist:       model.Flist,
		ComputedIP:  model.ComputedIP4,
		ComputedIP6: model.ComputedIP6,
		YggIP:       model.YggIP,
		IP:          model.WGIP,
		CPU:         model.CPU,
		Memory:      model.Memory,
	}
}

// Assign chosen NodeIds to cluster node. with both way conversions to/from Reservations array.
func (r *Client) assignNodesIDsForCluster(ctx context.Context, cluster *K8sCluster) error {
	// all units unified in bytes

	workloads := []*PlannedReservation{}

	ms := PlannedReservation{
		WorkloadName: cluster.Master.Name,
		FarmID:       cluster.Master.FarmID,
		MRU:          uint64(cluster.Master.Memory * int(gridtypes.Megabyte)),
		SRU:          uint64(cluster.Master.DiskSize * int(gridtypes.Gigabyte)),
		PublicIps:    cluster.Master.PublicIP,
		NodeID:       cluster.Master.NodeID,
	}

	workloads = append(workloads, &ms)

	for idx := range cluster.Workers {

		wr := PlannedReservation{
			WorkloadName: cluster.Workers[idx].Name,
			FarmID:       cluster.Workers[idx].FarmID,
			MRU:          uint64(cluster.Workers[idx].Memory * int(gridtypes.Megabyte)),
			SRU:          uint64(cluster.Workers[idx].DiskSize * int(gridtypes.Gigabyte)),
			PublicIps:    cluster.Workers[idx].PublicIP,
			NodeID:       cluster.Workers[idx].NodeID,
		}

		workloads = append(workloads, &wr)
	}

	err := r.AssignNodes(ctx, workloads)
	if err != nil {
		return err
	}

	if cluster.Master.NodeID == 0 {
		for _, workload := range workloads {
			if workload.WorkloadName == cluster.Master.Name {
				cluster.Master.NodeID = uint32(workload.NodeID)
			}
		}
	}

	for idx := range cluster.Workers {
		if cluster.Workers[idx].NodeID == 0 {
			for _, workload := range workloads {
				if workload.WorkloadName == cluster.Workers[idx].Name {
					cluster.Workers[idx].NodeID = uint32(workload.NodeID)
				}
			}
		}
	}

	return nil
}

// AddK8sWorker adds a worker to a deployed kubernetes cluster
func (c *Client) AddK8sWorker(ctx context.Context, worker AddWorkerParams) error {
	clusterContracts, err := c.loadModelContracts(ctx, worker.ClusterName)
	if err != nil {
		return errors.Wrapf(err, "failed to get kubernetes cluster %s contracts", worker.ClusterName)
	}

	if len(clusterContracts.nodeContarcts) == 0 {
		return fmt.Errorf("found 0 contracts for cluster %s", worker.ClusterName)
	}

	znet, err := c.loadNetwork(worker.ClusterName, clusterContracts.network)
	if err != nil {
		return errors.Wrapf(err, "failed to load network for cluster %s", worker.ClusterName)
	}

	if !doesNetworkIncludeNode(znet.Nodes, worker.Worker.NodeID) {
		znet.Nodes = append(znet.Nodes, worker.Worker.NodeID)
		err = c.client.DeployNetwork(ctx, &znet)
		if err != nil {
			return errors.Wrap(err, "failed to deploy network")
		}
	}

	cluster, err := c.loadK8s(worker.ClusterName, clusterContracts.nodeContarcts)
	if err != nil {
		return errors.Wrap(err, "failed to load kubernetes cluster")
	}

	cluster.Workers = append(cluster.Workers, newK8sNodeFromModel(worker.Worker))
	if err := c.client.DeployK8sCluster(ctx, &cluster); err != nil {
		return errors.Wrap(err, "failed to update kubernetes cluster")
	}

	return nil
}

// RemoveK8sWorker removes a worker from a deployed kubernetes cluster
func (c *Client) RemoveK8sWorker(ctx context.Context, worker RemoveWorkerParams) error {
	log.Info().Msgf("removing worker %s", worker.WorkerName)

	k8sContracts, err := c.loadModelContracts(ctx, worker.ClusterName)
	if err != nil {
		return errors.Wrapf(err, "failed to get kubernetes cluster %s contracts", worker.ClusterName)
	}

	if len(k8sContracts.nodeContarcts) == 0 {
		return fmt.Errorf("found 0 contracts for cluster %s", worker.ClusterName)
	}

	znet, err := c.loadNetwork(worker.ClusterName, k8sContracts.network)
	if err != nil {
		return errors.Wrapf(err, "failed to load network for cluster %s", worker.ClusterName)
	}

	cluster, err := c.loadK8s(worker.MasterName, k8sContracts.nodeContarcts)
	if err != nil {
		return errors.Wrap(err, "failed to load kubernetes cluster")
	}

	workerIdx, err := getWorkerIndex(&cluster, worker.WorkerName)
	if err != nil {
		return err
	}

	workerNodeID := cluster.Workers[workerIdx].Node

	cluster.Workers = append(cluster.Workers[:workerIdx], cluster.Workers[workerIdx+1:]...)

	nodeIDs := []uint32{}
	for _, worker := range cluster.Workers {
		nodeIDs = append(nodeIDs, worker.Node)
	}
	nodeIDs = append(nodeIDs, cluster.Master.Node)

	if err := c.client.DeployK8sCluster(ctx, &cluster); err != nil {
		return err
	}

	// TODO: check if there is no other worker on workerNodeID before updating network
	for idx, nodeID := range znet.Nodes {
		if nodeID == workerNodeID {
			znet.Nodes = append(znet.Nodes[:idx], znet.Nodes[idx+1:]...)
			break
		}
	}

	if err := c.client.DeployNetwork(ctx, &znet); err != nil {
		return err
	}

	return nil
}

func getWorkerIndex(cluster *workloads.K8sCluster, workerName string) (int, error) {
	for idx, worker := range cluster.Workers {
		if worker.Name == workerName {
			return idx, nil
		}
	}

	return 0, fmt.Errorf("failed to find a worker with name %s", workerName)
}
