package ibm

import "github.com/blushft/go-diagrams/diagram"

type socialContainer struct {
	path string
	opts []diagram.NodeOption
}

var Social = &socialContainer{
	opts: diagram.OptionSet{diagram.Provider("ibm"), diagram.NodeShape("none")},
	path: "assets/ibm/social",
}

func (c *socialContainer) Messaging(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/social/messaging.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *socialContainer) Networking(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/social/networking.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *socialContainer) Communities(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/social/communities.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *socialContainer) FileSync(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/social/file-sync.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *socialContainer) LiveCollaboration(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/social/live-collaboration.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
