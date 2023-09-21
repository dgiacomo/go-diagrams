package digitalocean

import "github.com/blushft/go-diagrams/diagram"

type databaseContainer struct {
	path string
	opts []diagram.NodeOption
}

var Database = &databaseContainer{
	opts: diagram.OptionSet{diagram.Provider("digitalocean"), diagram.NodeShape("none")},
	path: "assets/digitalocean/database",
}

func (c *databaseContainer) DbaasStandby(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/database/dbaas-standby.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DbaasPrimaryStandbyMore(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/database/dbaas-primary-standby-more.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DbaasPrimary(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/database/dbaas-primary.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DbaasReadOnly(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/database/dbaas-read-only.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
