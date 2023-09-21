package ibm

import "github.com/blushft/go-diagrams/diagram"

type securityContainer struct {
	path string
	opts []diagram.NodeOption
}

var Security = &securityContainer{
	opts: diagram.OptionSet{diagram.Provider("ibm"), diagram.NodeShape("none")},
	path: "assets/ibm/security",
}

func (c *securityContainer) SecurityMonitoringIntelligence(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/security/security-monitoring-intelligence.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) SecurityServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/security/security-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) DataSecurity(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/security/data-security.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Firewall(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/security/firewall.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) GovernanceRiskCompliance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/security/governance-risk-compliance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) PhysicalSecurity(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/security/physical-security.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) InfrastructureSecurity(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/security/infrastructure-security.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Gateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/security/gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) IdentityAccessManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/security/identity-access-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) IdentityProvider(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/security/identity-provider.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) TrustendComputing(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/security/trustend-computing.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Vpn(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/security/vpn.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) ApiSecurity(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/security/api-security.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) BlockchainSecurityService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/security/blockchain-security-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
