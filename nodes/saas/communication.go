package saas

import "github.com/blushft/go-diagrams/diagram"

type communicationContainer struct {
	path string
	opts []diagram.NodeOption
}

var Communication = &communicationContainer{
	opts: diagram.OptionSet{diagram.Provider("saas"), diagram.NodeShape("none")},
	path: "assets/saas/communication",
}

func (c *communicationContainer) Twilio(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/communication/twilio.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
