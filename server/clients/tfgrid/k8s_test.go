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

func TestK8s(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cl := mocks.NewMockTFGridClient(ctrl)

	r := Runner{
		client: cl,
	}

	t.Run("k8s_deploy_success", func(t *testing.T) {
		projectName := "project1"
		model := K8sCluster{
			Master: &K8sNode{
				Name:      "master",
				NodeID:    1,
				DiskSize:  10,
				PublicIP:  false,
				PublicIP6: true,
				Planetary: true,
				Flist:     "hamada",
				CPU:       1,
				Memory:    2,
			},
			Workers: []K8sNode{
				{
					Name:      "w1",
					NodeID:    2,
					DiskSize:  10,
					PublicIP:  true,
					PublicIP6: false,
					Planetary: false,
					Flist:     "hamada2",
					CPU:       3,
					Memory:    5,
				},
			},
			Name:   "cluster1",
			Token:  "token1",
			SSHKey: "key1",
		}

		want := K8sCluster{
			Master: &K8sNode{
				Name:        "master",
				NodeID:      1,
				DiskSize:    10,
				PublicIP:    false,
				PublicIP6:   true,
				Planetary:   true,
				Flist:       "hamada",
				CPU:         1,
				Memory:      2,
				ComputedIP4: "ip4",
				ComputedIP6: "ip6",
				WGIP:        "wgip",
				YggIP:       "yggip",
			},
			Workers: []K8sNode{
				{
					Name:        "w1",
					NodeID:      2,
					DiskSize:    10,
					PublicIP:    true,
					PublicIP6:   false,
					Planetary:   false,
					Flist:       "hamada2",
					CPU:         3,
					Memory:      5,
					ComputedIP4: "ip4",
					ComputedIP6: "ip6",
					WGIP:        "wgip",
					YggIP:       "yggip",
				},
			},
			Name:        "cluster1",
			Token:       "token1",
			NetworkName: generateNetworkName(model.Name),
			SSHKey:      "key1",
		}

		cl.
			EXPECT().
			GetProjectContracts(gomock.Any(), projectName).
			Return(graphql.Contracts{}, nil)

		ipRange, err := gridtypes.ParseIPNet("10.1.0.0/16")
		assert.NoError(t, err)

		znet := workloads.ZNet{
			Name:         generateNetworkName(model.Name),
			Nodes:        []uint32{2, 1},
			IPRange:      ipRange,
			SolutionType: projectName,
		}

		cl.EXPECT().DeployNetwork(gomock.Any(), &znet).Return(nil)

		model.NetworkName = generateNetworkName(model.Name)
		k8s := newK8sClusterFromModel(model, projectName)
		cl.EXPECT().DeployK8sCluster(gomock.Any(), &k8s).DoAndReturn(func(ctx context.Context, wl *workloads.K8sCluster) error {
			wl.Master.ComputedIP = "ip4"
			wl.Master.ComputedIP6 = "ip6"
			wl.Master.IP = "wgip"
			wl.Master.YggIP = "yggip"
			wl.Workers[0].ComputedIP = "ip4"
			wl.Workers[0].ComputedIP6 = "ip6"
			wl.Workers[0].IP = "wgip"
			wl.Workers[0].YggIP = "yggip"
			wl.NodeDeploymentID = map[uint32]uint64{1: 1}
			return nil
		})

		got, err := r.K8sDeploy(context.Background(), model, projectName)
		assert.NoError(t, err)

		assert.Equal(t, want, got)
	})

	t.Run("k8s_get_success", func(t *testing.T) {
		projectName := "project1"
		clusterName := "cluster1"
		token := "token"
		sshKey := "key"
		rmbClient := mocks.NewMockClient(ctrl)

		model := K8sCluster{
			Name: clusterName,
			Master: &K8sNode{
				Name:      "master",
				NodeID:    1,
				DiskSize:  10,
				Flist:     "hamada",
				PublicIP:  true,
				Planetary: true,
				CPU:       1,
				Memory:    2,
			},
			Workers: []K8sNode{
				{
					Name:      "w1",
					NodeID:    2,
					DiskSize:  10,
					Flist:     "hamada",
					PublicIP:  true,
					Planetary: true,
					CPU:       1,
					Memory:    2,
				},
			},
			Token:  token,
			SSHKey: sshKey,
		}

		want := K8sCluster{
			Name: clusterName,
			Master: &K8sNode{
				Name:        "master",
				NodeID:      1,
				DiskSize:    10,
				Flist:       "hamada",
				CPU:         1,
				Memory:      2,
				PublicIP:    true,
				Planetary:   true,
				ComputedIP4: "1.1.1.1/16",
				WGIP:        "3.3.3.3",
				YggIP:       "4.4.4.4",
			},
			Workers: []K8sNode{
				{
					Name:        "w1",
					NodeID:      2,
					DiskSize:    10,
					Flist:       "hamada",
					CPU:         1,
					Memory:      2,
					PublicIP:    true,
					Planetary:   true,
					ComputedIP4: "1.1.1.1/16",
					WGIP:        "3.3.3.3",
					YggIP:       "4.4.4.4",
				},
			},
			Token:       token,
			NetworkName: generateNetworkName(clusterName),
			SSHKey:      sshKey,
		}

		cl.EXPECT().GetProjectContracts(gomock.Any(), projectName).Return(graphql.Contracts{
			NodeContracts: []graphql.Contract{
				{
					ContractID: "1",
					NodeID:     1,
				},
				{
					ContractID: "2",
					NodeID:     2,
				},
			},
		}, nil)

		nodeClient := client.NewNodeClient(1, rmbClient, 10)
		cl.EXPECT().GetNodeClient(uint32(1)).Return(nodeClient, nil)
		masterWorkloads, err := generateK8sNodeWorkloads(*model.Master, sshKey, token, clusterName, false)
		assert.NoError(t, err)
		masterDeployment := workloads.NewGridDeployment(1, masterWorkloads)
		dl := gridtypes.Deployment{}
		rmbClient.EXPECT().Call(gomock.Any(), gomock.Any(), "zos.deployment.get", gomock.Any(), &dl).
			DoAndReturn(func(ctx context.Context, twin uint32, fn string, data, result interface{}) error {
				var res *gridtypes.Deployment = result.(*gridtypes.Deployment)
				*res = masterDeployment
				return nil
			})

		nodeClient = client.NewNodeClient(2, rmbClient, 10)
		cl.EXPECT().GetNodeClient(uint32(2)).Return(nodeClient, nil)
		workerWorkloads, err := generateK8sNodeWorkloads(model.Workers[0], sshKey, token, clusterName, true)
		assert.NoError(t, err)
		workerDeployment := workloads.NewGridDeployment(1, workerWorkloads)
		dl = gridtypes.Deployment{}
		rmbClient.EXPECT().Call(gomock.Any(), gomock.Any(), "zos.deployment.get", gomock.Any(), &dl).
			DoAndReturn(func(ctx context.Context, twin uint32, fn string, data, result interface{}) error {
				var res *gridtypes.Deployment = result.(*gridtypes.Deployment)
				*res = workerDeployment
				return nil
			})

		got, err := r.K8sGet(context.Background(), clusterName, projectName)
		assert.NoError(t, err)

		assert.Equal(t, want, got)
	})
}

func generateK8sNodeWorkloads(node K8sNode, sshKey string, token string, clusterName string, isWorker bool) ([]gridtypes.Workload, error) {
	envVars := map[string]string{
		"SSH_KEY":           sshKey,
		"K3S_TOKEN":         token,
		"K3S_DATA_DIR":      "/mydisk",
		"K3S_FLANNEL_IFACE": "eth0",
		"K3S_NODE_NAME":     node.Name,
		"K3S_URL":           "",
	}

	if isWorker {
		envVars["K3S_URL"] = "master_ip"
	}

	vmRes := zos.ZMachineResult{
		YggIP: "4.4.4.4",
	}

	vmResultBytes, err := json.Marshal(vmRes)
	if err != nil {
		return nil, err
	}

	pubIPRes := zos.PublicIPResult{
		IP: gridtypes.MustParseIPNet("1.1.1.1/16"),
	}

	pubIPBytes, err := json.Marshal(pubIPRes)
	if err != nil {
		return nil, err
	}

	return []gridtypes.Workload{
		{
			Version: 0,
			Name:    gridtypes.Name(node.Name),
			Type:    zos.ZMachineType,
			Data: gridtypes.MustMarshal(zos.ZMachine{
				FList: node.Flist,
				Network: zos.MachineNetwork{
					PublicIP: gridtypes.Name(fmt.Sprintf("%sip", node.Name)),
					Interfaces: []zos.MachineInterface{
						{
							Network: gridtypes.Name(generateNetworkName(clusterName)),
							IP:      net.ParseIP("3.3.3.3"),
						},
					},
					Planetary: true,
				},
				ComputeCapacity: zos.MachineCapacity{
					CPU:    uint8(node.CPU),
					Memory: gridtypes.Unit(uint(node.Memory)) * gridtypes.Megabyte,
				},
				Entrypoint: "/sbin/zinit init",
				Mounts: []zos.MachineMount{
					{
						Name:       gridtypes.Name(fmt.Sprintf("%sdisk", node.Name)),
						Mountpoint: "/mydisk",
					},
				},
				Env: envVars,
			}),
			Result: gridtypes.Result{
				Created: gridtypes.Now(),
				State:   gridtypes.StateOk,
				Data:    json.RawMessage(vmResultBytes),
			},
		},
		{
			Version: 0,
			Name:    gridtypes.Name(fmt.Sprintf("%sip", node.Name)),
			Type:    zos.PublicIPType,
			Data: gridtypes.MustMarshal(zos.PublicIP{
				V4: node.PublicIP,
			}),
			Result: gridtypes.Result{
				Created: gridtypes.Now(),
				State:   gridtypes.StateOk,
				Data:    json.RawMessage(pubIPBytes),
			},
		},
		{
			Name:        gridtypes.Name(fmt.Sprintf("%sdisk", node.Name)),
			Version:     0,
			Type:        zos.ZMountType,
			Description: "",
			Data: gridtypes.MustMarshal(zos.ZMount{
				Size: gridtypes.Unit(node.DiskSize) * gridtypes.Gigabyte,
			}),
		},
	}, nil
}
