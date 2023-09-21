package ibm

import "github.com/blushft/go-diagrams/diagram"

type infrastructureContainer struct {
	path string
	opts []diagram.NodeOption
}

var Infrastructure = &infrastructureContainer{
	opts: diagram.OptionSet{diagram.Provider("ibm"), diagram.NodeShape("none")},
	path: "assets/ibm/infrastructure",
}

func (c *infrastructureContainer) MonitoringLogging(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/infrastructure/monitoring-logging.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) Channels(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/infrastructure/channels.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) Diagnostics(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/infrastructure/diagnostics.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) EnterpriseMessaging(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/infrastructure/enterprise-messaging.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) EventFeed(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/infrastructure/event-feed.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) InterserviceCommunication(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/infrastructure/interservice-communication.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) LoadBalancingRouting(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/infrastructure/load-balancing-routing.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) MobileBackend(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/infrastructure/mobile-backend.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) Monitoring(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/infrastructure/monitoring.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) ServiceDiscoveryConfiguration(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/infrastructure/service-discovery-configuration.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) CloudMessaging(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/infrastructure/cloud-messaging.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) MicroservicesMesh(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/infrastructure/microservices-mesh.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) PeerServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/infrastructure/peer-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) InfrastructureServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/infrastructure/infrastructure-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) TransformationConnectivity(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/infrastructure/transformation-connectivity.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) Dashboard(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/infrastructure/dashboard.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) EdgeServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/infrastructure/edge-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) MobileProviderNetwork(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/infrastructure/mobile-provider-network.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
