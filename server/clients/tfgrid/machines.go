package tfgrid

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/workloads"
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

type AddMachineParams struct {
	ModelName string  `json:"model_name"`
	Machine   Machine `json:"machine"`
}

type RemoveMachineParams struct {
	ModelName   string `json:"model_name"`
	MachineName string `json:"machine_name"`
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
func (r *Client) MachinesDeploy(ctx context.Context, model MachinesModel) (MachinesModel, error) {
	projectName := generateProjectName(model.Name)

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

	ret, err := r.loadModel(model.Name, nodeDeploymentID, znet.NodeDeploymentID)
	if err != nil {
		return MachinesModel{}, errors.Wrapf(err, "failed to rebuild machines model")
	}

	return ret, nil
}

func (c *Client) loadModel(modelName string, nodeDeployments map[uint32]uint64, networkDeployments map[uint32]uint64) (MachinesModel, error) {
	znet, err := c.loadNetwork(modelName, networkDeployments)
	if err != nil {
		return MachinesModel{}, err
	}

	deployments := []workloads.Deployment{}
	for nodeID, contractID := range nodeDeployments {
		dl, err := c.loadDeployment(modelName, nodeID, contractID)
		if err != nil {
			return MachinesModel{}, errors.Wrap(err, "failed to load deployments")
		}
		deployments = append(deployments, dl)
	}

	ret, err := c.buildModel(modelName, deployments, &znet)
	if err != nil {
		return MachinesModel{}, errors.Wrap(err, "failed to load machines model")
	}

	return ret, nil
}

func (c *Client) buildModel(name string, dls []workloads.Deployment, znet *workloads.ZNet) (MachinesModel, error) {
	model := MachinesModel{
		Name:     name,
		Network:  networkToModel(znet),
		Machines: []Machine{},
		// TODO: get description and metdata
		Description: "",
		Metadata:    "",
	}

	for _, dl := range dls {
		model.Name = dl.Name
		farmID, err := c.client.GetNodeFarm(dl.NodeID)
		if err != nil {
			return MachinesModel{}, err
		}

		disks := getDiskMap(&dl)
		qsfss := getQSFSMap(&dl)

		for _, vm := range dl.Vms {
			machine := machineFromVM(dl.NodeID, &vm, disks, farmID, qsfss)
			model.Machines = append(model.Machines, machine)
		}
	}

	return model, nil
}

func networkToModel(znet *workloads.ZNet) Network {
	return Network{
		AddWireguardAccess: znet.AddWGAccess,
		IPRange:            znet.IPRange.String(),
		Name:               znet.Name,
		WireguardConfig:    znet.AccessWGConfig,
	}
}

func getDiskMap(dl *workloads.Deployment) map[string]workloads.Disk {
	diskMap := map[string]workloads.Disk{}
	for _, disk := range dl.Disks {
		diskMap[disk.Name] = disk
	}

	return diskMap
}

func getQSFSMap(dl *workloads.Deployment) map[string]workloads.QSFS {
	qsfsMap := map[string]workloads.QSFS{}
	for _, qsfs := range dl.QSFS {
		qsfsMap[qsfs.Name] = qsfs
	}

	return qsfsMap
}

func (m *MachinesModel) generateDiskNames() {
	for _, machine := range m.Machines {
		for idx := range machine.Disks {
			machine.Disks[idx].Name = generateDiskName(machine.Name, idx)
		}
	}
}

func (r *Client) deployMachinesWorkloads(ctx context.Context, model *MachinesModel, projectName string) (map[uint32]uint64, error) {
	nodeDeploymentID := map[uint32]uint64{}

	model.generateDiskNames()

	nodeMachineMap := map[uint32][]*Machine{}
	for idx, machine := range model.Machines {
		nodeMachineMap[machine.NodeID] = append(nodeMachineMap[machine.NodeID], &model.Machines[idx])
	}

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

func (r *Client) extractWorkloads(machine *Machine, networkName string) (workloads.VM, []workloads.Disk, []workloads.QSFS) {
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

func (r *Client) MachinesDelete(ctx context.Context, modelName string) error {
	projectName := generateProjectName(modelName)

	if err := r.client.CancelProject(ctx, projectName); err != nil {
		return errors.Wrapf(err, "failed to cancel contracts")
	}

	return nil
}

func (c *Client) MachinesGet(ctx context.Context, modelName string) (MachinesModel, error) {
	modelContracts, err := c.loadModelContracts(ctx, modelName)
	if err != nil {
		return MachinesModel{}, errors.Wrapf(err, "failed to machines model %s contracts", modelName)
	}

	if len(modelContracts.nodeContarcts) == 0 {
		return MachinesModel{}, fmt.Errorf("found 0 contracts for model %s", modelName)
	}

	ret, err := c.loadModel(modelName, modelContracts.nodeContarcts, modelContracts.network)
	if err != nil {
		return MachinesModel{}, errors.Wrapf(err, "failed to rebuild machines model")
	}

	return ret, nil
}

func machineFromVM(nodeID uint32, vm *workloads.VM, diskMap map[string]workloads.Disk, farmID uint32, qsfsMap map[string]workloads.QSFS) Machine {
	zlogs := []Zlog{}
	for _, zlog := range vm.Zlogs {
		zlogs = append(zlogs, Zlog{
			Output: zlog.Output,
		})
	}

	var disks []Disk
	var qsfss []QSFS
	for _, mount := range vm.Mounts {
		disk, ok := diskMap[mount.DiskName]
		if ok {
			disks = append(disks, diskToModel(disk, mount.MountPoint))
			continue
		}

		qsfs, ok := qsfsMap[mount.DiskName]
		if ok {
			qsfss = append(qsfss, qsfsToModel(qsfs, mount.MountPoint))
		}
	}

	machine := Machine{
		NodeID:      nodeID,
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
		Disks:       disks,
		FarmID:      farmID,
		QSFSs:       qsfss,
	}
	// vm.Mounts[0].
	return machine
}

func diskToModel(disk workloads.Disk, mountpoint string) Disk {
	return Disk{
		MountPoint:  mountpoint,
		SizeGB:      disk.SizeGB,
		Description: disk.Description,
		Name:        disk.Name,
	}
}

func qsfsToModel(qsfs workloads.QSFS, mountpoint string) QSFS {
	return QSFS{
		MountPoint:           mountpoint,
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
		Metadata:             metadataToModel(qsfs.Metadata),
		Groups:               groupsToModel(qsfs.Groups),
		Name:                 qsfs.Name,
		MetricsEndpoint:      qsfs.MetricsEndpoint,
	}
}

func metadataToModel(metadata workloads.Metadata) Metadata {
	return Metadata{
		Type:                metadata.Type,
		Prefix:              metadata.Prefix,
		EncryptionAlgorithm: metadata.Prefix,
		EncryptionKey:       metadata.EncryptionKey,
		Backends:            backendsToModel(metadata.Backends),
	}
}

func backendsToModel(backends workloads.Backends) Backends {
	ret := Backends{}
	for _, b := range backends {
		ret = append(ret, Backend{
			Address:   b.Address,
			Namespace: b.Namespace,
			Password:  b.Password,
		})
	}

	return ret
}

func groupsToModel(groups workloads.Groups) Groups {
	ret := Groups{}
	for _, g := range groups {
		ret = append(ret, Group{
			Backends: backendsToModel(g.Backends),
		})
	}

	return ret
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
func (r *Client) assignNodesIDsForMachines(ctx context.Context, machines *MachinesModel) error {
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
			NodeID:       machines.Machines[idx].NodeID,
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

func (c *Client) MachineAdd(ctx context.Context, params AddMachineParams) (MachinesModel, error) {
	log.Info().Msgf("adding machine %s", params.Machine.Name)

	modelContracts, err := c.loadModelContracts(ctx, params.ModelName)
	if err != nil {
		return MachinesModel{}, errors.Wrapf(err, "failed to get machines model %s contracts", params.ModelName)
	}

	if len(modelContracts.nodeContarcts) == 0 {
		return MachinesModel{}, fmt.Errorf("found 0 contracts for model %s", params.ModelName)
	}

	if err := c.ensureNodeBelongsToNetwork(ctx, params.ModelName, modelContracts.network, params.Machine.NodeID); err != nil {
		return MachinesModel{}, err
	}

	if err := c.updateDeployment(ctx, modelContracts.nodeContarcts, &params); err != nil {
		return MachinesModel{}, err
	}

	ret, err := c.loadModel(params.ModelName, modelContracts.nodeContarcts, modelContracts.network)
	if err != nil {
		return MachinesModel{}, errors.Wrapf(err, "failed to load model %s", params.ModelName)
	}

	return ret, nil
}

func (c *Client) updateDeployment(ctx context.Context, oldDeployments map[uint32]uint64, params *AddMachineParams) error {
	dl, err := c.prepareDeploymentForUpdate(oldDeployments, params)
	if err != nil {
		return err
	}

	contractID, err := c.client.DeployDeployment(ctx, &dl)
	if err != nil {
		return errors.Wrap(err, "failed to deploy")
	}

	oldDeployments[dl.NodeID] = contractID

	return nil
}

func (c *Client) prepareDeploymentForUpdate(oldDeployments map[uint32]uint64, params *AddMachineParams) (workloads.Deployment, error) {
	networkName := generateNetworkName(params.ModelName)
	vm, disks, qsfss := c.extractWorkloads(&params.Machine, networkName)

	if contractID, ok := oldDeployments[params.Machine.NodeID]; ok {
		// there is an old deployment on this node; load and update this deployment.
		dl, err := c.loadDeployment(params.ModelName, params.Machine.NodeID, contractID)
		if err != nil {
			return workloads.Deployment{}, errors.Wrap(err, "failed to load deployments")
		}

		dl.Vms = append(dl.Vms, vm)
		dl.QSFS = append(dl.QSFS, qsfss...)
		dl.Disks = append(dl.Disks, disks...)

		return dl, nil
	}

	return workloads.NewDeployment(
		params.ModelName,
		params.Machine.NodeID,
		generateProjectName(params.ModelName),
		nil,
		networkName,
		disks,
		nil,
		[]workloads.VM{vm},
		qsfss), nil

}

func (c *Client) MachineRemove(ctx context.Context, params RemoveMachineParams) (MachinesModel, error) {
	log.Info().Msgf("removeing machine %s", params.MachineName)

	modelContracts, err := c.loadModelContracts(ctx, params.ModelName)
	if err != nil {
		return MachinesModel{}, errors.Wrapf(err, "failed to get model %s contracts", params.ModelName)
	}

	if len(modelContracts.nodeContarcts) == 0 {
		return MachinesModel{}, fmt.Errorf("found 0 contracts for model %s", params.ModelName)
	}

	model, err := c.loadModel(params.ModelName, modelContracts.nodeContarcts, modelContracts.network)
	if err != nil {
		return MachinesModel{}, errors.Wrapf(err, "failed to build model %s", params.ModelName)
	}

	for _, machine := range model.Machines {
		if machine.Name == params.MachineName {
			dl, err := getMachineDeployment(dls, params.MachineName)
			if err != nil {
				return MachinesModel{}, err
			}

			if len(dl.Vms) == 1 {
				if err := c.client.CancelDeployment(ctx, dl); err != nil {
					return MachinesModel{}, err
				}

				removeNodeFromZnet(&znet, dl.NodeID)
				if err := c.client.DeployNetwork(ctx, &znet); err != nil {
					return MachinesModel{}, err
				}

				return MachinesModel{}, nil
			}

			removeMachineFromDeployment(dl, &machine)
			_, err = c.client.DeployDeployment(ctx, dl)
			if err != nil {
				return MachinesModel{}, err
			}

			return MachinesModel{}, nil
		}
	}

	return MachinesModel{}, fmt.Errorf("failed to find machine with name %s", params.MachineName)
}

func removeMachineFromDeployment(dl *workloads.Deployment, machine *Machine) {
	for idx := range dl.Vms {
		if dl.Vms[idx].Name == machine.Name {
			dl.Vms = append(dl.Vms[:idx], dl.Vms[idx+1:]...)
			break
		}
	}

	for _, disk := range machine.Disks {
		for idx := range dl.Disks {
			if dl.Disks[idx].Name == disk.Name {
				dl.Disks = append(dl.Disks[:idx], dl.Disks[idx+1:]...)
				break
			}
		}
	}

	for _, qsfs := range machine.QSFSs {
		for idx := range dl.QSFS {
			if dl.QSFS[idx].Name == qsfs.Name {
				dl.QSFS = append(dl.QSFS[:idx], dl.QSFS[idx+1:]...)
				break
			}
		}
	}
}

func removeNodeFromZnet(znet *workloads.ZNet, nodeID uint32) {
	for idx, node := range znet.Nodes {
		if node == nodeID {
			znet.Nodes = append(znet.Nodes[:idx], znet.Nodes[idx+1:]...)
		}
	}
}

func getMachineDeployment(dls []workloads.Deployment, machineName string) (*workloads.Deployment, error) {
	for _, dl := range dls {
		for _, vm := range dl.Vms {
			if vm.Name == machineName {
				return &dl, nil
			}
		}
	}

	return nil, fmt.Errorf("failed to find deployment with machine %s", machineName)
}
