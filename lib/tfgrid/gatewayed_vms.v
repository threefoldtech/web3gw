module tfgrid

struct CapacityPackage {
	cpu u32
	memory u32
	size u32
}

enum Capacity {
	small
	medium
	large
}

type CapacityPackages = map[Capacity]CapacityPackage

gatewayed_vms_packages := CapacityPackages{
	small: CapacityPackage{
		cpu: 1,
		memory: 2048,
		size: 10
	},
	medium: CapacityPackage{
		cpu: 2,
		memory: 2048,
		size: 20
	},
	large: CapacityPackage{
		cpu: 4,
		memory: 8192,
		size: 50
	}
}

fn (client TFGridClient) gatewayed_vms_create(vms GatewayedVMsCreate) ! (tfgrid.MachinesResult, []tfgrid.GatewayNameResult){
	replicas := vms.number_of_machines

	net := tfgrid.Network{
		add_wireguard_access: gws.public_network
	}

	disk := tfgrid.Disk{
		size: vms.disk_size
		mountpoint: '/mnt/disk'
	}

	mut vm := tfgrid.Machine{
		name: rand.hex(8)
		farm_id: vm.farm_id
		public_ip: vm.public_ips

		disks: []tfgrid.Disk{}
		env_vars: map[string]string{}
	}

	vm.cpu = gatewayed_vms_packages[vms.capacity].cpu
	vm.memory = gatewayed_vms_packages[vms.capacity].memory
	vm.rootfs_size = gatewayed_vms_packages[vms.capacity].size

	vm.disks = []tfgrid.Disk{disk}
	vm.env_vars["SSH_KEY"] = vms.ssh_key

	mut machines := []tfgrid.Machine{}
	for i in 0..replicas {
		vm.name = rand.hex(8)
		machines << vm
	}

	vms := tfgrid.MachinesModel{
		name: 	  vms.model_name
		network:  net
		machines: machines
	}

	machines := client.machines_deploy(vms)!

	backends := []string{}
	for __vm in res.machines {
		backends << vm.ygg_ip
	}

	gateways := []tfgrid.GatewayNameResult{}

	for backend in backends {
		mut gw := tfgrid.GatewayName{
			name: 'gw' + rand.hex(8)

			// TODO: change this to pick a random gateway node
			node_id: 11
			tls_passthrough: false
			backends: ['http://${backends}:${vms.backend_port}']
		}
		gateways << client.gateways_deploy_name(gw)!
	}
	return machines, gateways
}