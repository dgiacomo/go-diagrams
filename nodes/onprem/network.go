package onprem

import "github.com/blushft/go-diagrams/diagram"

type networkContainer struct {
	path string
	opts []diagram.NodeOption
}

var Network = &networkContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/network",
}

func (c *networkContainer) Envoy(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/envoy.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Etcd(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/etcd.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Nginx(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/nginx.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Tomcat(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/tomcat.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Yarp(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/yarp.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Haproxy(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/haproxy.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Jbossas(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/jbossas.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Linkerd(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/linkerd.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Pomerium(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/pomerium.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Traefik(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/traefik.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Zookeeper(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/zookeeper.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Apache(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/apache.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Opnsense(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/opnsense.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Gunicorn(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/gunicorn.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Istio(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/istio.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) OpenServiceMesh(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/open-service-mesh.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Consul(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/consul.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Internet(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/internet.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Powerdns(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/powerdns.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Tyk(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/tyk.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Vyos(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/vyos.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Ocelot(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/ocelot.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Pfsense(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/pfsense.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Ambassador(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/ambassador.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Bind9(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/bind-9.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Caddy(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/caddy.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Glassfish(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/glassfish.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Jetty(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/jetty.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Kong(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/kong.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Wildfly(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/network/wildfly.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
