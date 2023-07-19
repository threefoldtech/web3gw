package tfgrid

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/graphql"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/state"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/workloads"
	"github.com/threefoldtech/web3_proxy/server/clients/tfgrid/mocks"
	"github.com/threefoldtech/zos/pkg/gridtypes"
)

func TestMachines(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cl := mocks.NewMockTFGridClient(ctrl)

	r := Client{
		GridClient: cl,
		Projects:   make(map[string]ProjectState),
	}

	t.Run("machines_deploy_success", func(t *testing.T) {
		nodeID := uint32(1)
		modelName := "model1"
		projectName := generateProjectName(modelName)
		networkContratID := uint64(1)
		deploymentContractID := uint64(2)
		model := NetworkDeployment{
			Name: modelName,
			Network: NetworkConfiguration{
				IPRange:            "10.1.0.0/16",
				AddWireguardAccess: false,
			},
			VMs: []VMConfiguration{
				{
					NodeID:    nodeID,
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

		want := NetworkDeployment{
			Name: modelName,
			Network: NetworkConfiguration{
				AddWireguardAccess: false,
				IPRange:            "10.1.0.0/16",
				Name:               generateNetworkName(model.Name),
			},
			VMs: []VMConfiguration{
				{
					NodeID:    nodeID,
					Name:      "vm1",
					FarmID:    1,
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
			Nodes:        []uint32{nodeID},
			IPRange:      ipRange,
			SolutionType: projectName,
		}

		cl.EXPECT().DeployNetwork(context.Background(), &znet).DoAndReturn(func(ctx context.Context, znet *workloads.ZNet) error {
			znet.NodeDeploymentID = map[uint32]uint64{nodeID: networkContratID}
			return nil
		})

		model.Network.Name = generateNetworkName(model.Name)

		// TODO: deployment should not be any
		cl.EXPECT().DeployDeployment(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, clientDeployment *workloads.Deployment) error {
			clientDeployment.ContractID = deploymentContractID
			clientDeployment.NodeDeploymentID = map[uint32]uint64{nodeID: deploymentContractID}
			return nil
		})

		cl.EXPECT().GetNodeFarm(nodeID).Return(uint32(1), nil)

		cl.EXPECT().SetContractState(map[uint32]state.ContractIDs{nodeID: {deploymentContractID, networkContratID}})
		cl.EXPECT().LoadNetwork(generateNetworkName(modelName)).Return(workloads.ZNet{
			Name:             generateNetworkName(model.Name),
			Nodes:            []uint32{nodeID},
			IPRange:          ipRange,
			SolutionType:     projectName,
			NodeDeploymentID: map[uint32]uint64{nodeID: networkContratID},
		}, nil)

		cl.EXPECT().LoadDeployment(modelName, nodeID).Return(workloads.Deployment{
			Name:             modelName,
			NodeID:           nodeID,
			SolutionType:     projectName,
			SolutionProvider: nil,
			NetworkName:      generateNetworkName(modelName),
			Disks: []workloads.Disk{
				{
					SizeGB: 10,
					Name:   generateDiskName("vm1", 0),
				},
			},
			Vms: []workloads.VM{
				{
					Name:      "vm1",
					PublicIP:  true,
					Planetary: true,
					Zlogs: []workloads.Zlog{
						{
							Output: "hamada",
						},
					},

					EnvVars:    map[string]string{"hello": "world"},
					ComputedIP: "1.1.1.1/16",
					YggIP:      "4.4.4.4",
					IP:         "1.1.1.1",
					Mounts: []workloads.Mount{
						{
							DiskName:   generateDiskName("vm1", 0),
							MountPoint: "point1",
						},
					},
				},
			},
			ContractID: 2,
		}, nil)

		got, err := r.DeployNetwork(context.Background(), model)
		assert.NoError(t, err)

		assert.Equal(t, want, got)
	})

	t.Run("machines_get_success", func(t *testing.T) {
		nodeID := uint32(1)
		modelName := "model2"
		projectName := generateProjectName(modelName)
		networkName := generateNetworkName(modelName)
		networkContractID := uint64(1)
		nodeContractID := uint64(2)
		vmName := "vm1"

		want := NetworkDeployment{
			Name: modelName,
			Network: NetworkConfiguration{
				AddWireguardAccess: false,
				IPRange:            "10.1.0.0/16",
				Name:               generateNetworkName(modelName),
			},
			VMs: []VMConfiguration{
				{
					NodeID:    nodeID,
					FarmID:    1,
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
					EnvVars:     map[string]string{"hello": "world"},
				},
			},
		}

		networkDeploymentData, err := json.Marshal(workloads.DeploymentData{
			Type: "network",
			Name: networkName,
		})
		assert.NoError(t, err)

		nodesDeploymentData, err := json.Marshal(workloads.DeploymentData{
			Type: "vm",
			Name: modelName,
		})
		assert.NoError(t, err)

		cl.EXPECT().GetProjectContracts(gomock.Any(), projectName).Return(graphql.Contracts{
			NodeContracts: []graphql.Contract{
				{
					ContractID:     "1",
					NodeID:         nodeID,
					DeploymentData: string(networkDeploymentData),
				},
				{
					ContractID:     "2",
					NodeID:         nodeID,
					DeploymentData: string(nodesDeploymentData),
				},
			},
		}, nil)

		cl.EXPECT().SetContractState(map[uint32]state.ContractIDs{nodeID: {networkContractID, nodeContractID}})
		cl.EXPECT().LoadNetwork(networkName).Return(workloads.ZNet{
			Name:        networkName,
			Nodes:       []uint32{nodeID},
			AddWGAccess: false,
			IPRange:     gridtypes.MustParseIPNet("10.1.0.0/16"),
		}, nil)

		cl.EXPECT().GetNodeFarm(nodeID).Return(uint32(1), nil)

		cl.EXPECT().LoadDeployment(modelName, nodeID).Return(workloads.Deployment{
			Name:             modelName,
			NodeID:           nodeID,
			SolutionType:     projectName,
			SolutionProvider: nil,
			NetworkName:      networkName,
			Disks: []workloads.Disk{
				{
					SizeGB: 10,
					Name:   generateDiskName("vm1", 0),
				},
			},
			Vms: []workloads.VM{
				{
					Name:      "vm1",
					PublicIP:  true,
					Planetary: true,
					Zlogs: []workloads.Zlog{
						{
							Output: "hamada",
						},
					},

					EnvVars:    map[string]string{"hello": "world"},
					ComputedIP: "1.1.1.1/16",
					YggIP:      "4.4.4.4",
					IP:         "1.1.1.1",
					Mounts: []workloads.Mount{
						{
							DiskName:   generateDiskName("vm1", 0),
							MountPoint: "point1",
						},
					},
					CPU:        2,
					Memory:     10,
					Entrypoint: "entry point",
				},
			},
			ContractID: nodeContractID,
		}, nil)

		got, err := r.GetNetworkDeployment(context.Background(), modelName)
		assert.NoError(t, err)

		assert.Equal(t, want, got)
	})
}
