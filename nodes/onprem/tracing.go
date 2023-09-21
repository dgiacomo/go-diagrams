package onprem

import "github.com/blushft/go-diagrams/diagram"

type tracingContainer struct {
	path string
	opts []diagram.NodeOption
}

var Tracing = &tracingContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/tracing",
}

func (c *tracingContainer) Jaeger(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/tracing/jaeger.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *tracingContainer) Tempo(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/tracing/tempo.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
