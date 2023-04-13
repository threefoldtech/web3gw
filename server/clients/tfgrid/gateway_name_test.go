package tfgrid

import (
	"context"
	"encoding/json"
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

func TestGatewayName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cl := mocks.NewMockTFGridClient(ctrl)

	r := Runner{
		client: cl,
	}

	t.Run("gateway_name_deploy_success", func(t *testing.T) {
		projectName := "project1"
		want := GatewayNameModel{
			NodeID: 1,
			Name:   "hamada",
			Backends: []zos.Backend{
				"backend1",
				"b2",
			},
			TLSPassthrough: false,
			Description:    "desc1",
			FQDN:           "hamada.name.com",
			NameContractID: 1,
			ContractID:     2,
		}

		model := GatewayNameModel{
			NodeID: 1,
			Name:   "hamada",
			Backends: []zos.Backend{
				"backend1",
				"b2",
			},
			TLSPassthrough: false,
			Description:    "desc1",
		}

		cl.
			EXPECT().
			GetProjectContracts(gomock.Any(), projectName).
			Return(graphql.Contracts{}, nil)

		gw := newGWNameProxyFromModel(model, projectName)

		cl.EXPECT().DeployGWName(gomock.Any(), &gw).DoAndReturn(func(ctx context.Context, wl *workloads.GatewayNameProxy) error {
			wl.NameContractID = 1
			wl.ContractID = 2
			return nil
		})

		rmbClient := mocks.NewMockClient(ctrl)
		nodeClient := client.NewNodeClient(uint32(1), rmbClient, 10)
		cl.EXPECT().GetNodeClient(uint32(1)).Return(nodeClient, nil)

		cfg := client.PublicConfig{
			Domain: "name.com",
		}
		rmbClient.
			EXPECT().
			Call(gomock.Any(), gomock.Any(), "zos.network.public_config_get", gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, twin uint32, fn string, data, result interface{}) error {
				var res *client.PublicConfig = result.(*client.PublicConfig)
				*res = cfg
				return nil
			})
		got, err := r.GatewayNameDeploy(context.Background(), model, projectName)
		assert.NoError(t, err)

		assert.Equal(t, want, got)
	})

	t.Run("gateway_name_get_success", func(t *testing.T) {
		projectName := "project1"
		nodeID := uint32(1)
		nameContractID := uint64(1)
		nodeContractID := uint64(2)
		want := GatewayNameModel{
			NodeID: 1,
			Name:   "hamada",
			Backends: []zos.Backend{
				"backend1",
				"b2",
			},
			TLSPassthrough: false,
			Description:    "desc1",
			FQDN:           "hamada.name.com",
			NameContractID: nameContractID,
			ContractID:     nodeContractID,
		}

		cl.EXPECT().GetProjectContracts(gomock.Any(), projectName).Return(graphql.Contracts{
			NodeContracts: []graphql.Contract{
				{
					ContractID: "2",
					NodeID:     nodeID,
				},
			},
			NameContracts: []graphql.Contract{
				{
					ContractID: "1",
					NodeID:     nodeID,
				},
			},
		}, nil)

		result := zos.GatewayProxyResult{
			FQDN: want.FQDN,
		}
		resultBytes, err := json.Marshal(result)
		assert.NoError(t, err)

		rmbClient := mocks.NewMockClient(ctrl)
		workload := gridtypes.Workload{
			Version: 0,
			Name:    "name1",
			Type:    zos.GatewayNameProxyType,
			Data: gridtypes.MustMarshal(zos.GatewayNameProxy{
				Name: want.Name,
				GatewayBase: zos.GatewayBase{
					TLSPassthrough: want.TLSPassthrough,
					Backends:       want.Backends,
				},
			}),
			Description: want.Description,
			Result: gridtypes.Result{
				Created: gridtypes.Now(),
				State:   gridtypes.StateOk,
				Data:    json.RawMessage(resultBytes),
			},
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

		got, err := r.GatewayNameGet(context.Background(), projectName)
		assert.NoError(t, err)

		assert.Equal(t, want, got)
	})
}
