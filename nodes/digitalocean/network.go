package digitalocean

import "github.com/blushft/go-diagrams/diagram"

type networkContainer struct {
	path string
	opts []diagram.NodeOption
}

var Network = &networkContainer{
	opts: diagram.OptionSet{diagram.Provider("digitalocean"), diagram.NodeShape("none")},
	path: "assets/digitalocean/network",
}

func (c *networkContainer) FloatingIp(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/network/floating-ip.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) InternetGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/network/internet-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ManagedVpn(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/network/managed-vpn.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Vpc(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/network/vpc.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) LoadBalancer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/network/load-balancer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Certificate(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/network/certificate.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) DomainRegistration(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/network/domain-registration.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Domain(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/network/domain.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Firewall(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/network/firewall.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
