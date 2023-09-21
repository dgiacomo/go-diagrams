package onprem

import "github.com/blushft/go-diagrams/diagram"

type identityContainer struct {
	path string
	opts []diagram.NodeOption
}

var Identity = &identityContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/identity",
}

func (c *identityContainer) Dex(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/identity/dex.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
