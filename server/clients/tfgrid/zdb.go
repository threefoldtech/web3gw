package tfgrid

import (
	"context"
	"math/rand"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/state"
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

func (c *Client) ZDBDeploy(ctx context.Context, zdb ZDB) (ZDB, error) {
	// validate no workloads with the same name
	if err := c.validateProjectName(ctx, zdb.Name); err != nil {
		return ZDB{}, err
	}

	if err := c.ensureZDBNodeIDExist(zdb); err != nil {
		return ZDB{}, err
	}

	gridZDB := toGridZDB(zdb)

	if err := c.deployZDB(ctx, &gridZDB, zdb.NodeID); err != nil {
		return ZDB{}, err
	}

	return c.ZDBGet(ctx, zdb.Name)
}

func (c *Client) deployZDB(ctx context.Context, gridZDB *workloads.ZDB, nodeID uint32) error {
	log.Debug().Msgf("Deploying zdb: %+v", *gridZDB)

	dl := workloads.NewDeployment(gridZDB.Name, nodeID, generateProjectName(gridZDB.Name), nil, "", nil, []workloads.ZDB{*gridZDB}, nil, nil)
	if err := c.client.DeployDeployment(ctx, &dl); err != nil {
		return errors.Wrapf(err, "failed to deploy zdb with name: %s", gridZDB.Name)
	}

	projectName := generateProjectName(gridZDB.Name)

	c.Projects[projectName] = ProjectState{
		nodeContracts: map[uint32]state.ContractIDs{
			nodeID: {dl.NodeDeploymentID[nodeID]},
		},
	}

	return nil
}

func (c *Client) ZDBDelete(ctx context.Context, modelName string) error {
	if err := c.cancelModel(ctx, modelName); err != nil {
		return errors.Wrapf(err, "Failed to cancel zdb with name: %s", modelName)
	}

	return nil
}

func (r *Client) ZDBGet(ctx context.Context, modelName string) (ZDB, error) {
	log.Debug().Msgf("retreiving zdb %s", modelName)

	zdb, nodeID, err := r.loadZDB(ctx, modelName)
	if err != nil {
		return ZDB{}, err
	}

	ret := ZDBToModel(zdb, nodeID)

	return ret, nil
}

func toGridZDB(zdb ZDB) workloads.ZDB {
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
