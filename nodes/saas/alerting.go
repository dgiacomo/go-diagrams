package saas

import "github.com/blushft/go-diagrams/diagram"

type alertingContainer struct {
	path string
	opts []diagram.NodeOption
}

var Alerting = &alertingContainer{
	opts: diagram.OptionSet{diagram.Provider("saas"), diagram.NodeShape("none")},
	path: "assets/saas/alerting",
}

func (c *alertingContainer) Pushover(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/alerting/pushover.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *alertingContainer) Xmatters(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/alerting/xmatters.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *alertingContainer) Newrelic(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/alerting/newrelic.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *alertingContainer) Opsgenie(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/alerting/opsgenie.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *alertingContainer) Pagerduty(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/alerting/pagerduty.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
