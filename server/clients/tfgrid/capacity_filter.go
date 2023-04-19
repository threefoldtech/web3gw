package tfgrid

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

func (r *Client) FilterNodes(ctx context.Context, options FilterOptions) (FilterResult, error) {
	var res FilterResult
	var err error

	ctx2, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	hasFarmerBot := r.HasFarmerBot(ctx2, options.FarmID)

	if options.FarmID != 0 && hasFarmerBot {
		log.Info().Msg("Calling farmerbot")
		res, err = r.FilterNodesWithFarmerBot(ctx, options)
	} else {
		log.Info().Msg("Calling gridproxy")
		res, err = r.FilterNodesWithGridProxy(ctx, options)
	}

	return res, err
}
