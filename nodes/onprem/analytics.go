package onprem

import "github.com/blushft/go-diagrams/diagram"

type analyticsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Analytics = &analyticsContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/analytics",
}

func (c *analyticsContainer) Databricks(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/analytics/databricks.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Superset(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/analytics/superset.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Norikra(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/analytics/norikra.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Singer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/analytics/singer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Spark(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/analytics/spark.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Beam(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/analytics/beam.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Dbt(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/analytics/dbt.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Hive(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/analytics/hive.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Metabase(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/analytics/metabase.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Dremio(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/analytics/dremio.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Flink(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/analytics/flink.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Storm(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/analytics/storm.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Hadoop(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/analytics/hadoop.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Powerbi(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/analytics/powerbi.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Presto(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/analytics/presto.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *analyticsContainer) Tableau(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/analytics/tableau.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
