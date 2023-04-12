package tfgrid

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/threefoldtech/grid3-go/workloads"
	"github.com/threefoldtech/zos/pkg/gridtypes"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"
)

// Machines model ensures that each node has one deployment that includes all workloads
type MachinesModel struct {
	Name        string    `json:"name"`     // this is the model name, should be unique
	Network     Network   `json:"network"`  // network specs
	Machines    []Machine `json:"machines"` // machines specs
	Metadata    string    `json:"metadata"`
	Description string    `json:"description"`
}

type Network struct {
	AddWireguardAccess bool   `json:"add_wireguard_access"` // true to add access node
	IPRange            string `json:"ip_range"`

	// computed
	Name            string `json:"name"` // network name will be (projectname.network)
	WireguardConfig string `json:"wireguard_config"`
}

type Machine struct {
	NodeID      uint32            `json:"node_id"`
	FarmID      uint32            `json:"farm_id"`
	Name        string            `json:"name"`
	Flist       string            `json:"flist"`
	PublicIP    bool              `json:"public_ip"`
	PublicIP6   bool              `json:"public_ip6"`
	Planetary   bool              `json:"planetary"`
	Description string            `json:"description"`
	CPU         int               `json:"cpu"`
	Memory      int               `json:"memory"`
	RootfsSize  int               `json:"rootfs_size"`
	Entrypoint  string            `json:"entrypoint"`
	Zlogs       []Zlog            `json:"zlogs"`
	Disks       []Disk            `json:"disks"`
	QSFSs       []QSFS            `json:"qsfss"`
	EnvVars     map[string]string `json:"env_vars"`

	// computed
	ComputedIP4 string `json:"computed_ip4"`
	ComputedIP6 string `json:"computed_ip6"`
	WGIP        string `json:"wireguard_ip"`
	YggIP       string `json:"ygg_ip"`
}

// Zlog logger struct
type Zlog struct {
	Output string `json:"output"`
}

// Disk struct
type Disk struct {
	MountPoint  string `json:"mountpoint"`
	SizeGB      int    `json:"size"`
	Description string `json:"description"`

	// computed
	Name string `json:"name"`
}

// QSFS struct
type QSFS struct {
	MountPoint           string   `json:"mountpoint"`
	Description          string   `json:"description"`
	Cache                int      `json:"cache"`
	MinimalShards        uint32   `json:"minimal_shards"`
	ExpectedShards       uint32   `json:"expected_shards"`
	RedundantGroups      uint32   `json:"redundant_groups"`
	RedundantNodes       uint32   `json:"redundant_nodes"`
	MaxZDBDataDirSize    uint32   `json:"max_zdb_data_dir_size"`
	EncryptionAlgorithm  string   `json:"encryption_algorithm"`
	EncryptionKey        string   `json:"encryption_key"`
	CompressionAlgorithm string   `json:"compression_algorithm"`
	Metadata             Metadata `json:"metadata"`
	Groups               Groups   `json:"groups"`

	// computed
	Name            string `json:"name"`
	MetricsEndpoint string `json:"metrics_endpoint"`
}

// Metadata for QSFS
type Metadata struct {
	Type                string   `json:"type"`
	Prefix              string   `json:"prefix"`
	EncryptionAlgorithm string   `json:"encryption_algorithm"`
	EncryptionKey       string   `json:"encryption_key"`
	Backends            Backends `json:"backends"`
}

// Group is a zos group
type Group struct {
	Backends Backends `json:"backends"`
}

// Backend is a zos backend
type Backend zos.ZdbBackend

// Groups is a list of groups
type Groups []Group

// Backends is a list of backends
type Backends []Backend

// nodes should always be provided
func (r *Runner) MachinesDeploy(ctx context.Context, model MachinesModel, projectName string) (MachinesModel, error) {
	/*
		- validate incoming deployment
			- project name has to be unique
		- construct network deployer
			- get nodes from all machines
			- build network deployer using these nodes
		- deploy network
		- construct deployment deployer
		- deploy deployment
		- construct machines model and return it
	*/

	// validation
	if err := r.validateProjectName(ctx, projectName); err != nil {
		return MachinesModel{}, err
	}

	if err := r.assignNodesIDsForMachines(ctx, &model); err != nil {
		return MachinesModel{}, errors.Wrapf(err, "Couldn't find node for all machines model")
	}

	// deploy network
	nodes := []uint32{}
	for _, machine := range model.Machines {
		nodes = append(nodes, machine.NodeID)
	}

	znet, err := r.deployNetwork(ctx, model.Name, nodes, model.Network.IPRange, model.Network.AddWireguardAccess, projectName)
	if err != nil {
		return MachinesModel{}, err
	}

	// deploy deployment
	nodeDeploymentID, err := r.deployMachinesWorkloads(ctx, &model, projectName)
	if err != nil {
		// TODO: if error happens midway, all created contracts should be deleted
		return MachinesModel{}, err
	}

	net := Network{
		Name:               znet.Name,
		AddWireguardAccess: znet.AddWGAccess,
		IPRange:            znet.IPRange.String(),
		WireguardConfig:    znet.AccessWGConfig,
	}

	// construct result
	resModel, err := r.constructMachinesModelFromContracts(ctx, nodeDeploymentID, model.Name, net)
	if err != nil {
		return MachinesModel{}, err
	}

	return resModel, nil
}

func (m *MachinesModel) generateDiskNames() {
	for _, machine := range m.Machines {
		for idx := range machine.Disks {
			machine.Disks[idx].Name = generateDiskName(machine.Name, idx)
		}
	}
}

func (r *Runner) deployMachinesWorkloads(ctx context.Context, model *MachinesModel, projectName string) (map[uint32]uint64, error) {
	model.generateDiskNames()

	nodeMachineMap := map[uint32][]*Machine{}
	for idx, machine := range model.Machines {
		nodeMachineMap[machine.NodeID] = append(nodeMachineMap[machine.NodeID], &model.Machines[idx])
	}

	nodeDeploymentID := map[uint32]uint64{}

	networkName := generateNetworkName(model.Name)

	for nodeID, machines := range nodeMachineMap {
		vms := []workloads.VM{}
		QSFSs := []workloads.QSFS{}
		disks := []workloads.Disk{}

		for _, machine := range machines {
			nodeVM, nodeDisks, nodeQSFSs := r.extractWorkloads(machine, networkName)
			vms = append(vms, nodeVM)
			QSFSs = append(QSFSs, nodeQSFSs...)
			disks = append(disks, nodeDisks...)
		}

		clientDeployment := workloads.NewDeployment(model.Name, nodeID, projectName, nil, networkName, disks, nil, vms, QSFSs)
		contractID, err := r.client.DeployDeployment(ctx, &clientDeployment)
		if err != nil {
			return nil, errors.Wrap(err, "failed to deploy")
		}

		nodeDeploymentID[nodeID] = contractID
	}

	return nodeDeploymentID, nil
}

func (r *Runner) extractWorkloads(machine *Machine, networkName string) (workloads.VM, []workloads.Disk, []workloads.QSFS) {
	disks := []workloads.Disk{}
	qsfss := []workloads.QSFS{}
	mounts := []workloads.Mount{}
	zlogs := []workloads.Zlog{}

	for idx, disk := range machine.Disks {
		diskName := generateDiskName(machine.Name, idx)
		disks = append(disks, workloads.Disk{
			Name:        diskName,
			SizeGB:      disk.SizeGB,
			Description: disk.Description,
		})
		mounts = append(mounts, workloads.Mount{
			DiskName:   diskName,
			MountPoint: disk.MountPoint,
		})
	}

	for idx, qsfs := range machine.QSFSs {
		metaBackends := []workloads.Backend{}
		for _, b := range qsfs.Metadata.Backends {
			metaBackends = append(metaBackends, workloads.Backend{
				Address:   b.Address,
				Namespace: b.Namespace,
				Password:  b.Password,
			})
		}
		groups := []workloads.Group{}
		for _, group := range qsfs.Groups {
			bs := workloads.Backends{}
			for _, b := range group.Backends {
				bs = append(bs, workloads.Backend{
					Address:   b.Address,
					Namespace: b.Namespace,
					Password:  b.Password,
				})
			}
			groups = append(groups, workloads.Group{Backends: bs})
		}

		qsfss = append(qsfss, workloads.QSFS{
			Name:                 generateQSFSName(machine.Name, idx),
			Description:          qsfs.Description,
			Cache:                qsfs.Cache,
			MinimalShards:        qsfs.MinimalShards,
			ExpectedShards:       qsfs.ExpectedShards,
			RedundantGroups:      qsfs.RedundantGroups,
			RedundantNodes:       qsfs.RedundantNodes,
			MaxZDBDataDirSize:    qsfs.MaxZDBDataDirSize,
			EncryptionAlgorithm:  qsfs.EncryptionAlgorithm,
			EncryptionKey:        qsfs.EncryptionKey,
			CompressionAlgorithm: qsfs.CompressionAlgorithm,
			Metadata: workloads.Metadata{
				Type:                qsfs.Metadata.Type,
				Prefix:              qsfs.Metadata.Prefix,
				EncryptionAlgorithm: qsfs.Metadata.EncryptionAlgorithm,
				EncryptionKey:       qsfs.Metadata.EncryptionKey,
				Backends:            metaBackends,
			},
			Groups:          groups,
			MetricsEndpoint: qsfs.MetricsEndpoint,
		})
	}

	for _, zlog := range machine.Zlogs {
		zlogs = append(zlogs, workloads.Zlog{
			Zmachine: machine.Name,
			Output:   zlog.Output,
		})
	}

	vm := workloads.VM{
		Name:        machine.Name,
		Flist:       machine.Flist,
		PublicIP:    machine.PublicIP,
		PublicIP6:   machine.PublicIP6,
		Planetary:   machine.Planetary,
		Description: machine.Description,
		CPU:         machine.CPU,
		Memory:      machine.Memory,
		RootfsSize:  machine.RootfsSize,
		Entrypoint:  machine.Entrypoint,
		Mounts:      mounts,
		Zlogs:       zlogs,
		EnvVars:     machine.EnvVars,
		NetworkName: networkName,
	}

	return vm, disks, qsfss
}

func (r *Runner) MachinesDelete(ctx context.Context, projectName string) error {
	if err := r.client.CancelProject(ctx, projectName); err != nil {
		return errors.Wrapf(err, "failed to cancel contracts")
	}

	return nil
}

func (r *Runner) MachinesGet(ctx context.Context, modelName string, projectName string) (MachinesModel, error) {
	contracts, err := r.client.GetProjectContracts(ctx, projectName)
	if err != nil {
		return MachinesModel{}, errors.Wrapf(err, "failed to retreive contracts with project name %s", projectName)
	}

	if len(contracts.NodeContracts) == 0 {
		return MachinesModel{}, fmt.Errorf("found 0 contracts for project %s", projectName)
	}

	nodeDeploymentID := map[uint32]uint64{}
	for _, c := range contracts.NodeContracts {
		contractID, err := strconv.Atoi(c.ContractID)
		if err != nil {
			return MachinesModel{}, errors.Wrapf(err, "failed to parse contract with id (%s)", c.ContractID)
		}
		nodeDeploymentID[c.NodeID] = uint64(contractID)
	}
	net := Network{
		Name: generateNetworkName(modelName),
	}

	model, err := r.constructMachinesModelFromContracts(ctx, nodeDeploymentID, modelName, net)
	if err != nil {
		return MachinesModel{}, errors.Wrapf(err, "failed to construct model for project")
	}

	return model, nil
}

func (r *Runner) constructMachinesModelFromContracts(ctx context.Context, nodeDeploymentID map[uint32]uint64, modelName string, net Network) (MachinesModel, error) {
	model := MachinesModel{
		Name:    modelName,
		Network: net,
	}
	for nodeID, contractID := range nodeDeploymentID {

		nodeClient, err := r.client.GetNodeClient(nodeID)
		if err != nil {
			return MachinesModel{}, errors.Wrapf(err, "failed to get node %d client", nodeID)
		}

		dl, err := nodeClient.DeploymentGet(ctx, contractID)
		if err != nil {
			return MachinesModel{}, errors.Wrapf(err, "failed to get deployment with contract id %d", contractID)
		}

		machineMap := map[string]*Machine{}
		machineMountPoints := map[string]string{}
		// first get machines and znet
		for idx := range dl.Workloads {
			if dl.Workloads[idx].Type == zos.ZMachineType {
				vm, err := workloads.NewVMFromWorkload(&dl.Workloads[idx], &dl)
				if err != nil {
					return MachinesModel{}, errors.Wrapf(err, "failed to parse vm %s data", dl.Workloads[idx].Name)
				}

				machine := machineFromVM(&vm)
				machine.NodeID = nodeID
				machineMap[machine.Name] = &machine

				for _, mp := range vm.Mounts {
					machineMountPoints[mp.DiskName] = mp.MountPoint
				}
			}

			if dl.Workloads[idx].Type == zos.NetworkType && model.Network.IPRange == "" {
				net, err := workloads.NewNetworkFromWorkload(dl.Workloads[idx], nodeID)
				if err != nil {
					return MachinesModel{}, errors.Wrapf(err, "failed to parse network %s data", dl.Workloads[idx].Name)
				}

				model.Network.IPRange = net.IPRange.String()
			}
		}

		// get disks and qsfss
		for idx := range dl.Workloads {
			if dl.Workloads[idx].Type == zos.ZMountType {
				disk, err := workloads.NewDiskFromWorkload(&dl.Workloads[idx])
				if err != nil {
					return MachinesModel{}, errors.Wrapf(err, "failed to parse disk %s data", dl.Workloads[idx].Name)
				}

				machineName, err := getMachineNameFromMount(disk.Name)
				if err != nil {
					return MachinesModel{}, errors.Wrapf(err, "failed to extract machine name from disk with name %s", disk.Name)
				}

				machine, ok := machineMap[machineName]
				if !ok {
					return MachinesModel{}, errors.Wrapf(err, "disk (%s) is not mounted on any machine", disk.Name)
				}

				machine.Disks = append(machine.Disks, Disk{
					Name:        disk.Name,
					SizeGB:      disk.SizeGB,
					Description: disk.Description,
					MountPoint:  machineMountPoints[disk.Name],
				})
			} else if dl.Workloads[idx].Type == zos.QuantumSafeFSType {
				qsfs, err := workloads.NewQSFSFromWorkload(&dl.Workloads[idx])
				if err != nil {
					return MachinesModel{}, errors.Wrapf(err, "failed to parse qsfs %s data", qsfs.Name)
				}

				machineName, err := getMachineNameFromMount(qsfs.Name)
				if err != nil {
					return MachinesModel{}, errors.Wrapf(err, "failed to extract machine name from qsfs with name %s", qsfs.Name)
				}

				machine, ok := machineMap[machineName]
				if !ok {
					return MachinesModel{}, errors.Wrapf(err, "qsfs (%s) is not mounted on any machine", qsfs.Name)
				}

				metaBackends := []Backend{}
				for _, b := range qsfs.Metadata.Backends {
					metaBackends = append(metaBackends, Backend{
						Address:   b.Address,
						Namespace: b.Namespace,
						Password:  b.Password,
					})
				}

				groups := []Group{}
				for _, group := range qsfs.Groups {
					bs := Backends{}
					for _, b := range group.Backends {
						bs = append(bs, Backend{
							Address:   b.Address,
							Namespace: b.Namespace,
							Password:  b.Password,
						})
					}
					groups = append(groups, Group{Backends: bs})
				}

				machine.QSFSs = append(machine.QSFSs, QSFS{
					MountPoint:           machineMountPoints[machineName],
					Description:          qsfs.Description,
					Cache:                qsfs.Cache,
					MinimalShards:        qsfs.MinimalShards,
					ExpectedShards:       qsfs.ExpectedShards,
					RedundantGroups:      qsfs.RedundantGroups,
					RedundantNodes:       qsfs.RedundantNodes,
					MaxZDBDataDirSize:    qsfs.MaxZDBDataDirSize,
					EncryptionAlgorithm:  qsfs.EncryptionAlgorithm,
					EncryptionKey:        qsfs.EncryptionKey,
					CompressionAlgorithm: qsfs.CompressionAlgorithm,
					Metadata: Metadata{
						Type:                qsfs.Metadata.Type,
						Prefix:              qsfs.Metadata.Prefix,
						EncryptionAlgorithm: qsfs.Metadata.EncryptionAlgorithm,
						EncryptionKey:       qsfs.Metadata.EncryptionKey,
						Backends:            metaBackends,
					},
					Groups:          groups,
					Name:            qsfs.Name,
					MetricsEndpoint: qsfs.MetricsEndpoint,
				})

			}
		}

		machines := []Machine{}
		for _, m := range machineMap {
			machines = append(machines, *m)
		}

		model.Machines = append(model.Machines, machines...)
	}

	return model, nil
}

func getMachineNameFromMount(name string) (string, error) {
	// disk or qsfs name should be in the form: vmname_disk/qsfs_X
	s := strings.Split(name, "_")
	if len(s) == 0 {
		return "", fmt.Errorf("workload name is invalid")
	}
	return s[0], nil
}

func machineFromVM(vm *workloads.VM) Machine {
	zlogs := []Zlog{}
	for _, zlog := range vm.Zlogs {
		zlogs = append(zlogs, Zlog{
			Output: zlog.Output,
		})
	}
	machine := Machine{
		NodeID:      0,
		Name:        vm.Name,
		Flist:       vm.Flist,
		PublicIP:    vm.PublicIP,
		PublicIP6:   vm.PublicIP6,
		Planetary:   vm.Planetary,
		Description: vm.Description,
		CPU:         vm.CPU,
		Memory:      vm.Memory,
		RootfsSize:  vm.RootfsSize,
		Entrypoint:  vm.Entrypoint,
		EnvVars:     vm.EnvVars,
		ComputedIP4: vm.ComputedIP,
		ComputedIP6: vm.ComputedIP6,
		WGIP:        vm.IP,
		YggIP:       vm.YggIP,
		Zlogs:       zlogs,
	}
	return machine
}

func generateNetworkName(modelName string) string {
	return fmt.Sprintf("%s_network", modelName)
}

func generateDiskName(machineName string, id int) string {
	return fmt.Sprintf("%s_disk_%d", machineName, id)
}

func generateQSFSName(machineName string, id int) string {
	return fmt.Sprintf("%s_qsfs_%d", machineName, id)
}

// Assign chosen NodeIds to machines vm. with both way conversions to/from Reservations array.
func (r *Runner) assignNodesIDsForMachines(ctx context.Context, machines *MachinesModel) error {
	// all units unified in bytes

	workloads := []*PlannedReservation{}

	for idx := range machines.Machines {
		neededSRU := 0
		neededHRU := 0
		for _, disk := range machines.Machines[idx].Disks {
			neededSRU += disk.SizeGB * int(gridtypes.Gigabyte)
		}
		for _, qsfs := range machines.Machines[idx].QSFSs {
			neededHRU += int(qsfs.Cache) * int(gridtypes.Gigabyte)
		}
		neededSRU += machines.Machines[idx].RootfsSize * int(gridtypes.Megabyte)

		workloads = append(workloads, &PlannedReservation{
			WorkloadName: machines.Machines[idx].Name,
			MRU:          uint64(machines.Machines[idx].Memory * int(gridtypes.Megabyte)),
			SRU:          uint64(neededSRU),
			HRU:          uint64(neededHRU),
			FarmID:       machines.Machines[idx].FarmID,
		})
	}

	err := r.AssignNodes(ctx, workloads)
	if err != nil {
		return err
	}

	for idx := range machines.Machines {
		if machines.Machines[idx].NodeID == 0 {
			for _, workload := range workloads {
				if workload.WorkloadName == machines.Machines[idx].Name {
					machines.Machines[idx].NodeID = uint32(workload.NodeID)
				}
			}
		}
	}

	return nil
}
