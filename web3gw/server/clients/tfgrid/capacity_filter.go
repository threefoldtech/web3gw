package tfgrid

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

func (r *Client) FilterNodes(ctx context.Context, options NodeFilterOptions) ([]uint32, error) {
	var res []uint32
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

func (r *Client) FilterFarms(ctx context.Context, options FarmFilterOptions) ([]uint32, error) {
	var res []uint32
	var err error

	log.Info().Msg("Calling gridproxy")
	res, err = r.FilterFarmsWithGridProxy(ctx, options)

	return res, err
}

func (r *Client) FilterContracts(ctx context.Context, options ContractFilterOptions) ([]uint32, error) {
	var res []uint32
	var err error

	log.Info().Msg("Calling gridproxy")
	res, err = r.FilterContractsWithGridProxy(ctx, options)

	return res, err
}

func (r *Client) FilterTwins(ctx context.Context, options TwinFilterOptions) ([]uint32, error) {
	var res []uint32
	var err error

	log.Info().Msg("Calling gridproxy")
	res, err = r.FilterTwinsWithGridProxy(ctx, options)

	return res, err
}
