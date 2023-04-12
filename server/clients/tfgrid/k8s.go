package tfgrid

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/threefoldtech/grid3-go/graphql"
	"github.com/threefoldtech/grid3-go/workloads"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"
)

// K8sCluster struct for k8s cluster
type K8sCluster struct {
	Name        string    `json:"name"`
	Master      *K8sNode  `json:"master"`
	Workers     []K8sNode `json:"workers"`
	Token       string    `json:"token"`
	NetworkName string    `json:"network_name"`
	SSHKey      string    `json:"ssh_key"`
}

// K8sNode kubernetes data
type K8sNode struct {
	Name      string `json:"name"`
	NodeID    uint32 `json:"node_id"`
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

func (r *Runner) K8sDeploy(ctx context.Context, cluster K8sCluster, projectName string) (K8sCluster, error) {
	// validate project name is unique
	if err := r.validateProjectName(ctx, projectName); err != nil {
		return K8sCluster{}, err
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
	k8s := newK8sClusterFromModel(cluster, projectName)

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

func (r *Runner) K8sDelete(ctx context.Context, projectName string) error {
	err := r.client.CancelProject(ctx, projectName)
	if err != nil {
		return errors.Wrapf(err, "failed to cancel project: %s", projectName)
	}

	return nil
}

func (r *Runner) K8sGet(ctx context.Context, clusterName string, projectName string) (K8sCluster, error) {
	// get all contracts by project name
	contracts, err := r.client.GetProjectContracts(ctx, projectName)
	if err != nil {
		return K8sCluster{}, errors.Wrapf(err, "failed to get contracts for project: %s", projectName)
	}

	if len(contracts.NodeContracts) == 0 {
		return K8sCluster{}, fmt.Errorf("found 0 contracts for project %s", projectName)
	}

	cluster, err := r.reconstructClusterFromContractIDs(ctx, clusterName, contracts)
	if err != nil {
		return K8sCluster{}, err
	}

	return cluster, nil
}

func NewClientK8sNodeFromK8sNode(k8sNode K8sNode) workloads.K8sNode {
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

func (r *Runner) reconstructClusterFromContractIDs(ctx context.Context, clusterName string, contracts graphql.Contracts) (K8sCluster, error) {
	result := K8sCluster{
		Name:        clusterName,
		Master:      &K8sNode{},
		Workers:     []K8sNode{},
		NetworkName: generateNetworkName(clusterName),
	}

	diskNameNodeNameMap := map[string]string{}
	nodeNameDiskSizeMap := map[string]int{}

	for _, contract := range contracts.NodeContracts {
		nodeClient, err := r.client.GetNodeClient(contract.NodeID)
		if err != nil {
			return K8sCluster{}, errors.Wrapf(err, "failed to get node %d client", contract.NodeID)
		}

		contractID, err := strconv.ParseUint(contract.ContractID, 10, 64)
		if err != nil {
			return K8sCluster{}, errors.Wrapf(err, "Couldn't convert ContractID: %s", contract.ContractID)
		}

		deployment, err := nodeClient.DeploymentGet(ctx, contractID)
		if err != nil {
			return K8sCluster{}, errors.Wrapf(err, "failed to get deployment with contract id %d", contractID)
		}

		for _, workload := range deployment.Workloads {
			if workload.Type == zos.ZMachineType {
				vm, err := workloads.NewVMFromWorkload(&workload, &deployment)
				if err != nil {
					return K8sCluster{}, errors.Wrapf(err, "Failed to get vm from workload: %s", workload.Name)
				}

				if len(vm.Mounts) == 1 {
					diskNameNodeNameMap[vm.Mounts[0].DiskName] = vm.Name
				}

				if isWorker(vm) {
					worker := NewK8sNodeFromVM(vm)
					worker.NodeID = contract.NodeID

					result.Workers = append(result.Workers, worker)
				} else {
					masterNode := NewK8sNodeFromVM(vm)
					masterNode.NodeID = contract.NodeID

					result.Master = &masterNode
					result.SSHKey = vm.EnvVars["SSH_KEY"]
					result.Token = vm.EnvVars["K3S_TOKEN"]
				}
			}
		}

		for _, workload := range deployment.Workloads {
			if workload.Type == zos.ZMountType {
				disk, err := workloads.NewDiskFromWorkload(&workload)
				if err != nil {
					return K8sCluster{}, errors.Wrapf(err, "Failed to get disk from workload: %s", workload.Name)
				}

				nodeName := diskNameNodeNameMap[disk.Name]
				nodeNameDiskSizeMap[nodeName] = disk.SizeGB
			}
		}
	}

	result.Master.DiskSize = nodeNameDiskSizeMap[result.Master.Name]
	for idx := range result.Workers {
		result.Workers[idx].DiskSize = nodeNameDiskSizeMap[result.Workers[idx].Name]
	}

	return result, nil
}

func NewK8sNodeFromVM(vm workloads.VM) K8sNode {
	return K8sNode{
		Name:      vm.Name,
		PublicIP:  vm.PublicIP,
		PublicIP6: vm.PublicIP6,
		Planetary: vm.Planetary,
		Flist:     vm.Flist,
		CPU:       vm.CPU,
		Memory:    vm.Memory,

		ComputedIP4: vm.ComputedIP,
		ComputedIP6: vm.ComputedIP6,
		WGIP:        vm.IP,
		YggIP:       vm.YggIP,
	}
}

func (k *K8sNode) assignComputedNodeValues(node workloads.K8sNode) {
	k.ComputedIP4 = node.ComputedIP
	k.ComputedIP6 = node.ComputedIP6
	k.WGIP = node.IP
	k.YggIP = node.YggIP
}

func isWorker(vm workloads.VM) bool {
	return len(vm.EnvVars["K3S_URL"]) != 0
}

func newK8sClusterFromModel(model K8sCluster, projectName string) workloads.K8sCluster {
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
		SolutionType: projectName,
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
