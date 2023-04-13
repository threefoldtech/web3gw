package tfgrid

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/graphql"
	client "github.com/threefoldtech/tfgrid-sdk-go/grid-client/node"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/workloads"
	"github.com/threefoldtech/web3_proxy/server/clients/tfgrid/mocks"
	"github.com/threefoldtech/zos/pkg/gridtypes"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"
)

func TestGatewayFQDN(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cl := mocks.NewMockTFGridClient(ctrl)

	r := Runner{
		client: cl,
	}

	t.Run("fqdn_deploy_success", func(t *testing.T) {
		projectName := "project1"
		nodeID := uint32(1)
		contractID := uint64(1)
		backends := []zos.Backend{
			"backend1",
			"backend2",
		}
		fqdn := "hamada.com"

		want := GatewayFQDNModel{
			NodeID:         nodeID,
			Backends:       backends,
			FQDN:           fqdn,
			Name:           "name1",
			TLSPassthrough: false,
			Description:    "description1",
			ContractID:     contractID,
		}

		wl := workloads.GatewayFQDNProxy{
			NodeID:         nodeID,
			Backends:       backends,
			FQDN:           fqdn,
			Name:           "name1",
			TLSPassthrough: false,
			Description:    "description1",
			SolutionType:   projectName,
		}

		cl.
			EXPECT().
			GetProjectContracts(gomock.Any(), projectName).
			Return(graphql.Contracts{}, nil)

		cl.EXPECT().DeployGWFQDN(gomock.Any(), &wl).DoAndReturn(func(ctx context.Context, wl *workloads.GatewayFQDNProxy) error {
			wl.ContractID = 1
			return nil
		})

		model := GatewayFQDNModel{
			NodeID:         want.NodeID,
			Backends:       want.Backends,
			FQDN:           want.FQDN,
			Name:           want.Name,
			TLSPassthrough: want.TLSPassthrough,
			Description:    want.Description,
		}

		got, err := r.GatewayFQDNDeploy(context.Background(), model, projectName)
		assert.NoError(t, err)

		assert.Equal(t, want, got, "target gateway fqdn is not equal to result gateway fqdn")
	})

	t.Run("fqdn_deploy_fail_project_name_not_unique", func(t *testing.T) {
		projectName := "project1"
		fqdnModel := GatewayFQDNModel{
			NodeID: 1,
			Backends: []zos.Backend{
				"backend1",
				"backend2",
			},
			FQDN:           "hamada.com",
			Name:           "name1",
			TLSPassthrough: false,
			Description:    "description1",
		}

		cl.EXPECT().GetProjectContracts(gomock.Any(), projectName).Return(graphql.Contracts{
			NodeContracts: []graphql.Contract{
				{
					ContractID: "1",
				},
			},
		}, nil)

		_, err := r.GatewayFQDNDeploy(context.Background(), fqdnModel, projectName)
		assert.Error(t, err)
	})

	t.Run("fqdn_get_success", func(t *testing.T) {
		projectName := "project1"
		nodeID := uint32(1)
		contractID := uint64(1)
		want := GatewayFQDNModel{
			NodeID: nodeID,
			Backends: []zos.Backend{
				"backend1",
				"backend2",
			},
			FQDN:           "hamada.com",
			Name:           "name1",
			TLSPassthrough: false,
			Description:    "description1",
			ContractID:     contractID,
		}

		cl.EXPECT().GetProjectContracts(gomock.Any(), projectName).Return(graphql.Contracts{
			NodeContracts: []graphql.Contract{
				{
					ContractID: "1",
					NodeID:     nodeID,
				},
			},
		}, nil)

		rmbClient := mocks.NewMockClient(ctrl)
		workload := gridtypes.Workload{
			Version: 0,
			Name:    "name1",
			Type:    zos.GatewayFQDNProxyType,
			Data: gridtypes.MustMarshal(zos.GatewayFQDNProxy{
				FQDN: "hamada.com",
				GatewayBase: zos.GatewayBase{
					TLSPassthrough: false,
					Backends: []zos.Backend{
						"backend1",
						"backend2",
					},
				},
			}),
			Description: "description1",
		}
		resDeployment := workloads.NewGridDeployment(1, []gridtypes.Workload{workload})
		dl := gridtypes.Deployment{}
		rmbClient.EXPECT().Call(gomock.Any(), gomock.Any(), "zos.deployment.get", gomock.Any(), &dl).
			DoAndReturn(func(ctx context.Context, twin uint32, fn string, data, result interface{}) error {
				var res *gridtypes.Deployment = result.(*gridtypes.Deployment)
				*res = resDeployment
				return nil
			})

		nodeClient := client.NewNodeClient(nodeID, rmbClient, 10)
		cl.EXPECT().GetNodeClient(nodeID).Return(nodeClient, nil)

		got, err := r.GatewayFQDNGet(context.Background(), projectName)
		assert.NoError(t, err)

		assert.Equal(t, want, got)
	})
}
