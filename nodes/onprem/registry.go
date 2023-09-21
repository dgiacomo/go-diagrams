package onprem

import "github.com/blushft/go-diagrams/diagram"

type registryContainer struct {
	path string
	opts []diagram.NodeOption
}

var Registry = &registryContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/registry",
}

func (c *registryContainer) Harbor(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/registry/harbor.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *registryContainer) Jfrog(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/registry/jfrog.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
