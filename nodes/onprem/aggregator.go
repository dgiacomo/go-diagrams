package onprem

import "github.com/blushft/go-diagrams/diagram"

type aggregatorContainer struct {
	path string
	opts []diagram.NodeOption
}

var Aggregator = &aggregatorContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/aggregator",
}

func (c *aggregatorContainer) Fluentd(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/aggregator/fluentd.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *aggregatorContainer) Vector(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/aggregator/vector.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
