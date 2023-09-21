package onprem

import "github.com/blushft/go-diagrams/diagram"

type groupwareContainer struct {
	path string
	opts []diagram.NodeOption
}

var Groupware = &groupwareContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/groupware",
}

func (c *groupwareContainer) Nextcloud(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/groupware/nextcloud.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
