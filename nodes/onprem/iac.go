package onprem

import "github.com/blushft/go-diagrams/diagram"

type iacContainer struct {
	path string
	opts []diagram.NodeOption
}

var Iac = &iacContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/iac",
}

func (c *iacContainer) Ansible(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/iac/ansible.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iacContainer) Atlantis(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/iac/atlantis.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iacContainer) Awx(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/iac/awx.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iacContainer) Puppet(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/iac/puppet.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iacContainer) Terraform(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/iac/terraform.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
