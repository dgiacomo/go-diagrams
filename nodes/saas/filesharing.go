package saas

import "github.com/blushft/go-diagrams/diagram"

type filesharingContainer struct {
	path string
	opts []diagram.NodeOption
}

var Filesharing = &filesharingContainer{
	opts: diagram.OptionSet{diagram.Provider("saas"), diagram.NodeShape("none")},
	path: "assets/saas/filesharing",
}

func (c *filesharingContainer) Nextcloud(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/filesharing/nextcloud.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
