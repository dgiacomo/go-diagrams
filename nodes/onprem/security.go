package onprem

import "github.com/blushft/go-diagrams/diagram"

type securityContainer struct {
	path string
	opts []diagram.NodeOption
}

var Security = &securityContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/security",
}

func (c *securityContainer) Bitwarden(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/security/bitwarden.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Trivy(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/security/trivy.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Vault(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/security/vault.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
