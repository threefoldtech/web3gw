package tfgrid

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/graphql"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/state"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/workloads"
	"github.com/threefoldtech/web3gw/web3gw/server/clients/tfgrid/mocks"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"
)

func TestGatewayName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cl := mocks.NewMockTFGridClient(ctrl)

	r := Client{
		GridClient: cl,
		Projects:   make(map[string]ProjectState),
	}

	t.Run("gateway_name_deploy_success", func(t *testing.T) {
		modelName := "hamada"
		projectName := projectNameFromName(modelName)
		nodeID := uint32(1)
		nameContractID := uint64(1)
		contractID := uint64(2)
		domain := "name.com"
		want := GatewayNameModel{
			NodeID: nodeID,
			Name:   modelName,
			Backends: []zos.Backend{
				"backend1",
				"b2",
			},
			TLSPassthrough: false,
			Description:    "desc1",
			FQDN:           fmt.Sprintf("%s.%s", modelName, domain),
			NameContractID: nameContractID,
			ContractID:     contractID,
		}

		model := GatewayNameModel{
			NodeID: nodeID,
			Name:   modelName,
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

		cl.EXPECT().SetContractState(map[uint32]state.ContractIDs{nodeID: {contractID}})

		cl.EXPECT().LoadGatewayName(modelName, nodeID).Return(workloads.GatewayNameProxy{
			NodeID: nodeID,
			Name:   modelName,
			Backends: []zos.Backend{
				"backend1",
				"b2",
			},
			TLSPassthrough: false,
			Description:    "desc1",
			FQDN:           fmt.Sprintf("%s.%s", modelName, domain),
			NameContractID: nameContractID,
			ContractID:     contractID,
		}, nil)

		gw := toGridGWName(model)

		cl.EXPECT().DeployGWName(gomock.Any(), &gw).DoAndReturn(func(ctx context.Context, wl *workloads.GatewayNameProxy) error {
			wl.NameContractID = nameContractID
			wl.ContractID = contractID
			return nil
		})

		got, err := r.GatewayNameDeploy(context.Background(), model)
		assert.NoError(t, err)

		assert.Equal(t, want, got)
	})

	t.Run("gateway_name_get_success", func(t *testing.T) {
		modelName := "hamada2"
		projectName := projectNameFromName(modelName)
		nodeID := uint32(1)
		nameContractID := uint64(1)
		nodeContractID := uint64(2)
		domain := "name.com"

		want := GatewayNameModel{
			NodeID: nodeID,
			Name:   modelName,
			Backends: []zos.Backend{
				"backend1",
				"b2",
			},
			TLSPassthrough: false,
			Description:    "desc1",
			FQDN:           fmt.Sprintf("%s.%s", modelName, domain),
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

		cl.EXPECT().SetContractState(map[uint32]state.ContractIDs{nodeID: {nodeContractID}})
		cl.EXPECT().LoadGatewayName(modelName, nodeID).Return(workloads.GatewayNameProxy{
			NodeID: nodeID,
			Name:   modelName,
			Backends: []zos.Backend{
				"backend1",
				"b2",
			},
			TLSPassthrough: false,
			Description:    "desc1",
			FQDN:           fmt.Sprintf("%s.%s", modelName, domain),
			NameContractID: nameContractID,
			ContractID:     nodeContractID,
		}, nil)

		got, err := r.GatewayNameGet(context.Background(), modelName)
		assert.NoError(t, err)

		assert.Equal(t, want, got)
	})
}
