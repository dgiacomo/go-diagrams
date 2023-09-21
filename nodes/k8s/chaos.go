package k8s

import "github.com/blushft/go-diagrams/diagram"

type chaosContainer struct {
	path string
	opts []diagram.NodeOption
}

var Chaos = &chaosContainer{
	opts: diagram.OptionSet{diagram.Provider("k8s"), diagram.NodeShape("none")},
	path: "assets/k8s/chaos",
}

func (c *chaosContainer) ChaosMesh(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/chaos/chaos-mesh.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *chaosContainer) LitmusChaos(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/chaos/litmus-chaos.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
