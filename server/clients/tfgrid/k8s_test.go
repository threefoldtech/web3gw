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

func TestK8s(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cl := mocks.NewMockTFGridClient(ctrl)

	r := Client{
		client: cl,
	}

	t.Run("k8s_deploy_success", func(t *testing.T) {
		modelName := "cluster1"
		projectName := generateProjectName(modelName)
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
			Name:   modelName,
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
			Name:        modelName,
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
			Nodes:        []uint32{1, 2},
			IPRange:      ipRange,
			SolutionType: projectName,
		}

		cl.EXPECT().DeployNetwork(gomock.Any(), &znet).Return(nil)

		model.NetworkName = generateNetworkName(model.Name)
		k8s := newK8sClusterFromModel(model)
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

		got, err := r.K8sDeploy(context.Background(), model)
		assert.NoError(t, err)

		assert.Equal(t, want, got)
	})

	t.Run("k8s_get_success", func(t *testing.T) {
		clusterName := "cluster1"
		projectName := generateProjectName(clusterName)
		token := "token"
		sshKey := "key"

		want := K8sCluster{
			Name: clusterName,
			Master: &K8sNode{
				Name:        "master",
				NodeID:      1,
				FarmID:      1,
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
					FarmID:      1,
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

		networkDeploymentData, err := json.Marshal(workloads.DeploymentData{
			Type: "network",
			Name: generateNetworkName(clusterName),
		})
		assert.NoError(t, err)

		nodesDeploymentData, err := json.Marshal(workloads.DeploymentData{
			Type: "kubernetes",
			Name: clusterName,
		})
		assert.NoError(t, err)

		cl.EXPECT().GetProjectContracts(gomock.Any(), projectName).Return(graphql.Contracts{
			NodeContracts: []graphql.Contract{
				{
					ContractID:     "1",
					NodeID:         1,
					DeploymentData: string(nodesDeploymentData),
				},
				{
					ContractID:     "2",
					NodeID:         2,
					DeploymentData: string(nodesDeploymentData),
				},
				{
					ContractID:     "3",
					NodeID:         1,
					DeploymentData: string(networkDeploymentData),
				},
				{
					ContractID:     "4",
					NodeID:         2,
					DeploymentData: string(networkDeploymentData),
				},
			},
		}, nil)

		cl.EXPECT().SetNodeDeploymentState(map[uint32]state.ContractIDs{1: {1}, 2: {2}})
		cl.EXPECT().LoadK8s("master", []uint32{1, 2}).Return(workloads.K8sCluster{
			Master: &workloads.K8sNode{
				Name:        "master",
				Node:        1,
				DiskSize:    10,
				Flist:       "hamada",
				CPU:         1,
				Memory:      2,
				PublicIP:    true,
				Planetary:   true,
				ComputedIP:  "1.1.1.1/16",
				IP:          "3.3.3.3",
				YggIP:       "4.4.4.4",
				NetworkName: generateNetworkName(clusterName),
				Token:       token,
				SSHKey:      sshKey,
			},
			Workers: []workloads.K8sNode{
				{
					Name:        "w1",
					Node:        2,
					DiskSize:    10,
					Flist:       "hamada",
					CPU:         1,
					Memory:      2,
					PublicIP:    true,
					Planetary:   true,
					ComputedIP:  "1.1.1.1/16",
					IP:          "3.3.3.3",
					YggIP:       "4.4.4.4",
					NetworkName: generateNetworkName(clusterName),
					Token:       token,
					SSHKey:      sshKey,
				},
			},
			Token:        token,
			NetworkName:  generateNetworkName(clusterName),
			SSHKey:       sshKey,
			SolutionType: projectName,
		}, nil)

		cl.EXPECT().GetNodeFarm(uint32(1)).Return(uint32(1), nil).AnyTimes()
		cl.EXPECT().GetNodeFarm(uint32(2)).Return(uint32(1), nil).AnyTimes()

		got, err := r.K8sGet(context.Background(), GetClusterParams{
			ClusterName: clusterName,
			MasterName:  "master",
		})
		assert.NoError(t, err)

		assert.Equal(t, want, got)
	})
}
