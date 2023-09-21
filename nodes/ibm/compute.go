package ibm

import "github.com/blushft/go-diagrams/diagram"

type computeContainer struct {
	path string
	opts []diagram.NodeOption
}

var Compute = &computeContainer{
	opts: diagram.OptionSet{diagram.Provider("ibm"), diagram.NodeShape("none")},
	path: "assets/ibm/compute",
}

func (c *computeContainer) Instance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/compute/instance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Key(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/compute/key.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) PowerInstance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/compute/power-instance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) BareMetalServer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/compute/bare-metal-server.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ImageService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/compute/image-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
