package tfgrid

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

func (c *Client) validateProjectName(ctx context.Context, modelName string) error {
	projectName := generateProjectName(modelName)

	if _, ok := c.Projects[projectName]; ok {
		return fmt.Errorf("invalid project name. project %s is not unique", projectName)
	}

	contracts, err := c.GridClient.GetProjectContracts(ctx, projectName)
	if err != nil {
		return errors.Wrapf(err, "failed to retreive contracts with project name %s", projectName)
	}

	if len(contracts.NameContracts) > 0 || len(contracts.NodeContracts) > 0 || len(contracts.RentContracts) > 0 {
		return fmt.Errorf("invalid project name. project name (%s) is not unique", projectName)
	}

	return nil
}

func (c *Client) cancelModel(ctx context.Context, modelName string) error {
	projectName := generateProjectName(modelName)
	if st, ok := c.Projects[projectName]; ok {
		// project contracts are stored locally
		for _, contractID := range st.nameContracts {
			if err := c.GridClient.CancelContract(ctx, contractID); err != nil {
				return err
			}
		}

		for _, contractIDs := range st.nodeContracts {
			for _, cid := range contractIDs {
				if err := c.GridClient.CancelContract(ctx, cid); err != nil {
					return err
				}
			}
		}

		delete(c.Projects, projectName)

		return nil
	}

	// project contracts are not stored locally, fetch from graphql, then cancel

	if err := c.GridClient.CancelProject(ctx, projectName); err != nil {
		return err
	}

	return nil
}
