package onprem

import "github.com/blushft/go-diagrams/diagram"

type monitoringContainer struct {
	path string
	opts []diagram.NodeOption
}

var Monitoring = &monitoringContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/monitoring",
}

func (c *monitoringContainer) Datadog(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/monitoring/datadog.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Prometheus(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/monitoring/prometheus.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Sentry(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/monitoring/sentry.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Zabbix(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/monitoring/zabbix.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Nagios(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/monitoring/nagios.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Dynatrace(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/monitoring/dynatrace.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Humio(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/monitoring/humio.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Mimir(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/monitoring/mimir.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Newrelic(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/monitoring/newrelic.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) PrometheusOperator(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/monitoring/prometheus-operator.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Splunk(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/monitoring/splunk.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Cortex(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/monitoring/cortex.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Grafana(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/monitoring/grafana.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Thanos(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/monitoring/thanos.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
