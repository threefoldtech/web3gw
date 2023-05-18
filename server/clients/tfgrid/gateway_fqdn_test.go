package tfgrid

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/graphql"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/state"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/workloads"
	"github.com/threefoldtech/web3_proxy/server/clients/tfgrid/mocks"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"
)

func TestGatewayFQDN(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cl := mocks.NewMockTFGridClient(ctrl)

	r := Client{
		GridClient: cl,
		Projects:   make(map[string]ProjectState),
	}

	t.Run("fqdn_deploy_success", func(t *testing.T) {
		modelName := "name1"
		projectName := generateProjectName(modelName)
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
			Name:           modelName,
			TLSPassthrough: false,
			Description:    "description1",
			ContractID:     contractID,
		}

		wl := workloads.GatewayFQDNProxy{
			NodeID:         nodeID,
			Backends:       backends,
			FQDN:           fqdn,
			Name:           modelName,
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

		got, err := r.GatewayFQDNDeploy(context.Background(), model)
		assert.NoError(t, err)

		assert.Equal(t, want, got, "target gateway fqdn is not equal to result gateway fqdn")
	})

	t.Run("fqdn_deploy_fail_project_name_not_unique", func(t *testing.T) {
		modelName := "name2"
		projectName := generateProjectName(modelName)
		fqdnModel := GatewayFQDNModel{
			NodeID: 1,
			Backends: []zos.Backend{
				"backend1",
				"backend2",
			},
			FQDN:           "hamada.com",
			Name:           modelName,
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

		_, err := r.GatewayFQDNDeploy(context.Background(), fqdnModel)
		assert.Error(t, err)
	})

	t.Run("fqdn_get_success", func(t *testing.T) {
		modelName := "name3"
		projectName := generateProjectName(modelName)
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

		cl.EXPECT().GetProjectContracts(context.Background(), projectName).Return(graphql.Contracts{
			NodeContracts: []graphql.Contract{
				{
					ContractID: "1",
					NodeID:     nodeID,
				},
			},
		}, nil)

		cl.EXPECT().SetContractState(map[uint32]state.ContractIDs{nodeID: {contractID}})
		cl.EXPECT().LoadGatewayFQDN(modelName, nodeID).Return(workloads.GatewayFQDNProxy{
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
		}, nil)

		got, err := r.GatewayFQDNGet(context.Background(), modelName)
		assert.NoError(t, err)

		assert.Equal(t, want, got)
	})
}
