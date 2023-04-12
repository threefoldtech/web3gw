package tfgrid

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/threefoldtech/grid3-go/workloads"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"
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

func (r *Runner) ZDBDeploy(ctx context.Context, zdb ZDB, projectName string) (ZDB, error) {
	// validate no workloads with the same name
	if err := r.validateProjectName(ctx, projectName); err != nil {
		return ZDB{}, err
	}

	// deploy
	zdbs := []workloads.ZDB{
		newClientWorkloadFromZDB(zdb),
	}
	log.Info().Msgf("Deploying zdb: %+v", zdbs)

	clientDeployment := workloads.NewDeployment(zdb.Name, zdb.NodeID, projectName, nil, "", nil, zdbs, nil, nil)
	contractID, err := r.client.DeployDeployment(ctx, &clientDeployment)
	if err != nil {
		return ZDB{}, errors.Wrapf(err, "failed to deploy zdb with name: %s", zdb.Name)
	}

	// get the result with the computed values
	nodeClient, err := r.client.GetNodeClient(zdb.NodeID)
	if err != nil {
		return ZDB{}, err
	}

	dl, err := nodeClient.DeploymentGet(ctx, contractID)
	if err != nil {
		return ZDB{}, errors.Wrapf(err, "failed to retreive deployment with contract id %d", contractID)
	}

	if len(dl.Workloads) != 1 {
		return ZDB{}, errors.Wrapf(err, "deployment %d should have 1 workload, but %d were found", contractID, len(dl.Workloads))
	}

	loadedZDB, err := workloads.NewZDBFromWorkload(&dl.Workloads[0])
	if err != nil {
		return ZDB{}, errors.Wrapf(err, "failed to construct zdb from workload")
	}

	result := newZDBFromClientZDB(loadedZDB)
	result.NodeID = zdb.NodeID

	// NOTE: clean the state after deploying
	return result, nil
}

func (r *Runner) ZDBDelete(ctx context.Context, projectName string) error {
	// TODO: fix canceling
	if err := r.client.CancelProject(ctx, projectName); err != nil {
		return errors.Wrapf(err, "Failed to cancel cluster with name: %s", projectName)
	}

	return nil
}

func (r *Runner) ZDBGet(ctx context.Context, projectName string) (ZDB, error) {
	// get the contract
	contracts, err := r.client.GetProjectContracts(ctx, projectName)
	if err != nil {
		return ZDB{}, errors.Wrapf(err, "failed to get contracts for project: %s", projectName)
	}

	if len(contracts.NodeContracts) != 1 {
		return ZDB{}, fmt.Errorf("contracts of project %s should be 1, but %d were found", projectName, len(contracts.NodeContracts))
	}

	contract := contracts.NodeContracts[0]

	cl, err := r.client.GetNodeClient(contract.NodeID)
	if err != nil {
		return ZDB{}, errors.Wrapf(err, "failed to get client for node: %d", contract.NodeID)
	}

	cid, err := strconv.ParseUint(contract.ContractID, 10, 64)
	if err != nil {
		return ZDB{}, errors.Wrapf(err, "failed to parse contract Id: %s", contract.ContractID)
	}

	dl, err := cl.DeploymentGet(ctx, cid)
	if err != nil {
		return ZDB{}, errors.Wrapf(err, "failed to get deployment with contract Id: %s", contract.ContractID)
	}

	for _, workload := range dl.Workloads {
		if workload.Type == zos.ZDBType {
			zdb, err := workloads.NewZDBFromWorkload(&workload)
			if err != nil {
				return ZDB{}, errors.Wrapf(err, "failed to get zdb from workload: %s", workload.Name)
			}

			result := newZDBFromClientZDB(zdb)
			result.NodeID = contract.NodeID
			return result, nil
		}
	}

	return ZDB{}, fmt.Errorf("found zdb workloads in contract %d", contract.NodeID)
}

func newClientWorkloadFromZDB(zdb ZDB) workloads.ZDB {
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

func newZDBFromClientZDB(wl workloads.ZDB) ZDB {
	return ZDB{
		Name:        wl.Name,
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
