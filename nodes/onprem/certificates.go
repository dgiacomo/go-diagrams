package onprem

import "github.com/blushft/go-diagrams/diagram"

type certificatesContainer struct {
	path string
	opts []diagram.NodeOption
}

var Certificates = &certificatesContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/certificates",
}

func (c *certificatesContainer) CertManager(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/certificates/cert-manager.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *certificatesContainer) LetsEncrypt(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/certificates/lets-encrypt.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
