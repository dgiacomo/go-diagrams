package ibm

import "github.com/blushft/go-diagrams/diagram"

type networkContainer struct {
	path string
	opts []diagram.NodeOption
}

var Network = &networkContainer{
	opts: diagram.OptionSet{diagram.Provider("ibm"), diagram.NodeShape("none")},
	path: "assets/ibm/network",
}

func (c *networkContainer) Region(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/region.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Router(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/router.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) TransitGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/transit-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) VpnGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/vpn-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Gateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) PublicGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/public-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) LoadBalancingRouting(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/load-balancing-routing.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Enterprise(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/enterprise.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Firewall(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/firewall.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) FloatingIp(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/floating-ip.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) InternetServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/internet-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) LoadBalancerListener(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/load-balancer-listener.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Rules(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/rules.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Bridge(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/bridge.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) DirectLink(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/direct-link.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) VpnPolicy(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/vpn-policy.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Subnet(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/subnet.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) VpnConnection(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/vpn-connection.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Vpc(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/vpc.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) LoadBalancerPool(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/load-balancer-pool.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) LoadBalancer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/network/load-balancer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
