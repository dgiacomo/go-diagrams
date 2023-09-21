package programming

import "github.com/blushft/go-diagrams/diagram"

type runtimeContainer struct {
	path string
	opts []diagram.NodeOption
}

var Runtime = &runtimeContainer{
	opts: diagram.OptionSet{diagram.Provider("programming"), diagram.NodeShape("none")},
	path: "assets/programming/runtime",
}

func (c *runtimeContainer) Dapr(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/runtime/dapr.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
