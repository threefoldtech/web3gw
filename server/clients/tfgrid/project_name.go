package tfgrid

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

func (r *Runner) validateProjectName(ctx context.Context, projectName string) error {
	contracts, err := r.client.GetProjectContracts(ctx, projectName)
	if err != nil {
		return errors.Wrapf(err, "failed to retreive contracts with project name %s", projectName)
	}

	if len(contracts.NameContracts) > 0 || len(contracts.NodeContracts) > 0 || len(contracts.RentContracts) > 0 {
		return fmt.Errorf("invalid project name. project name (%s) is not unique", projectName)
	}

	return nil
}
