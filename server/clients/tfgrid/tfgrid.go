package tfgrid

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/deployer"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/graphql"
	client "github.com/threefoldtech/tfgrid-sdk-go/grid-client/node"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/state"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/workloads"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-proxy/pkg/types"
)

type TFGridClient interface {
	DeployGWFQDN(ctx context.Context, gw *workloads.GatewayFQDNProxy) error
	DeployGWName(ctx context.Context, gw *workloads.GatewayNameProxy) error
	DeployK8sCluster(ctx context.Context, k8s *workloads.K8sCluster) error
	DeployNetwork(ctx context.Context, znet *workloads.ZNet) error
	DeployDeployment(ctx context.Context, d *workloads.Deployment) (uint64, error)
	CancelProject(ctx context.Context, projectName string) error
	GetProjectContracts(ctx context.Context, projectName string) (graphql.Contracts, error)
	GetNodeClient(nodeID uint32) (*client.NodeClient, error)
	RMBCall(ctx context.Context, twin uint32, fn string, data interface{}, result interface{}) error
	FilterNodes(filter types.NodeFilter, pagination types.Limit) (res []types.Node, totalCount int, err error)
	FilterFarms(filter types.FarmFilter, pagination types.Limit) (res []types.Farm, totalCount int, err error)
	GetNode(nodeID uint32) (res types.NodeWithNestedCapacity, err error)
	GetNodeDomain(ctx context.Context, nodeID uint32) (string, error)
	GetNodeFarm(nodeID uint32) (uint32, error)

	LoadNetwork(networkName string) (workloads.ZNet, error)
	LoadGatewayFQDN(nodeID uint32, modelName string) (workloads.GatewayFQDNProxy, error)
	LoadGatewayName(nodeID uint32, modelName string) (workloads.GatewayNameProxy, error)
	LoadK8s(nodeIDs []uint32, modelName string) (workloads.K8sCluster, error)
	LoadDeployment(nodeID uint32, modelName string) (workloads.Deployment, error)
	LoadZDB(nodeID uint32, modelName string) (workloads.ZDB, error)
	SetNodeDeploymentState(nodeContracts map[uint32][]uint64)
	SetNetworkState(nodeContract map[uint32]uint64)
	CancelDeployment(ctx context.Context, dl *workloads.Deployment) error
}

type tfgridClient struct {
	client *deployer.TFPluginClient
}

func NewTFGridClient(c *deployer.TFPluginClient) TFGridClient {
	return &tfgridClient{
		client: c,
	}
}

func (c *tfgridClient) DeployGWFQDN(ctx context.Context, gw *workloads.GatewayFQDNProxy) error {
	if err := c.client.GatewayFQDNDeployer.Deploy(ctx, gw); err != nil {
		return errors.Wrapf(err, "failed to deploy gateway fqdn")
	}

	return nil
}

func (c *tfgridClient) DeployGWName(ctx context.Context, gw *workloads.GatewayNameProxy) error {
	if err := c.client.GatewayNameDeployer.Deploy(ctx, gw); err != nil {
		return errors.Wrapf(err, "failed to deploy gateway %s", gw.Name)
	}

	return nil
}

func (c *tfgridClient) DeployK8sCluster(ctx context.Context, k8s *workloads.K8sCluster) error {
	if err := c.client.K8sDeployer.Deploy(ctx, k8s); err != nil {
		return errors.Wrapf(err, "Failed to deploy K8s Cluster")
	}

	return nil
}

func (c *tfgridClient) DeployNetwork(ctx context.Context, znet *workloads.ZNet) error {
	if err := c.client.NetworkDeployer.Deploy(ctx, znet); err != nil {
		return errors.Wrap(err, "failed to deploy network")
	}

	return nil
}
func (c *tfgridClient) DeployDeployment(ctx context.Context, d *workloads.Deployment) (uint64, error) {
	if err := c.client.DeploymentDeployer.Deploy(ctx, d); err != nil {
		return 0, errors.Wrap(err, "failed to deploy deployment")
	}

	return d.ContractID, nil
}
func (c *tfgridClient) CancelProject(ctx context.Context, projectName string) error {
	if err := c.client.CancelByProjectName(projectName); err != nil {
		return errors.Wrapf(err, "failed to cancel project %s", projectName)
	}

	return nil
}

func (c *tfgridClient) GetNodeClient(nodeID uint32) (*client.NodeClient, error) {
	nodeClient, err := c.client.NcPool.GetNodeClient(c.client.SubstrateConn, nodeID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	return nodeClient, nil
}
func (c *tfgridClient) GetProjectContracts(ctx context.Context, projectName string) (graphql.Contracts, error) {
	contracts, err := c.client.ContractsGetter.ListContractsOfProjectName(projectName)
	if err != nil {
		return graphql.Contracts{}, errors.Wrapf(err, "failed to get project (%s) contracts", projectName)
	}

	return contracts, nil
}

func (c *tfgridClient) RMBCall(ctx context.Context, twin uint32, fn string, data interface{}, result interface{}) error {
	return c.client.RMB.Call(ctx, twin, fn, data, result)
}
func (c *tfgridClient) FilterNodes(filter types.NodeFilter, pagination types.Limit) (res []types.Node, totalCount int, err error) {
	return c.client.GridProxyClient.Nodes(filter, pagination)
}
func (c *tfgridClient) FilterFarms(filter types.FarmFilter, pagination types.Limit) (res []types.Farm, totalCount int, err error) {
	return c.client.GridProxyClient.Farms(filter, pagination)
}
func (c *tfgridClient) GetNode(nodeID uint32) (res types.NodeWithNestedCapacity, err error) {
	return c.client.GridProxyClient.Node(nodeID)
}

func (c *tfgridClient) GetNodeFarm(nodeID uint32) (uint32, error) {
	node, err := c.client.GridProxyClient.Node(nodeID)
	if err != nil {
		return 0, err
	}

	return uint32(node.FarmID), nil
}

func (c *tfgridClient) GetNodeDomain(ctx context.Context, nodeID uint32) (string, error) {
	nodeClient, err := c.GetNodeClient(nodeID)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get node %d client", nodeID)
	}

	cfg, err := nodeClient.NetworkGetPublicConfig(ctx)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get node %d public config", nodeID)
	}

	return cfg.Domain, nil
}

func (c *tfgridClient) SetNodeDeploymentState(nodeContracts map[uint32][]uint64) {
	c.client.State.CurrentNodeDeployments = make(map[uint32]state.ContractIDs)
	for node, contracts := range nodeContracts {
		c.client.State.CurrentNodeDeployments[node] = contracts
	}
}

func (c *tfgridClient) SetNetworkState(nodeContract map[uint32]uint64) {
	c.client.State.CurrentNodeNetworks = make(map[uint32]state.ContractIDs)
	for node, contract := range nodeContract {
		c.client.State.CurrentNodeNetworks[node] = []uint64{contract}
	}
}

func (c *tfgridClient) LoadNetwork(networkName string) (workloads.ZNet, error) {
	return c.client.State.LoadNetworkFromGrid(networkName)
}

func (c *tfgridClient) LoadGatewayFQDN(nodeID uint32, modelName string) (workloads.GatewayFQDNProxy, error) {
	return c.client.State.LoadGatewayFQDNFromGrid(nodeID, modelName, modelName)
}

func (c *tfgridClient) LoadGatewayName(nodeID uint32, modelName string) (workloads.GatewayNameProxy, error) {
	return c.client.State.LoadGatewayNameFromGrid(nodeID, modelName, modelName)
}

func (c *tfgridClient) LoadK8s(nodeIDs []uint32, modelName string) (workloads.K8sCluster, error) {
	return c.client.State.LoadK8sFromGrid(nodeIDs, modelName)
}

func (c *tfgridClient) LoadDeployment(nodeID uint32, modelName string) (workloads.Deployment, error) {
	return c.client.State.LoadDeploymentFromGrid(nodeID, modelName)
}

func (c *tfgridClient) LoadZDB(nodeID uint32, modelName string) (workloads.ZDB, error) {
	return c.client.State.LoadZdbFromGrid(nodeID, modelName, modelName)
}

func (c *tfgridClient) CancelDeployment(ctx context.Context, dl *workloads.Deployment) error {
	return c.client.DeploymentDeployer.Cancel(ctx, dl)
}

func generateProjectName(modelName string) (projectName string) {
	return fmt.Sprintf("%s.web3proxy", modelName)
}
