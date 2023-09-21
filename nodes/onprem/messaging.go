package onprem

import "github.com/blushft/go-diagrams/diagram"

type messagingContainer struct {
	path string
	opts []diagram.NodeOption
}

var Messaging = &messagingContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/messaging",
}

func (c *messagingContainer) Centrifugo(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/messaging/centrifugo.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
