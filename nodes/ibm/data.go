package ibm

import "github.com/blushft/go-diagrams/diagram"

type dataContainer struct {
	path string
	opts []diagram.NodeOption
}

var Data = &dataContainer{
	opts: diagram.OptionSet{diagram.Provider("ibm"), diagram.NodeShape("none")},
	path: "assets/ibm/data",
}

func (c *dataContainer) GroundTruth(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/data/ground-truth.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *dataContainer) DeviceIdentityService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/data/device-identity-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *dataContainer) EnterpriseUserDirectory(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/data/enterprise-user-directory.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *dataContainer) ConversationTrainedDeployed(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/data/conversation-trained-deployed.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *dataContainer) DataSources(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/data/data-sources.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *dataContainer) EnterpriseData(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/data/enterprise-data.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *dataContainer) FileRepository(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/data/file-repository.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *dataContainer) TmsDataInterface(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/data/tms-data-interface.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *dataContainer) Caches(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/data/caches.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *dataContainer) Cloud(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/data/cloud.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *dataContainer) Model(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/data/model.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *dataContainer) DataServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/data/data-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *dataContainer) DeviceRegistry(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/data/device-registry.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
