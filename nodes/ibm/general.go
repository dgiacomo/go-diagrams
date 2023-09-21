package ibm

import "github.com/blushft/go-diagrams/diagram"

type generalContainer struct {
	path string
	opts []diagram.NodeOption
}

var General = &generalContainer{
	opts: diagram.OptionSet{diagram.Provider("ibm"), diagram.NodeShape("none")},
	path: "assets/ibm/general",
}

func (c *generalContainer) CloudServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/cloud-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) IdentityProvider(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/identity-provider.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) MicroservicesMesh(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/microservices-mesh.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) TextToSpeech(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/text-to-speech.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) TransformationConnectivity(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/transformation-connectivity.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) CloudMessaging(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/cloud-messaging.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) GovernanceRiskCompliance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/governance-risk-compliance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) IotCloud(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/iot-cloud.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) ServiceDiscoveryConfiguration(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/service-discovery-configuration.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) InfrastructureSecurity(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/infrastructure-security.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) OfflineCapabilities(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/offline-capabilities.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Scalable(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/scalable.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Cloudant(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/cloudant.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) DataSecurity(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/data-security.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) IbmPublicCloud(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/ibm-public-cloud.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Internet(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/internet.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Enterprise(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/enterprise.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) ObjectStorage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/object-storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) CognitiveServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/cognitive-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) IdentityAccessManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/identity-access-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) MonitoringLogging(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/monitoring-logging.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Monitoring(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/monitoring.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) IbmContainers(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/ibm-containers.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) MicroservicesApplication(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/microservices-application.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Openwhisk(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/openwhisk.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) PeerCloud(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/peer-cloud.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) RetrieveRank(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/general/retrieve-rank.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
