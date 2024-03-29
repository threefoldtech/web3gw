package tfgrid

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/state"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/workloads"
	"github.com/threefoldtech/zos/pkg/gridtypes"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"
	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// Machines model ensures that each node has one deployment that includes all workloads
type NetworkDeployment struct {
	Name        string               `json:"name"`        // this is the name of the deployment, should be unique
	Metadata    string               `json:"metadata"`    // metadata for the model
	Description string               `json:"description"` // description of the model
	Network     NetworkConfiguration `json:"network"`     // network configuration
	VMs         []VMConfiguration    `json:"vms"`         // vm configurations
}

type NetworkConfiguration struct {
	Name               string `json:"name"`                 // network name will be (projectname.network)
	AddWireguardAccess bool   `json:"add_wireguard_access"` // true to add access node
	IPRange            string `json:"ip_range"`             // network ip range, must have a subnet mask of 16

	// computed
	WireguardConfig string `json:"wireguard_config"` // wireguard configuration, if any
}

type AddVMToNetworkDeployment struct {
	VMConfiguration
	Network string `json:"network"`
}

type RemoveVMFromNetworkDeployment struct {
	VM      string `json:"vm"`
	Network string `json:"network"`
}

type VMConfiguration struct {
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

type gridMachinesModel struct {
	modelName   string
	network     *workloads.ZNet
	deployments map[uint32]*workloads.Deployment
}

func (c *Client) DeployNetwork(ctx context.Context, args NetworkDeployment) (NetworkDeployment, error) {
	// validation
	if err := c.validateProjectName(ctx, args.Name); err != nil {
		return NetworkDeployment{}, err
	}

	if err := c.assignNodesIDsForMachines(ctx, args); err != nil {
		return NetworkDeployment{}, errors.Wrapf(err, "Couldn't find node for all machines model")
	}

	gridMachinesModel, err := toGridMachinesModel(&args)
	if err != nil {
		return NetworkDeployment{}, err
	}

	if err := c.deployMachinesModel(ctx, &gridMachinesModel); err != nil {
		return NetworkDeployment{}, err
	}

	return c.GetNetworkDeployment(ctx, args.Name)
}

func (r *Client) CancelNetworkDeployment(ctx context.Context, modelName string) error {
	if err := r.cancelModel(ctx, modelName); err != nil {
		return errors.Wrapf(err, "failed to cancel model %s contracts", modelName)
	}

	return nil
}

func (c *Client) GetNetworkDeployment(ctx context.Context, modelName string) (NetworkDeployment, error) {
	gridMachinesModel, err := c.loadGridMachinesModel(ctx, modelName)
	if err != nil {
		return NetworkDeployment{}, errors.Wrapf(err, "failed to load machines model %s deployments", modelName)
	}

	return c.toMachinesModel(&gridMachinesModel)
}

func (c *Client) AddVMToNetworkDeployment(ctx context.Context, params AddVMToNetworkDeployment) (NetworkDeployment, error) {
	log.Debug().Msgf("adding machine %s", params.Name)

	gridMachinesModel, err := c.loadGridMachinesModel(ctx, params.Network)
	if err != nil {
		return NetworkDeployment{}, err
	}

	if params.NodeID == 0 {
		if err := c.assignNodesIDForMachine(ctx, &params.VMConfiguration); err != nil {
			return NetworkDeployment{}, errors.Wrapf(err, "Couldn't find node for all machines model")
		}
	}

	if err := c.addMachine(ctx, &gridMachinesModel, &params); err != nil {
		return NetworkDeployment{}, errors.Wrapf(err, "failed to add machine %s", params.Name)
	}

	return c.GetNetworkDeployment(ctx, params.Network)
}

func (c *Client) RemoveVMFromNetworkDeployment(ctx context.Context, params RemoveVMFromNetworkDeployment) (NetworkDeployment, error) {
	log.Debug().Msgf("removeing machine %s", params.VM)

	gridMachinesModel, err := c.loadGridMachinesModel(ctx, params.Network)
	if err != nil {
		return NetworkDeployment{}, err
	}

	if err := c.removeMachine(ctx, &gridMachinesModel, &params); err != nil {
		return NetworkDeployment{}, errors.Wrapf(err, "failed to remove machine from model %s", params.Network)
	}

	return c.GetNetworkDeployment(ctx, params.Network)
}

func (c *Client) deployMachinesModel(ctx context.Context, model *gridMachinesModel) error {
	if err := c.deployZnet(ctx, model.network); err != nil {
		return err
	}

	if err := c.deployMachinesDeployments(ctx, model); err != nil {
		return err
	}

	if err := c.updateLocalState(model); err != nil {
		return err
	}

	return nil
}

func (c *Client) updateLocalState(g *gridMachinesModel) error {
	nodeContracts := map[uint32]state.ContractIDs{}
	for nodeID, dl := range g.deployments {
		nodeContracts[nodeID] = append(nodeContracts[nodeID], dl.ContractID)
	}

	for nodeID, contractID := range g.network.NodeDeploymentID {
		nodeContracts[nodeID] = append(nodeContracts[nodeID], contractID)
	}

	projectName := projectNameFromName(g.modelName)

	c.Projects[projectName] = ProjectState{
		nodeContracts: nodeContracts,
	}

	return nil
}

func (c *Client) deployZnet(ctx context.Context, znet *workloads.ZNet) error {
	if znet.AddWGAccess {
		privateKey, err := wgtypes.GeneratePrivateKey()
		if err != nil {
			return errors.Wrap(err, "failed to generate wireguard private key")
		}
		znet.ExternalSK = privateKey
	}

	if err := c.GridClient.DeployNetwork(ctx, znet); err != nil {
		return errors.Wrap(err, "failed to deploy znetwork")
	}

	return nil
}

func (r *Client) deployNetwork(ctx context.Context, modelName string, nodes []uint32, IPRange string, WGAccess bool) (*workloads.ZNet, error) {
	projectName := projectNameFromName(modelName)

	nodeList := []uint32{}
	nodeSet := map[uint32]struct{}{}
	for _, node := range nodes {
		if _, ok := nodeSet[node]; !ok {
			nodeList = append(nodeList, node)
			nodeSet[node] = struct{}{}
		}
	}

	ipRange, err := gridtypes.ParseIPNet(IPRange)
	if err != nil {
		return nil, errors.Wrapf(err, "network ip range (%s) is invalid", IPRange)
	}

	znet := workloads.ZNet{
		Name:         generateNetworkName(modelName),
		Nodes:        nodeList,
		IPRange:      ipRange,
		AddWGAccess:  WGAccess,
		SolutionType: projectName,
	}

	if znet.AddWGAccess {
		privateKey, err := wgtypes.GeneratePrivateKey()
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate wireguard private key")
		}
		znet.ExternalSK = privateKey
	}

	err = r.GridClient.DeployNetwork(ctx, &znet)
	if err != nil {
		return nil, errors.Wrap(err, "failed to deploy Z network")
	}

	return &znet, nil
}

func (c *Client) removeNodeFromNetwork(ctx context.Context, znet *workloads.ZNet, nodeID uint32) error {
	for idx, node := range znet.Nodes {
		if node == nodeID {
			znet.Nodes = append(znet.Nodes[:idx], znet.Nodes[idx+1:]...)
			return c.GridClient.DeployNetwork(ctx, znet)
		}
	}

	return nil
}

func generateNetworkName(modelName string) string {
	return fmt.Sprintf("%s_network", modelName)
}

func (c *Client) deployMachinesDeployments(ctx context.Context, g *gridMachinesModel) error {
	errGroup := errgroup.Group{}
	nodeDeploymentIDs := map[uint32]uint64{}
	for _, dl := range g.deployments {
		deployment := dl
		errGroup.Go(func() error {
			if err := c.GridClient.DeployDeployment(ctx, deployment); err != nil {
				return err
			}

			nodeDeploymentIDs[deployment.NodeID] = deployment.NodeDeploymentID[deployment.NodeID]
			return nil
		})
	}

	if err := errGroup.Wait(); err != nil {
		return err
	}

	for nodeID, dl := range g.deployments {
		dl.ContractID = nodeDeploymentIDs[nodeID]
		dl.NodeDeploymentID = map[uint32]uint64{nodeID: nodeDeploymentIDs[nodeID]}
	}

	return nil
}

func (c *Client) toMachinesModel(g *gridMachinesModel) (NetworkDeployment, error) {
	model := NetworkDeployment{
		Name:    g.modelName,
		Network: fromGridNetwork(g.network),
		VMs:     []VMConfiguration{},
	}

	for _, dl := range g.deployments {
		model.Name = dl.Name
		farmID, err := c.GridClient.GetNodeFarm(dl.NodeID)
		if err != nil {
			return NetworkDeployment{}, err
		}

		disks := getDiskMap(dl)
		qsfss := getQSFSMap(dl)

		for _, vm := range dl.Vms {
			grid_vm := toGridVM(dl.NodeID, &vm, disks, farmID, qsfss)
			model.VMs = append(model.VMs, grid_vm)
		}
	}

	return model, nil
}

func (machine *VMConfiguration) createReservationFromMachine() (string, *PlannedReservation) {
	neededSRU := 0
	neededHRU := 0
	for _, disk := range machine.Disks {
		neededSRU += disk.SizeGB * int(gridtypes.Gigabyte)
	}
	for _, qsfs := range machine.QSFSs {
		neededHRU += int(qsfs.Cache) * int(gridtypes.Gigabyte)
	}
	neededSRU += machine.RootfsSize * int(gridtypes.Megabyte)

	return machine.Name, &PlannedReservation{
		WorkloadName: machine.Name,
		MRU:          uint64(machine.Memory * int(gridtypes.Megabyte)),
		SRU:          uint64(neededSRU),
		HRU:          uint64(neededHRU),
		FarmID:       machine.FarmID,
		NodeID:       machine.NodeID,
	}
}

func (machine *VMConfiguration) assignNodeIdforMachine(reservations Reservations) {
	machine.NodeID = uint32(reservations[machine.Name].NodeID)
}

func (c *Client) assignNodesIDForMachine(ctx context.Context, machine *VMConfiguration) error {
	name, vm := machine.createReservationFromMachine()
	reservation := Reservations{name: vm}
	if err := c.AssignNodes(ctx, reservation); err != nil {
		return err
	}
	machine.assignNodeIdforMachine(reservation)
	return nil
}

// Assign chosen NodeIds to machines vm. with both way conversions to/from Reservations array.
func (c *Client) assignNodesIDsForMachines(ctx context.Context, args NetworkDeployment) error {
	// all units unified in bytes

	reservations := Reservations{}

	for idx := range args.VMs {
		name, vm := args.VMs[idx].createReservationFromMachine()
		reservations[name] = vm
	}

	err := c.AssignNodes(ctx, reservations)
	if err != nil {
		return err
	}

	for idx := range args.VMs {
		if args.VMs[idx].NodeID == 0 {
			args.VMs[idx].assignNodeIdforMachine(reservations)
		}
	}

	return nil
}

func (c *Client) addMachine(ctx context.Context, g *gridMachinesModel, params *AddVMToNetworkDeployment) error {
	if err := c.prepareModelForUpdate(g, params); err != nil {
		return err
	}

	if err := c.deployMachinesModel(ctx, g); err != nil {
		return err
	}

	return nil
}

func (c *Client) prepareModelForUpdate(g *gridMachinesModel, params *AddVMToNetworkDeployment) error {
	// update network
	if !slices.Contains(g.network.Nodes, params.NodeID) {
		g.network.Nodes = append(g.network.Nodes, params.NodeID)
	}

	// update deployment
	vm, disks, qsfss := extractMachineWorkloads(&params.VMConfiguration, g.network.Name)

	if dl, ok := g.deployments[params.NodeID]; ok && dl != nil {
		dl.Vms = append(dl.Vms, vm)
		dl.QSFS = append(dl.QSFS, qsfss...)
		dl.Disks = append(dl.Disks, disks...)
		return nil
	}

	newDl := workloads.NewDeployment(
		params.Name,
		params.NodeID,
		projectNameFromName(params.Name),
		nil,
		g.network.Name,
		disks,
		nil,
		[]workloads.VM{vm},
		qsfss)

	g.deployments[params.NodeID] = &newDl

	return nil
}

func (c *Client) removeMachine(ctx context.Context, g *gridMachinesModel, params *RemoveVMFromNetworkDeployment) error {
	model, err := c.toMachinesModel(g)
	if err != nil {
		return err
	}

	machine, err := model.findMachine(params.VM)
	if err != nil {
		return err
	}

	if err := c.removeMachineFromModel(ctx, g, machine); err != nil {
		return err
	}

	if _, ok := g.deployments[machine.NodeID]; !ok {
		if err := c.removeNodeFromNetwork(ctx, g.network, machine.NodeID); err != nil {
			return err
		}
	}

	if err := c.updateLocalState(g); err != nil {
		return err
	}

	return nil
}

func (c *Client) removeMachineFromModel(ctx context.Context, g *gridMachinesModel, machine *VMConfiguration) error {
	dl := g.deployments[machine.NodeID]

	if len(dl.Vms) == 1 {
		if err := c.GridClient.CancelDeployment(ctx, dl); err != nil {
			return err
		}

		delete(g.deployments, machine.NodeID)
		return nil
	}

	removeMachineFromDeployment(dl, machine)
	if err := c.GridClient.DeployDeployment(ctx, dl); err != nil {
		return err
	}

	return nil
}

func toGridMachinesModel(model *NetworkDeployment) (gridMachinesModel, error) {
	dls := toGridDeployments(model.VMs, model.Name)

	nodeIDs := []uint32{}
	for node := range dls {
		nodeIDs = append(nodeIDs, node)
	}

	znet, err := toGridZnet(&model.Network, nodeIDs, model.Name)
	if err != nil {
		return gridMachinesModel{}, err
	}

	return gridMachinesModel{
		modelName:   model.Name,
		network:     &znet,
		deployments: dls,
	}, nil
}

func toGridDeployments(machines []VMConfiguration, modelName string) map[uint32]*workloads.Deployment {
	dls := map[uint32]*workloads.Deployment{}

	nodeMachineMap := map[uint32][]*VMConfiguration{}
	for idx, machine := range machines {
		nodeMachineMap[machine.NodeID] = append(nodeMachineMap[machine.NodeID], &machines[idx])
	}

	networkName := generateNetworkName(modelName)

	for nodeID, machines := range nodeMachineMap {
		vms := []workloads.VM{}
		QSFSs := []workloads.QSFS{}
		disks := []workloads.Disk{}

		for _, machine := range machines {
			vm, machineDisks, machineQsfss := extractMachineWorkloads(machine, networkName)
			vms = append(vms, vm)
			QSFSs = append(QSFSs, machineQsfss...)
			disks = append(disks, machineDisks...)
		}

		clientDeployment := workloads.NewDeployment(modelName, nodeID, projectNameFromName(modelName), nil, networkName, disks, nil, vms, QSFSs)

		dls[nodeID] = &clientDeployment
	}

	return dls
}

func toGridZnet(network *NetworkConfiguration, nodeIDs []uint32, modelName string) (workloads.ZNet, error) {
	IPRange, err := gridtypes.ParseIPNet(network.IPRange)
	if err != nil {
		return workloads.ZNet{}, errors.Wrapf(err, "failed to parse network ip range %s", network.IPRange)
	}
	return workloads.ZNet{
		Name:         generateNetworkName(modelName),
		Nodes:        nodeIDs,
		IPRange:      IPRange,
		AddWGAccess:  network.AddWireguardAccess,
		SolutionType: projectNameFromName(modelName),
	}, nil
}

func fromGridNetwork(znet *workloads.ZNet) NetworkConfiguration {
	return NetworkConfiguration{
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

func generateGridDisk(disk *Disk, diskName string) workloads.Disk {
	return workloads.Disk{
		Name:        diskName,
		SizeGB:      disk.SizeGB,
		Description: disk.Description,
	}
}

func generateGridMount(diskName string, mountPoint string) workloads.Mount {
	return workloads.Mount{
		DiskName:   diskName,
		MountPoint: mountPoint,
	}
}

func generateGridQSFS(qsfs *QSFS, qsfsName string) workloads.QSFS {
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

	return workloads.QSFS{
		Name:                 qsfsName,
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
	}
}

func extractMachineWorkloads(machine *VMConfiguration, networkName string) (workloads.VM, []workloads.Disk, []workloads.QSFS) {
	disks := []workloads.Disk{}
	qsfss := []workloads.QSFS{}
	mounts := []workloads.Mount{}
	zlogs := []workloads.Zlog{}

	for idx, disk := range machine.Disks {
		diskName := generateDiskName(machine.Name, idx)
		disks = append(disks, generateGridDisk(&disk, diskName))
		mounts = append(mounts, generateGridMount(diskName, disk.MountPoint))
	}

	for idx, qsfs := range machine.QSFSs {
		qsfsName := generateQSFSName(machine.Name, idx)
		qsfss = append(qsfss, generateGridQSFS(&qsfs, qsfsName))
		mounts = append(mounts, generateGridMount(qsfsName, qsfs.MountPoint))
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

func toGridVM(nodeID uint32, vm *workloads.VM, diskMap map[string]workloads.Disk, farmID uint32, qsfsMap map[string]workloads.QSFS) VMConfiguration {
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
			disks = append(disks, fromGridDisk(disk, mount.MountPoint))
			continue
		}

		qsfs, ok := qsfsMap[mount.DiskName]
		if ok {
			qsfss = append(qsfss, fromGridQSFS(qsfs, mount.MountPoint))
		}
	}

	machine := VMConfiguration{
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

	return machine
}

func fromGridDisk(disk workloads.Disk, mountpoint string) Disk {
	return Disk{
		MountPoint:  mountpoint,
		SizeGB:      disk.SizeGB,
		Description: disk.Description,
		Name:        disk.Name,
	}
}

func fromGridQSFS(qsfs workloads.QSFS, mountpoint string) QSFS {
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
		Metadata:             fromGridMetadata(qsfs.Metadata),
		Groups:               fromGridGroups(qsfs.Groups),
		Name:                 qsfs.Name,
		MetricsEndpoint:      qsfs.MetricsEndpoint,
	}
}

func fromGridMetadata(metadata workloads.Metadata) Metadata {
	return Metadata{
		Type:                metadata.Type,
		Prefix:              metadata.Prefix,
		EncryptionAlgorithm: metadata.Prefix,
		EncryptionKey:       metadata.EncryptionKey,
		Backends:            fromGridBackends(metadata.Backends),
	}
}

func fromGridBackends(backends workloads.Backends) Backends {
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

func fromGridGroups(groups workloads.Groups) Groups {
	ret := Groups{}
	for _, g := range groups {
		ret = append(ret, Group{
			Backends: fromGridBackends(g.Backends),
		})
	}

	return ret
}

func generateDiskName(machineName string, id int) string {
	return fmt.Sprintf("%s_disk_%d", machineName, id)
}

func generateQSFSName(machineName string, id int) string {
	return fmt.Sprintf("%s_qsfs_%d", machineName, id)
}

func (n *NetworkDeployment) findMachine(machineName string) (*VMConfiguration, error) {
	for idx, machine := range n.VMs {
		if machine.Name == machineName {
			return &n.VMs[idx], nil
		}
	}

	return nil, fmt.Errorf("failed to find machine %s in model %s", machineName, n.Name)
}

func removeMachineFromDeployment(dl *workloads.Deployment, machine *VMConfiguration) {
	removeVMFromDeployment(dl, machine.Name)
	removeDisksFromDeployment(dl, machine.Disks)
	removeQSFSSFromDeployment(dl, machine.QSFSs)
}

func removeVMFromDeployment(dl *workloads.Deployment, machineName string) {
	for idx := range dl.Vms {
		if dl.Vms[idx].Name == machineName {
			dl.Vms = append(dl.Vms[:idx], dl.Vms[idx+1:]...)
			break
		}
	}
}

func removeDisksFromDeployment(dl *workloads.Deployment, disks []Disk) {
	for _, disk := range disks {
		for i := 0; i < len(dl.Disks); i++ {
			last := len(dl.Disks) - 1
			if dl.Disks[i].Name == disk.Name {
				dl.Disks[i], dl.Disks[last] = dl.Disks[last], dl.Disks[i]
				dl.Disks = dl.Disks[:last]
				i--
			}
		}
	}
}

func removeQSFSSFromDeployment(dl *workloads.Deployment, qsfss []QSFS) {
	for _, qsfs := range qsfss {
		for i := 0; i < len(dl.QSFS); i++ {
			last := len(dl.QSFS) - 1
			if dl.QSFS[i].Name == qsfs.Name {
				dl.QSFS[i], dl.QSFS[last] = dl.QSFS[last], dl.QSFS[i]
				dl.QSFS = dl.QSFS[:last]
				i--
			}
		}
	}
}
