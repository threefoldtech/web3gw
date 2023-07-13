package tfgrid

import (
	"context"

	"github.com/LeeSmet/go-jsonrpc"
	tfgridBase "github.com/threefoldtech/web3_proxy/server/clients/tfgrid"
	"github.com/threefoldtech/web3_proxy/server/pkg"
)

func (c *Client) DeployDiscourse(ctx context.Context, conState jsonrpc.State, discourse tfgridBase.Discourse) (tfgridBase.DiscourseResult, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.DiscourseResult{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.DeployDiscourse(ctx, discourse)
}

func (c *Client) GetDiscourse(ctx context.Context, conState jsonrpc.State, discourseName string) (tfgridBase.DiscourseResult, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.DiscourseResult{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.GetDiscourse(ctx, discourseName)
}

func (c *Client) DeleteDiscourse(ctx context.Context, conState jsonrpc.State, discourseName string) error {
	state := State(conState)
	if state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.cl.DeleteDiscourse(ctx, discourseName)
}

func (c *Client) DeployFunkwhale(ctx context.Context, conState jsonrpc.State, funkwhale tfgridBase.Funkwhale) (tfgridBase.FunkwhaleResult, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.FunkwhaleResult{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.Deployfunkwhale(ctx, funkwhale)
}

func (c *Client) GetFunkwhale(ctx context.Context, conState jsonrpc.State, funkwhaleName string) (tfgridBase.FunkwhaleResult, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.FunkwhaleResult{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.Getfunkwhale(ctx, funkwhaleName)
}

func (c *Client) DeleteFunkwhale(ctx context.Context, conState jsonrpc.State, funkwhaleName string) error {
	state := State(conState)
	if state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.cl.Deletefunkwhale(ctx, funkwhaleName)
}

func (c *Client) DeployPeertube(ctx context.Context, conState jsonrpc.State, peertube tfgridBase.Peertube) (tfgridBase.PeertubeResult, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.PeertubeResult{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.DeployPeertube(ctx, peertube)
}

func (c *Client) GetPeertube(ctx context.Context, conState jsonrpc.State, peertubeName string) (tfgridBase.PeertubeResult, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.PeertubeResult{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.GetPeertube(ctx, peertubeName)
}

func (c *Client) DeletePeertube(ctx context.Context, conState jsonrpc.State, peertubeName string) error {
	state := State(conState)
	if state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.cl.DeletePeertube(ctx, peertubeName)
}

func (c *Client) DeployPresearch(ctx context.Context, conState jsonrpc.State, presearch tfgridBase.Presearch) (tfgridBase.PresearchResult, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.PresearchResult{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.DeployPresearch(ctx, presearch)
}

func (c *Client) GetPresearch(ctx context.Context, conState jsonrpc.State, presearchName string) (tfgridBase.PresearchResult, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.PresearchResult{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.GetPresearch(ctx, presearchName)
}

func (c *Client) DeletePresearch(ctx context.Context, conState jsonrpc.State, presearchName string) error {
	state := State(conState)
	if state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.cl.DeletePresearch(ctx, presearchName)
}

func (c *Client) DeployTaiga(ctx context.Context, conState jsonrpc.State, taiga tfgridBase.Taiga) (tfgridBase.TaigaResult, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.TaigaResult{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.DeployTaiga(ctx, taiga)
}

func (c *Client) GetTaiga(ctx context.Context, conState jsonrpc.State, taigaName string) (tfgridBase.TaigaResult, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.TaigaResult{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.GetTaiga(ctx, taigaName)
}

func (c *Client) DeleteTaiga(ctx context.Context, conState jsonrpc.State, taigaName string) error {
	state := State(conState)
	if state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.cl.DeleteTaiga(ctx, taigaName)
}

func (c *Client) DeployVM(ctx context.Context, conState jsonrpc.State, vm tfgridBase.VM) (tfgridBase.VMResult, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.VMResult{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.DeployVM(ctx, vm)
}

func (c *Client) GetVM(ctx context.Context, conState jsonrpc.State, networkName string) (tfgridBase.VMResult, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.VMResult{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.GetVM(ctx, networkName)
}

func (c *Client) DeleteVM(ctx context.Context, conState jsonrpc.State, networkName string) error {
	state := State(conState)
	if state.cl == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.cl.DeleteVM(ctx, networkName)
}

func (c *Client) RemoveVM(ctx context.Context, conState jsonrpc.State, args tfgridBase.RemoveVM) (tfgridBase.VMResult, error) {
	state := State(conState)
	if state.cl == nil {
		return tfgridBase.VMResult{}, pkg.ErrClientNotConnected{}
	}

	return state.cl.RemoveVM(ctx, args)
}
