package gcp

import "github.com/blushft/go-diagrams/diagram"

type operationsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Operations = &operationsContainer{
	opts: diagram.OptionSet{diagram.Provider("gcp"), diagram.NodeShape("none")},
	path: "assets/gcp/operations",
}

func (c *operationsContainer) Monitoring(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/operations/monitoring.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
