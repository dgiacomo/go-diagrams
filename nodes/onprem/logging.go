package onprem

import "github.com/blushft/go-diagrams/diagram"

type loggingContainer struct {
	path string
	opts []diagram.NodeOption
}

var Logging = &loggingContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/logging",
}

func (c *loggingContainer) Rsyslog(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/logging/rsyslog.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *loggingContainer) SyslogNg(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/logging/syslog-ng.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *loggingContainer) Fluentbit(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/logging/fluentbit.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *loggingContainer) Graylog(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/logging/graylog.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *loggingContainer) Loki(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/logging/loki.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
