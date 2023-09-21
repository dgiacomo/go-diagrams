package onprem

import "github.com/blushft/go-diagrams/diagram"

type computeContainer struct {
	path string
	opts []diagram.NodeOption
}

var Compute = &computeContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/compute",
}

func (c *computeContainer) Nomad(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/compute/nomad.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Server(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/compute/server.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
