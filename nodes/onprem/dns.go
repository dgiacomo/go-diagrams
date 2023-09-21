package onprem

import "github.com/blushft/go-diagrams/diagram"

type dnsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Dns = &dnsContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/dns",
}

func (c *dnsContainer) Coredns(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/dns/coredns.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *dnsContainer) Powerdns(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/dns/powerdns.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
