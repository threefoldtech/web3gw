package tfgrid

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	proxyTypes "github.com/threefoldtech/tfgrid-sdk-go/grid-proxy/pkg/types"
)

var (
	Status  = "up"
	TrueVal = true
)

const (
	FarmerBotVersionAction  = "farmerbot.farmmanager.version"
	FarmerBotFindNodeAction = "farmerbot.nodemanager.findnode"
	FarmerBotRMBFunction    = "execute_job"
)

type FilterOptions struct {
	FarmID         uint32 `json:"farm_id"`
	PublicConfig   bool   `json:"public_config"`
	PublicIpsCount uint64 `json:"public_ips_count"`
	Dedicated      bool   `json:"dedicated"`
	MRU            uint64 `json:"mru"`
	HRU            uint64 `json:"hru"`
	SRU            uint64 `json:"sru"`
}

type FilterResult struct {
	FilterOption   FilterOptions `json:"filter_options"`
	AvailableNodes []uint32      `json:"available_nodes"`
}

type Reservations map[string]*PlannedReservation

type PlannedReservation struct {
	WorkloadName string
	NodeID       uint32
	FarmID       uint32
	MRU          uint64
	SRU          uint64
	HRU          uint64
	PublicIps    bool
}

type Args struct {
	RequiredHRU  *uint64  `json:"required_hru,omitempty"`
	RequiredSRU  *uint64  `json:"required_sru,omitempty"`
	RequiredCRU  *uint64  `json:"required_cru,omitempty"`
	RequiredMRU  *uint64  `json:"required_mru,omitempty"`
	NodeExclude  []uint32 `json:"node_exclude,omitempty"`
	Dedicated    *bool    `json:"dedicated,omitempty"`
	PublicConfig *bool    `json:"public_config,omitempty"`
	PublicIPs    *uint32  `json:"public_ips"`
	Certified    *bool    `json:"certified,omitempty"`
}

type Params struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type FarmerBotArgs struct {
	Args   []Args   `json:"args"`
	Params []Params `json:"params"`
}

type FarmerBotAction struct {
	Guid         string        `json:"guid"`
	TwinID       uint32        `json:"twinid"`
	Action       string        `json:"action"`
	Args         FarmerBotArgs `json:"args"`
	Result       FarmerBotArgs `json:"result"`
	State        string        `json:"state"`
	Start        uint64        `json:"start"`
	End          uint64        `json:"end"`
	GracePeriod  uint32        `json:"grace_period"`
	Error        string        `json:"error"`
	Timeout      uint32        `json:"timeout"`
	SourceTwinID uint32        `json:"src_twinid"`
	SourceAction string        `json:"src_action"`
	Dependencies []string      `json:"dependencies"`
}

func BuildGridProxyFilters(options FilterOptions, twinId uint64) proxyTypes.NodeFilter {
	proxyFilters := proxyTypes.NodeFilter{
		Status:       &Status,
		AvailableFor: &twinId,
	}

	if options.HRU != 0 {
		proxyFilters.FreeHRU = &options.HRU
	}

	if options.SRU != 0 {
		proxyFilters.FreeSRU = &options.SRU
	}

	if options.MRU != 0 {
		proxyFilters.FreeMRU = &options.MRU
	}

	if options.Dedicated {
		proxyFilters.Dedicated = &options.Dedicated
	}

	if options.PublicConfig {
		proxyFilters.IPv4 = &options.PublicConfig
		proxyFilters.Domain = &options.PublicConfig
	}

	if options.PublicIpsCount > 0 {
		proxyFilters.FreeIPs = &options.PublicIpsCount
	}

	if options.FarmID != 0 {
		proxyFilters.FarmIDs = []uint64{uint64(options.FarmID)}
	}

	return proxyFilters
}

func BuildFarmerBotParams(options FilterOptions) []Params {
	params := []Params{}
	if options.HRU != 0 {
		params = append(params, Params{Key: "required_hru", Value: options.HRU})
	}

	if options.SRU != 0 {
		params = append(params, Params{Key: "required_sru", Value: options.SRU})
	}

	if options.MRU != 0 {
		params = append(params, Params{Key: "required_mru", Value: options.MRU})
	}

	if options.Dedicated {
		params = append(params, Params{Key: "dedicated", Value: options.Dedicated})
	}

	if options.PublicConfig {
		params = append(params, Params{Key: "public_config", Value: options.PublicConfig})
	}

	if options.PublicIpsCount > 0 {
		params = append(params, Params{Key: "public_ips", Value: options.PublicIpsCount})
	}

	return params
}

func BuildFarmerBotAction(farmerTwinID uint32, sourceTwinID uint32, args []Args, params []Params, action string) FarmerBotAction {
	return FarmerBotAction{
		Guid:   uuid.NewString(),
		TwinID: farmerTwinID,
		Action: action,
		Args: FarmerBotArgs{
			Args:   args,
			Params: params,
		},
		Result: FarmerBotArgs{
			Args:   []Args{},
			Params: []Params{},
		},
		State:        "init",
		Start:        uint64(time.Now().Unix()),
		End:          0,
		GracePeriod:  0,
		Error:        "",
		Timeout:      6000,
		SourceTwinID: sourceTwinID,
		Dependencies: []string{},
	}
}

func (r *Client) GetFarmerTwinIDByFarmID(farmID uint32) (uint32, error) {
	farmid := uint64(farmID)
	farms, _, err := r.GridClient.FilterFarms(proxyTypes.FarmFilter{
		FarmID: &farmid,
	}, proxyTypes.Limit{
		Size: 1,
		Page: 1,
	})

	if err != nil || len(farms) == 0 {
		return 0, errors.Wrapf(err, "Couldn't get the FarmerTwinID for FarmID: %+v", farmID)
	}

	return uint32(farms[0].TwinID), nil
}

func GetFarmerBotResult(action FarmerBotAction, key string) (string, error) {
	if len(action.Result.Params) > 0 {
		for _, param := range action.Result.Params {
			if param.Key == key {
				return fmt.Sprint(param.Value), nil
			}
		}

	}

	return "", fmt.Errorf("Couldn't found a result for the same key: %s", key)
}

func (r *Client) FilterNodesWithFarmerBot(ctx context.Context, options FilterOptions) ([]uint32, error) {

	// construct farmerbot request
	params := BuildFarmerBotParams(options)

	// make farmerbot request
	farmerTwinID, err := r.GetFarmerTwinIDByFarmID(options.FarmID)
	if err != nil {
		return []uint32{}, errors.Wrapf(err, "Failed to get TwinID for FarmID %+v", options.FarmID)
	}

	sourceTwinID := r.TwinID

	data := BuildFarmerBotAction(farmerTwinID, sourceTwinID, []Args{}, params, FarmerBotFindNodeAction)

	var output FarmerBotAction

	err = r.GridClient.RMBCall(ctx, farmerTwinID, FarmerBotRMBFunction, data, &output)
	if err != nil {
		return []uint32{}, errors.Wrapf(err, "Failed calling farmerbot on farm %d", options.FarmID)
	}

	// build the result
	nodeIdStr, err := GetFarmerBotResult(output, "nodeid")
	if err != nil {
		return []uint32{}, err
	}

	nodeId, err := strconv.ParseUint(nodeIdStr, 10, 32)
	if err != nil {
		return []uint32{}, fmt.Errorf("can't parse node id")
	}

	return []uint32{uint32(nodeId)}, nil
}

func (r *Client) FilterNodesWithGridProxy(ctx context.Context, options FilterOptions) ([]uint32, error) {
	proxyFilters := BuildGridProxyFilters(options, uint64(r.TwinID))

	nodes, _, err := r.GridClient.FilterNodes(proxyFilters, proxyTypes.Limit{})
	if err != nil || len(nodes) == 0 {
		return []uint32{}, errors.Wrapf(err, "Couldn't find node for the provided filters: %+v", options)
	}

	nodesIDs := GetNodesIDs(nodes)

	return nodesIDs, nil
}

func GetNodesIDs(nodes []proxyTypes.Node) []uint32 {
	ids := []uint32{}

	for _, node := range nodes {
		ids = append(ids, uint32(node.NodeID))
	}

	return ids
}

func (r *Client) HasFarmerBot(ctx context.Context, farmID uint32) bool {
	args := []Args{}
	params := []Params{}

	farmerTwinID, err := r.GetFarmerTwinIDByFarmID(farmID)
	if err != nil {
		return false
	}

	sourceTwinID := r.TwinID

	data := BuildFarmerBotAction(farmerTwinID, sourceTwinID, args, params, FarmerBotVersionAction)

	var output FarmerBotAction

	err = r.GridClient.RMBCall(ctx, farmerTwinID, FarmerBotRMBFunction, data, &output)

	return err == nil
}

func (r *Client) checkNodeAvailability(nodeId uint32, workload PlannedReservation, reservedCapacity map[uint32]PlannedReservation) bool {
	// get node info
	node, err := r.GridClient.GetNode(nodeId)
	if err != nil {
		return false
	}

	// get free resources
	free := proxyTypes.Capacity{
		SRU: node.Capacity.Total.SRU - node.Capacity.Used.SRU,
		HRU: node.Capacity.Total.HRU - node.Capacity.Used.HRU,
		MRU: node.Capacity.Total.MRU - node.Capacity.Used.MRU,
	}

	// check if the free resource greater than the previous reserved capacity plus the current workload capacity
	if uint64(free.MRU) >= reservedCapacity[uint32(node.NodeID)].MRU+workload.MRU &&
		uint64(free.SRU) >= reservedCapacity[uint32(node.NodeID)].SRU+workload.SRU &&
		uint64(free.HRU) >= reservedCapacity[uint32(node.NodeID)].HRU+workload.HRU {
		return true
	}
	return false
}

func (r *Client) checkFarmAvailability(farmId uint64, workload PlannedReservation, reservedIps map[uint32]int) bool {
	// get farm info
	farms, _, err := r.GridClient.FilterFarms(proxyTypes.FarmFilter{
		FarmID: &farmId,
	}, proxyTypes.Limit{
		Size: 1,
		Page: 1,
	})

	if err != nil {
		return false
	}

	if len(farms[0].PublicIps) > reservedIps[workload.FarmID] {
		return true
	}

	return false
}

// Searching for node for each workload considering the reserved capacity by workloads in the same deployment.
// Assign the NodeID if found one or return it with NodeID: 0
func (c *Client) AssignNodes(ctx context.Context, workloads Reservations) error {

	reservedCapacity := make(map[uint32]PlannedReservation)
	reservedIps := make(map[uint32]int) // farmID -> numberOfPublicIps

	for _, workload := range workloads {
		if workload.NodeID == 0 {
			options := FilterOptions{
				FarmID: workload.FarmID,
				HRU:    workload.HRU,
				SRU:    workload.SRU,
				MRU:    workload.MRU,
			}

			if workload.PublicIps {
				options.PublicIpsCount = 1
			}

			var nodes []uint32
			var err error
			ctx2, cancel := context.WithTimeout(ctx, 3*time.Second)
			defer cancel()

			hasFarmerBot := c.HasFarmerBot(ctx2, options.FarmID)

			if !hasFarmerBot || options.FarmID == 0 {
				log.Info().Msg("Calling gridproxy")
				nodes, err = c.FilterNodesWithGridProxy(ctx, options)
				if err != nil {
					return errors.Wrap(err, "failed to filter nodes")
				}

				if len(nodes) == 0 {
					return fmt.Errorf("failed to find an elibile node satisfying specs")
				}

				if options.PublicIpsCount > 0 {
					farmIsValid := c.checkFarmAvailability(uint64(options.FarmID), *workload, reservedIps)
					if !farmIsValid {
						return fmt.Errorf("failed to find free public ips on farm %d", options.FarmID)
					}
					reservedIps[options.FarmID]++
				}

				selectedNodeId := uint32(0)
				for _, nodeId := range nodes {
					nodeIsValid := c.checkNodeAvailability(nodeId, *workload, reservedCapacity)
					if nodeIsValid {
						selectedNodeId = nodeId
						break
					}
				}

				if selectedNodeId == 0 {
					return fmt.Errorf("failed to find an elibile node satisfying specs")
				}

				workload.NodeID = selectedNodeId

				continue
			}

			log.Info().Msg("Calling farmerbot")
			nodes, err = c.FilterNodesWithFarmerBot(ctx, options)
			if err != nil {
				return errors.Wrap(err, "failed to filter nodes using farmerbot")
			}

			if len(nodes) == 0 {
				return fmt.Errorf("failed to find an elibile node satisfying specs")
			}

			workload.NodeID = nodes[0]
		}

		// update the reservedCapacity for the workload node with the resources. `either user provide NodeID or set by the filter`
		reservedCapacity[workload.NodeID] = PlannedReservation{
			MRU: reservedCapacity[workload.NodeID].MRU + workload.MRU,
			SRU: reservedCapacity[workload.NodeID].SRU + workload.SRU,
			HRU: reservedCapacity[workload.NodeID].HRU + workload.HRU,
		}

	}

	return nil
}
