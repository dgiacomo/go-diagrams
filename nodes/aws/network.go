package aws

import "github.com/blushft/go-diagrams/diagram"

type networkContainer struct {
	path string
	opts []diagram.NodeOption
}

var Network = &networkContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/network",
}

func (c *networkContainer) DirectConnect(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/direct-connect.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ElasticLoadBalancing(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/elastic-load-balancing.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ElbApplicationLoadBalancer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/elb-application-load-balancer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Endpoint(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/endpoint.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Route53HostedZone(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/route-53-hosted-zone.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) TransitGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/transit-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) VpcElasticNetworkAdapter(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/vpc-elastic-network-adapter.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) CloudfrontEdgeLocation(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/cloudfront-edge-location.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) VpcFlowLogs(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/vpc-flow-logs.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) NatGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/nat-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) SiteToSiteVpn(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/site-to-site-vpn.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) VpcCustomerGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/vpc-customer-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) VpcElasticNetworkInterface(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/vpc-elastic-network-interface.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ApiGatewayEndpoint(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/api-gateway-endpoint.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Cloudfront(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/cloudfront.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) AppMesh(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/app-mesh.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) CloudfrontDownloadDistribution(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/cloudfront-download-distribution.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) PrivateSubnet(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/private-subnet.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Route53(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/route-53.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) VpcRouter(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/vpc-router.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ApiGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/api-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) CloudfrontStreamingDistribution(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/cloudfront-streaming-distribution.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ElbNetworkLoadBalancer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/elb-network-load-balancer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) VpnConnection(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/vpn-connection.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) VpnGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/vpn-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ClientVpn(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/client-vpn.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) GlobalAccelerator(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/global-accelerator.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) InternetGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/internet-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Nacl(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/nacl.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) NetworkingAndContentDelivery(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/networking-and-content-delivery.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) RouteTable(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/route-table.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) VpcPeering(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/vpc-peering.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) VpcTrafficMirroring(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/vpc-traffic-mirroring.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ElbClassicLoadBalancer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/elb-classic-load-balancer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Privatelink(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/privatelink.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) PublicSubnet(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/public-subnet.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Vpc(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/vpc.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) CloudMap(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/network/cloud-map.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
