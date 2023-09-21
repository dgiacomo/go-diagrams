package elastic

import "github.com/blushft/go-diagrams/diagram"

type beatsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Beats = &beatsContainer{
	opts: diagram.OptionSet{diagram.Provider("elastic"), diagram.NodeShape("none")},
	path: "assets/elastic/beats",
}

func (c *beatsContainer) Functionbeat(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/beats/functionbeat.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *beatsContainer) Heartbeat(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/beats/heartbeat.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *beatsContainer) Metricbeat(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/beats/metricbeat.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *beatsContainer) Packetbeat(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/beats/packetbeat.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *beatsContainer) Winlogbeat(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/beats/winlogbeat.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *beatsContainer) Apm(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/beats/apm.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *beatsContainer) Auditbeat(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/beats/auditbeat.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *beatsContainer) Filebeat(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/beats/filebeat.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
