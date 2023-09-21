package elastic

import "github.com/blushft/go-diagrams/diagram"

type agentContainer struct {
	path string
	opts []diagram.NodeOption
}

var Agent = &agentContainer{
	opts: diagram.OptionSet{diagram.Provider("elastic"), diagram.NodeShape("none")},
	path: "assets/elastic/agent",
}

func (c *agentContainer) Agent(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/agent/agent.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *agentContainer) Endpoint(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/agent/endpoint.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *agentContainer) Fleet(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/agent/fleet.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *agentContainer) Integrations(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/agent/integrations.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
