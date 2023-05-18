module main

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.explorer

import flag
import log
import os


pub struct Taiga {
pub:
	name            string
	machine_farm_id u64
	gateway_farm_id u64
	cpu             u32
	memory          u32 // in mega bytes
	rootfs_size     u32 // in mega bytes
	disk_size       u32 // in giga bytes
	ssh_key         string
	admin_username  string
	admin_password  string
	admin_email     string
}

pub struct TaigaResult {
pub:
	name           string
	machine_ygg_ip string
	gateway_name   string
}

const (
	default_server_address = 'ws://127.0.0.1:8080'
)

fn deploy_taiga(mut client tfgrid.TFGridClient, mut explorer_client explorer.ExplorerClient, taiga Taiga, mut logger log.Logger) !TaigaResult {
	// determine which node will be used as gateway
	// get domain name for this node
	// provide fqdn to taiga machine
	gateway_nodes := explorer_client.nodes(explorer.NodesRequestParams{
		filters: explorer.NodeFilter{
			status: 'up'
			dedicated: false
			farm_ids: [taiga.gateway_farm_id]
			domain: true
		}
		pagination: explorer.Limit{
			size: 1
		}
	})!

	if gateway_nodes.nodes.len == 0{
		logger.error('failed to find an eligible node for gateway')
		return TaigaResult{}
	}

	machine_nodes := explorer_client.nodes(explorer.NodesRequestParams{
		filters: explorer.NodeFilter{
			status: 'up'
			dedicated: false
			farm_ids: [taiga.machine_farm_id]
			free_mru: u64(taiga.memory) * 1024 * 1024
			free_sru: u64(taiga.rootfs_size) * 1024 * 1024 + u64(taiga.disk_size) * 1024 * 1024 * 1024
		}
		pagination: explorer.Limit{
			size: 1
		}
	})!

	if machine_nodes.nodes.len == 0{
		logger.error('failed to find an eligible node for taiga vm')
		return TaigaResult{}
	}

	machine_node_id := machine_nodes.nodes[0].node_id
	gateway_node_id := gateway_nodes.nodes[0].node_id
	domain := gateway_nodes.nodes[0].public_config.domain

	defer {
		client.machines_delete(generate_taiga_machine_name(taiga.name)) or {}
	}

	// deploy machines
	machine := client.machines_deploy(tfgrid.MachinesModel{
		name: generate_taiga_machine_name(taiga.name)
		network: tfgrid.Network{
			add_wireguard_access: true
		}
		machines: [
			tfgrid.Machine{
				node_id: u32(machine_node_id)
				name:        'taiga_vm'
				cpu:         taiga.cpu
				memory:      taiga.memory
				rootfs_size: taiga.rootfs_size
				flist:       'https://hub.grid.tf/tf-official-apps/grid3_taiga_docker-latest.flist'
				env_vars:    {
					'DOMAIN_NAME':         '${taiga.name}.${domain}'
					'ADMIN_USERNAME':      taiga.admin_username
					'ADMIN_PASSWORD':      taiga.admin_password
					'ADMIN_EMAIL':         taiga.admin_email
					'SSH_KEY':             taiga.ssh_key
					'DEFAULT_FROM_EMAIL':  ''
					'EMAIL_USE_TLS':       ''
					'EMAIL_USE_SSL':       ''
					'EMAIL_HOST':          ''
					'EMAIL_PORT':          ''
					'EMAIL_HOST_USER':     ''
					'EMAIL_HOST_PASSWORD': ''
				}
				disks:       [
					tfgrid.Disk{
						size: taiga.disk_size
						mountpoint: '/mnt/disk1'
					},
				]
				planetary:   true
			},
		]
	}) or {
		logger.error('failed to deploy machine: ${err}')
		return TaigaResult{}
	}
	

	defer{
		client.gateways_delete_name(taiga.name) or {}
	}
	// deploy gateway
	gateway := client.gateways_deploy_name(tfgrid.GatewayName{
		name: taiga.name
		backends: ['http://${machine.machines[0].ygg_ip}:9000']
		node_id: u32(gateway_node_id),
	})!

	return TaigaResult{
		name: taiga.name
		machine_ygg_ip: 'machine.machines[0].ygg_ip'
		gateway_name: 'gateway.fqdn'
	}
}

fn generate_taiga_machine_name(taiga_name string) string {
	return '${taiga_name}_machine'
}


fn main(){
	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to the web3_proxy client. The web3_proxy client allows you to execute all remote procedure calls that the web3_proxy server can handle.')
	fp.limit_free_args(0, 0)!
	fp.description('')
	fp.skip_executable()
	mnemonic := fp.string('mnemonic', `m`, '', 'The mnemonic to be used to call any function')
	address := fp.string('address', `a`, '${default_server_address}', 'The address of the web3_proxy server to connect to.')
	debug_log := fp.bool('debug', 0, false, 'By setting this flag the client will print debug logs too.')
	_ := fp.finalize() or {
		eprintln(err)
		println(fp.usage())
		exit(1)
	}

	mut logger := log.Logger(&log.Log{
		level: if debug_log { .debug } else { .info }
	})

	mut myclient := rpcwebsocket.new_rpcwsclient(address, &logger) or {
		logger.error('Failed creating rpc websocket client: ${err}')
		exit(1)
	}

	_ := spawn myclient.run()

	mut tfgrid_client := tfgrid.new(mut myclient)

	mut explorer_client := explorer.new(mut myclient)
	explorer_client.load('dev') or{
		logger.error('failed to load exploere client instance: ${err}')
		exit(1)
	}
	

	tfgrid_client.load(tfgrid.Credentials{
		mnemonic: mnemonic // FILL IN YOUR MNEMONIC HERE
		network: 'dev'
	})!

	defer {
		tfgrid_client.logout() or {}
	}
	
	taiga := Taiga {
		name: 'hamadatiger'
		machine_farm_id: 1
		gateway_farm_id: 1
		cpu: 4
		memory: 8096
		rootfs_size: 51200
		disk_size: 100
		ssh_key: 'ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCQi7Qp0fs4WowSBQJonYHNWNJ5q776XbFO8HnUggyGse1Z4CFZyVpUdWaIzpFkQdivAACSKmqfE6jAunX7HuujTQhLhVgs/iCQ3ALQfQ118Jh1g2dz7S3/klHJs6JqfXLKtwDHzq2DuEDjls5PPoD6SVipjQo+kFO2tvKUYOrXryGW5VNPSUKtXZJX4kxtLzCANqENMSqZIBiJhXj7+JQq8Kp6E117dkLxh4BmPJmGS4stSAfgSdmEWgm0MgAbHkc2X+fLsWrcEBYaXE1b+n70bVXGDVEfeuMNZjBlsgULVR0DXY5zxegciOSNr1b7yF/ZdoALN0gmQg+AywPy92+oeY7EXLabDoDUKcE+42EHscXEkTHlhCieF/W9worCzGqpMwJuBDNvDu5kP1y/vB+ZfPVTlZ1kS77/OuDTr/zssQI/SgSszVXTyVSFIFIbXLGuUDscnmPHVPV4PnmeOa2aeF1cgX0o/JErQ8+iu2wqQKueZT4QAUFyknIgXloSBVs= mariocs@mario-codescalers'
		admin_username: 'admin_hamada'
		admin_password: 'admin_pass'
		admin_email: 'hamada@gridmail.com'
	}

	deploy_taiga(mut tfgrid_client, mut explorer_client, taiga, mut logger) or{
		logger.error('failed to deploy taiga isntance: ${err}')
		return
	}
}

