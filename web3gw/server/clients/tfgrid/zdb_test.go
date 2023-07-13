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
)

func TestZDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cl := mocks.NewMockTFGridClient(ctrl)

	r := Client{
		GridClient: cl,
		Projects:   map[string]ProjectState{},
	}

	t.Run("zdb_deploy_success", func(t *testing.T) {
		modelName := "zdb"
		projectName := generateProjectName(modelName)
		nodeID := uint32(1)
		contractID := uint64(1)

		model := ZDB{
			NodeID:   nodeID,
			Name:     modelName,
			Password: "pass",
			Public:   true,
			Size:     10,
			Mode:     "seq",
		}

		want := ZDB{
			NodeID:    nodeID,
			Name:      modelName,
			Password:  "pass",
			Public:    true,
			Size:      10,
			Mode:      "seq",
			Port:      9900,
			Namespace: "namespace",
			IPs:       []string{"1.1.1.1", "2.2.2.2"},
		}

		cl.
			EXPECT().
			GetProjectContracts(gomock.Any(), projectName).
			Return(graphql.Contracts{}, nil)

		zdbs := []workloads.ZDB{
			toGridZDB(model),
		}

		clientDeployment := workloads.NewDeployment(model.Name, model.NodeID, projectName, nil, "", nil, zdbs, nil, nil)
		cl.EXPECT().DeployDeployment(gomock.Any(), &clientDeployment).DoAndReturn(func(ctx context.Context, d *workloads.Deployment) error {
			d.ContractID = contractID
			d.NodeDeploymentID = map[uint32]uint64{nodeID: contractID}
			return nil
		})

		cl.EXPECT().SetContractState(map[uint32]state.ContractIDs{nodeID: {contractID}})

		cl.EXPECT().LoadZDB(modelName, nodeID).Return(workloads.ZDB{
			Name:      modelName,
			Password:  "pass",
			Public:    true,
			Size:      10,
			Mode:      "seq",
			Port:      9900,
			Namespace: "namespace",
			IPs:       []string{"1.1.1.1", "2.2.2.2"},
		}, nil)

		got, err := r.ZDBDeploy(context.Background(), model)
		assert.NoError(t, err)

		assert.Equal(t, want, got)
	})

	t.Run("zdb_get_success", func(t *testing.T) {
		modelName := "zdb2"
		projectName := generateProjectName(modelName)
		nodeID := uint32(1)
		contractID := uint64(1)

		want := ZDB{
			NodeID:    nodeID,
			Name:      modelName,
			Password:  "pass",
			Public:    true,
			Size:      10,
			Mode:      "seq",
			Port:      9900,
			Namespace: "namespace",
			IPs:       []string{"1.1.1.1", "2.2.2.2"},
		}

		cl.EXPECT().GetProjectContracts(gomock.Any(), projectName).Return(graphql.Contracts{
			NodeContracts: []graphql.Contract{
				{
					ContractID: "1",
					NodeID:     1,
				},
			},
		}, nil)

		cl.EXPECT().SetContractState(map[uint32]state.ContractIDs{nodeID: {contractID}})

		cl.EXPECT().LoadZDB(modelName, nodeID).Return(workloads.ZDB{
			Name:      modelName,
			Password:  "pass",
			Public:    true,
			Size:      10,
			Mode:      "seq",
			Port:      9900,
			Namespace: "namespace",
			IPs:       []string{"1.1.1.1", "2.2.2.2"},
		}, nil)

		got, err := r.ZDBGet(context.Background(), modelName)
		assert.NoError(t, err)

		assert.Equal(t, want, got)
	})
}
