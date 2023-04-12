package tfgrid

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/threefoldtech/grid3-go/graphql"
	client "github.com/threefoldtech/grid3-go/node"
	"github.com/threefoldtech/grid3-go/workloads"
	"github.com/threefoldtech/web3_proxy/server/clients/tfgrid/mocks"
	"github.com/threefoldtech/zos/pkg/gridtypes"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"
)

func TestMachines(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cl := mocks.NewMockTFGridClient(ctrl)

	r := Runner{
		client: cl,
	}

	t.Run("machines_deploy_success", func(t *testing.T) {
		projectName := "project1"
		rmbClient := mocks.NewMockClient(ctrl)
		model := MachinesModel{
			Name: "model1",
			Network: Network{
				IPRange:            "10.1.0.0/16",
				AddWireguardAccess: false,
			},
			Machines: []Machine{
				{
					NodeID:    1,
					Name:      "vm1",
					PublicIP:  true,
					Planetary: true,
					Zlogs: []Zlog{
						{
							Output: "hamada",
						},
					},
					Disks: []Disk{
						{
							MountPoint: "point1",
							SizeGB:     10,
						},
					},
					EnvVars: map[string]string{"hello": "world"},
				},
			},
		}

		want := MachinesModel{
			Name: "model1",
			Network: Network{
				AddWireguardAccess: false,
				IPRange:            "10.1.0.0/16",
				Name:               generateNetworkName(model.Name),
			},
			Machines: []Machine{
				{
					NodeID:    1,
					Name:      "vm1",
					PublicIP:  true,
					Planetary: true,
					Zlogs: []Zlog{
						{
							Output: "hamada",
						},
					},
					Disks: []Disk{
						{
							MountPoint: "point1",
							SizeGB:     10,
							Name:       generateDiskName("vm1", 0),
						},
					},
					EnvVars:     map[string]string{"hello": "world"},
					ComputedIP4: "1.1.1.1/16",
					YggIP:       "4.4.4.4",
					WGIP:        "1.1.1.1",
				},
			},
		}

		cl.
			EXPECT().
			GetProjectContracts(gomock.Any(), projectName).
			Return(graphql.Contracts{}, nil)

		ipRange, err := gridtypes.ParseIPNet(model.Network.IPRange)
		assert.NoError(t, err)

		znet := workloads.ZNet{
			Name:         generateNetworkName(model.Name),
			Nodes:        []uint32{1},
			IPRange:      ipRange,
			SolutionType: projectName,
		}

		cl.EXPECT().DeployNetwork(gomock.Any(), &znet).Return(&znet, nil)

		model.Network.Name = generateNetworkName(model.Name)

		// TODO: deployment should not be any
		cl.EXPECT().DeployDeployment(gomock.Any(), gomock.Any()).Return(nil)

		nodeClient := client.NewNodeClient(1, rmbClient, 10)
		cl.EXPECT().GetNodeClient(uint32(1)).Return(nodeClient, nil)
		vmRes := zos.ZMachineResult{
			YggIP: "4.4.4.4",
		}

		vmResultBytes, err := json.Marshal(vmRes)
		assert.NoError(t, err)

		pubIPRes := zos.PublicIPResult{
			IP: gridtypes.MustParseIPNet("1.1.1.1/16"),
		}
		pubIPBytes, err := json.Marshal(pubIPRes)
		assert.NoError(t, err)
		depWorkloads := []gridtypes.Workload{
			{
				Version: 0,
				Name:    gridtypes.Name(model.Machines[0].Name),
				Type:    zos.ZMachineType,
				Data: gridtypes.MustMarshal(zos.ZMachine{

					Network: zos.MachineNetwork{
						PublicIP: gridtypes.Name(fmt.Sprintf("%sip", model.Machines[0].Name)),
						Interfaces: []zos.MachineInterface{
							{
								Network: gridtypes.Name(generateNetworkName(model.Name)),
								IP:      net.ParseIP("1.1.1.1"),
							},
						},
						Planetary: true,
					},
					ComputeCapacity: zos.MachineCapacity{
						CPU:    uint8(model.Machines[0].CPU),
						Memory: gridtypes.Unit(uint(model.Machines[0].Memory)) * gridtypes.Megabyte,
					},
					Entrypoint: model.Machines[0].Entrypoint,
					Mounts: []zos.MachineMount{
						{
							Name:       gridtypes.Name(generateDiskName(model.Machines[0].Name, 0)),
							Mountpoint: model.Machines[0].Disks[0].MountPoint,
						},
					},
					Env: model.Machines[0].EnvVars,
				}),
				Result: gridtypes.Result{
					Created: gridtypes.Now(),
					State:   gridtypes.StateOk,
					Data:    json.RawMessage(vmResultBytes),
				},
			},
			{
				Version: 0,
				Name:    gridtypes.Name(fmt.Sprintf("%sip", model.Machines[0].Name)),
				Type:    zos.PublicIPType,
				Data: gridtypes.MustMarshal(zos.PublicIP{
					V4: model.Machines[0].PublicIP,
				}),
				Result: gridtypes.Result{
					Created: gridtypes.Now(),
					State:   gridtypes.StateOk,
					Data:    json.RawMessage(pubIPBytes),
				},
			},
			{
				Name:        gridtypes.Name(generateDiskName(model.Machines[0].Name, 0)),
				Version:     0,
				Type:        zos.ZMountType,
				Description: "",
				Data: gridtypes.MustMarshal(zos.ZMount{
					Size: gridtypes.Unit(model.Machines[0].Disks[0].SizeGB) * gridtypes.Gigabyte,
				}),
			},
			{
				Version: 0,
				Name:    gridtypes.Name(model.Machines[0].Name),
				Type:    zos.ZLogsType,
				Data: gridtypes.MustMarshal(zos.ZLogs{
					ZMachine: gridtypes.Name(model.Machines[0].Name),
					Output:   "hamada",
				}),
				Result: gridtypes.Result{
					Created: gridtypes.Now(),
					State:   gridtypes.StateOk,
				},
			},
		}
		zosDeployment := workloads.NewGridDeployment(1, depWorkloads)
		dl := gridtypes.Deployment{}
		rmbClient.EXPECT().Call(gomock.Any(), gomock.Any(), "zos.deployment.get", gomock.Any(), &dl).
			DoAndReturn(func(ctx context.Context, twin uint32, fn string, data, result interface{}) error {
				var res *gridtypes.Deployment = result.(*gridtypes.Deployment)
				*res = zosDeployment
				return nil
			})

		got, err := r.MachinesDeploy(context.Background(), model, projectName)
		assert.NoError(t, err)

		assert.Equal(t, want, got)
	})

	t.Run("machines_get_success", func(t *testing.T) {
		projectName := "project1"
		modelName := "model1"
		networkName := generateNetworkName(modelName)
		vmName := "vm1"
		rmbClient := mocks.NewMockClient(ctrl)

		want := MachinesModel{
			Name: modelName,
			Network: Network{
				AddWireguardAccess: false,
				IPRange:            "10.1.0.0/16",
				Name:               generateNetworkName(modelName),
			},
			Machines: []Machine{
				{
					NodeID:    1,
					Name:      vmName,
					CPU:       2,
					Memory:    10,
					PublicIP:  true,
					Planetary: true,
					Zlogs: []Zlog{
						{
							Output: "hamada",
						},
					},
					Disks: []Disk{
						{
							MountPoint: "point1",
							SizeGB:     10,
							Name:       generateDiskName(vmName, 0),
						},
					},
					ComputedIP4: "1.1.1.1/16",
					YggIP:       "4.4.4.4",
					WGIP:        "1.1.1.1",
					Entrypoint:  "entry point",
				},
			},
		}

		cl.EXPECT().GetProjectContracts(gomock.Any(), projectName).Return(graphql.Contracts{
			NodeContracts: []graphql.Contract{
				{
					ContractID: "1",
					NodeID:     1,
				},
			},
		}, nil)

		nodeClient := client.NewNodeClient(1, rmbClient, 10)
		cl.EXPECT().GetNodeClient(uint32(1)).Return(nodeClient, nil)
		vmRes := zos.ZMachineResult{
			YggIP: "4.4.4.4",
		}

		vmResultBytes, err := json.Marshal(vmRes)
		assert.NoError(t, err)

		pubIPRes := zos.PublicIPResult{
			IP: gridtypes.MustParseIPNet("1.1.1.1/16"),
		}
		pubIPBytes, err := json.Marshal(pubIPRes)
		assert.NoError(t, err)
		depWorkloads := []gridtypes.Workload{
			{
				Version: 0,
				Name:    gridtypes.Name(vmName),
				Type:    zos.ZMachineType,
				Data: gridtypes.MustMarshal(zos.ZMachine{

					Network: zos.MachineNetwork{
						PublicIP: gridtypes.Name(fmt.Sprintf("%sip", vmName)),
						Interfaces: []zos.MachineInterface{
							{
								Network: gridtypes.Name(generateNetworkName(modelName)),
								IP:      net.ParseIP("1.1.1.1"),
							},
						},
						Planetary: true,
					},
					ComputeCapacity: zos.MachineCapacity{
						CPU:    uint8(2),
						Memory: gridtypes.Unit(uint(10)) * gridtypes.Megabyte,
					},
					Entrypoint: "entry point",
					Mounts: []zos.MachineMount{
						{
							Name:       gridtypes.Name(generateDiskName(vmName, 0)),
							Mountpoint: "point1",
						},
					},
				}),
				Result: gridtypes.Result{
					Created: gridtypes.Now(),
					State:   gridtypes.StateOk,
					Data:    json.RawMessage(vmResultBytes),
				},
			},
			{
				Version: 0,
				Name:    gridtypes.Name(fmt.Sprintf("%sip", vmName)),
				Type:    zos.PublicIPType,
				Data: gridtypes.MustMarshal(zos.PublicIP{
					V4: true,
				}),
				Result: gridtypes.Result{
					Created: gridtypes.Now(),
					State:   gridtypes.StateOk,
					Data:    json.RawMessage(pubIPBytes),
				},
			},
			{
				Name:        gridtypes.Name(generateDiskName(vmName, 0)),
				Version:     0,
				Type:        zos.ZMountType,
				Description: "",
				Data: gridtypes.MustMarshal(zos.ZMount{
					Size: gridtypes.Unit(10) * gridtypes.Gigabyte,
				}),
			},
			{
				Version: 0,
				Name:    gridtypes.Name(vmName),
				Type:    zos.ZLogsType,
				Data: gridtypes.MustMarshal(zos.ZLogs{
					ZMachine: gridtypes.Name(vmName),
					Output:   "hamada",
				}),
				Result: gridtypes.Result{
					Created: gridtypes.Now(),
					State:   gridtypes.StateOk,
				},
			},
			{
				Version: 0,
				Name:    gridtypes.Name(networkName),
				Type:    zos.NetworkType,
				Data: gridtypes.MustMarshal(zos.Network{
					NetworkIPRange: gridtypes.MustParseIPNet("10.1.0.0/16"),
				}),
			},
		}
		zosDeployment := workloads.NewGridDeployment(1, depWorkloads)
		assert.NoError(t, err)
		dl := gridtypes.Deployment{}
		rmbClient.EXPECT().Call(gomock.Any(), gomock.Any(), "zos.deployment.get", gomock.Any(), &dl).
			DoAndReturn(func(ctx context.Context, twin uint32, fn string, data, result interface{}) error {
				var res *gridtypes.Deployment = result.(*gridtypes.Deployment)
				*res = zosDeployment
				return nil
			})

		got, err := r.MachinesGet(context.Background(), modelName, projectName)
		assert.NoError(t, err)

		assert.Equal(t, want, got)
	})
}
