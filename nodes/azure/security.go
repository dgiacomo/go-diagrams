package azure

import "github.com/blushft/go-diagrams/diagram"

type securityContainer struct {
	path string
	opts []diagram.NodeOption
}

var Security = &securityContainer{
	opts: diagram.OptionSet{diagram.Provider("azure"), diagram.NodeShape("none")},
	path: "assets/azure/security",
}

func (c *securityContainer) Defender(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/security/defender.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) ExtendedSecurityUpdates(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/security/extended-security-updates.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) KeyVaults(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/security/key-vaults.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) SecurityCenter(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/security/security-center.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Sentinel(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/security/sentinel.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) ApplicationSecurityGroups(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/security/application-security-groups.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) ConditionalAccess(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/security/conditional-access.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
