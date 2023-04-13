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

func TestZDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cl := mocks.NewMockTFGridClient(ctrl)

	r := Runner{
		client: cl,
	}

	t.Run("zdb_deploy_success", func(t *testing.T) {
		projectName := "project1"
		rmbClient := mocks.NewMockClient(ctrl)

		model := ZDB{
			NodeID:   1,
			Name:     "zdb",
			Password: "pass",
			Public:   true,
			Size:     10,
			Mode:     "seq",
		}

		want := ZDB{
			NodeID:    1,
			Name:      "zdb",
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
			newClientWorkloadFromZDB(model),
		}

		clientDeployment := workloads.NewDeployment(model.Name, model.NodeID, projectName, nil, "", nil, zdbs, nil, nil)
		contractID := uint64(1)
		cl.EXPECT().DeployDeployment(gomock.Any(), &clientDeployment).Return(contractID, nil)

		nodeClient := client.NewNodeClient(1, rmbClient, 10)
		cl.EXPECT().GetNodeClient(uint32(1)).Return(nodeClient, nil)
		zdbRes := zos.ZDBResult{
			Namespace: want.Namespace,
			IPs:       want.IPs,
			Port:      uint(want.Port),
		}

		zdbResBytes, err := json.Marshal(zdbRes)
		assert.NoError(t, err)

		wls := []gridtypes.Workload{
			{
				Version: 0,
				Name:    "zdb",
				Type:    zos.ZDBType,
				Data: gridtypes.MustMarshal(zos.ZDB{
					Size:     gridtypes.Unit(model.Size) * gridtypes.Gigabyte,
					Mode:     zos.ZDBMode(model.Mode),
					Password: model.Password,
					Public:   model.Public,
				}),
				Result: gridtypes.Result{
					Created: gridtypes.Now(),
					State:   gridtypes.StateOk,
					Data:    json.RawMessage(zdbResBytes),
				},
			},
		}
		zosDeployment := workloads.NewGridDeployment(1, wls)
		dl := gridtypes.Deployment{}
		rmbClient.EXPECT().Call(gomock.Any(), gomock.Any(), "zos.deployment.get", gomock.Any(), &dl).
			DoAndReturn(func(ctx context.Context, twin uint32, fn string, data, result interface{}) error {
				var res *gridtypes.Deployment = result.(*gridtypes.Deployment)
				*res = zosDeployment
				return nil
			})

		got, err := r.ZDBDeploy(context.Background(), model, projectName)
		assert.NoError(t, err)

		assert.Equal(t, want, got)
	})

	t.Run("zdb_get_success", func(t *testing.T) {
		projectName := "project1"
		rmbClient := mocks.NewMockClient(ctrl)

		want := ZDB{
			NodeID:    1,
			Name:      "zdb",
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

		nodeClient := client.NewNodeClient(1, rmbClient, 10)
		cl.EXPECT().GetNodeClient(uint32(1)).Return(nodeClient, nil)
		zdbRes := zos.ZDBResult{
			Namespace: want.Namespace,
			IPs:       want.IPs,
			Port:      uint(want.Port),
		}

		zdbResBytes, err := json.Marshal(zdbRes)
		assert.NoError(t, err)

		wls := []gridtypes.Workload{
			{
				Version: 0,
				Name:    "zdb",
				Type:    zos.ZDBType,
				Data: gridtypes.MustMarshal(zos.ZDB{
					Size:     gridtypes.Unit(10) * gridtypes.Gigabyte,
					Mode:     zos.ZDBMode("seq"),
					Password: "pass",
					Public:   true,
				}),
				Result: gridtypes.Result{
					Created: gridtypes.Now(),
					State:   gridtypes.StateOk,
					Data:    json.RawMessage(zdbResBytes),
				},
			},
		}
		zosDeployment := workloads.NewGridDeployment(1, wls)
		dl := gridtypes.Deployment{}
		rmbClient.EXPECT().Call(gomock.Any(), gomock.Any(), "zos.deployment.get", gomock.Any(), &dl).
			DoAndReturn(func(ctx context.Context, twin uint32, fn string, data, result interface{}) error {
				var res *gridtypes.Deployment = result.(*gridtypes.Deployment)
				*res = zosDeployment
				return nil
			})

		got, err := r.ZDBGet(context.Background(), projectName)
		assert.NoError(t, err)

		assert.Equal(t, want, got)
	})
}
