package onprem

import "github.com/blushft/go-diagrams/diagram"

type cdContainer struct {
	path string
	opts []diagram.NodeOption
}

var Cd = &cdContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/cd",
}

func (c *cdContainer) Spinnaker(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/cd/spinnaker.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *cdContainer) TektonCli(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/cd/tekton-cli.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *cdContainer) Tekton(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/cd/tekton.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
