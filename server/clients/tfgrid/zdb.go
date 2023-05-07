package tfgrid

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/workloads"
	proxyTypes "github.com/threefoldtech/tfgrid-sdk-go/grid-proxy/pkg/types"
)

type ZDB struct {
	NodeID      uint32 `json:"node_id"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Public      bool   `json:"public"`
	Size        int    `json:"size"`
	Description string `json:"description"`
	Mode        string `json:"mode"`

	// computed
	Port      uint32   `json:"port"`
	Namespace string   `json:"namespace"`
	IPs       []string `json:"ips"`
}

func (r *Client) ZDBDeploy(ctx context.Context, zdb ZDB) (ZDB, error) {
	projectName := generateProjectName(zdb.Name)

	// validate no workloads with the same name
	if err := r.validateProjectName(ctx, projectName); err != nil {
		return ZDB{}, err
	}

	if err := r.ensureZDBNodeIDExist(zdb); err != nil {
		return ZDB{}, err
	}

	// deploy
	zdbs := []workloads.ZDB{
		ZDBFromModel(zdb),
	}
	log.Info().Msgf("Deploying zdb: %+v", zdbs)

	clientDeployment := workloads.NewDeployment(zdb.Name, zdb.NodeID, projectName, nil, "", nil, zdbs, nil, nil)
	contractID, err := r.client.DeployDeployment(ctx, &clientDeployment)
	if err != nil {
		return ZDB{}, errors.Wrapf(err, "failed to deploy zdb with name: %s", zdb.Name)
	}

	r.client.SetNodeDeploymentState(map[uint32][]uint64{zdb.NodeID: {contractID}})

	z, err := r.client.LoadZDB(zdb.NodeID, zdb.Name)
	if err != nil {
		return ZDB{}, errors.Wrap(err, "failed to load zdb")
	}

	ret := ZDBToModel(z, zdb.NodeID)

	return ret, nil
}

func (r *Client) ZDBDelete(ctx context.Context, projectName string) error {
	if err := r.client.CancelProject(ctx, projectName); err != nil {
		return errors.Wrapf(err, "Failed to cancel cluster with name: %s", projectName)
	}

	return nil
}

func (r *Client) ZDBGet(ctx context.Context, modelName string) (ZDB, error) {
	projectName := generateProjectName(modelName)

	contracts, err := r.client.GetProjectContracts(ctx, projectName)
	if err != nil {
		return ZDB{}, errors.Wrapf(err, "failed to get contracts for project: %s", projectName)
	}

	if len(contracts.NodeContracts) != 1 {
		return ZDB{}, fmt.Errorf("contracts of project %s should be 1, but %d were found", projectName, len(contracts.NodeContracts))
	}

	contract := contracts.NodeContracts[0]
	cid, err := strconv.ParseUint(contract.ContractID, 10, 64)
	if err != nil {
		return ZDB{}, errors.Wrapf(err, "failed to parse contract Id: %s", contract.ContractID)
	}

	r.client.SetNodeDeploymentState(map[uint32][]uint64{contract.NodeID: {cid}})

	zdb, err := r.client.LoadZDB(contract.NodeID, modelName)
	if err != nil {
		return ZDB{}, err
	}

	ret := ZDBToModel(zdb, contract.NodeID)

	return ret, nil
}

func ZDBFromModel(zdb ZDB) workloads.ZDB {
	return workloads.ZDB{
		Name:        zdb.Name,
		Password:    zdb.Password,
		Public:      zdb.Public,
		Size:        zdb.Size,
		Description: zdb.Description,
		Mode:        zdb.Mode,
		Port:        zdb.Port,
		Namespace:   zdb.Namespace,
	}
}

func ZDBToModel(wl workloads.ZDB, nodeID uint32) ZDB {
	return ZDB{
		Name:        wl.Name,
		NodeID:      nodeID,
		Password:    wl.Password,
		Public:      wl.Public,
		Size:        wl.Size,
		Description: wl.Description,
		Mode:        wl.Mode,
		Port:        wl.Port,
		Namespace:   wl.Namespace,
		IPs:         wl.IPs,
	}
}

func (r *Client) ensureZDBNodeIDExist(zdb ZDB) error {
	// capacity filter
	if zdb.NodeID == 0 {
		nodeId, err := r.getNodeForZdb(uint64(zdb.Size))
		if err != nil {
			return errors.Wrapf(err, "Couldn't find a gateway node")
		}

		zdb.NodeID = nodeId
	}
	return nil
}

func (r *Client) getNodeForZdb(size uint64) (uint32, error) {
	options := proxyTypes.NodeFilter{
		Status:  &Status,
		FreeHRU: &size,
	}

	nodes, count, err := r.client.FilterNodes(options, proxyTypes.Limit{})
	if err != nil || count == 0 {
		return 0, errors.Wrapf(err, "Couldn't find node for the provided filters: %+v", options)
	}

	return uint32(nodes[rand.Intn(count)].NodeID), nil
}
