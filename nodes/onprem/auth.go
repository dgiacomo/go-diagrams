package onprem

import "github.com/blushft/go-diagrams/diagram"

type authContainer struct {
	path string
	opts []diagram.NodeOption
}

var Auth = &authContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/auth",
}

func (c *authContainer) Boundary(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/auth/boundary.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *authContainer) BuzzfeedSso(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/auth/buzzfeed-sso.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *authContainer) Oauth2Proxy(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/auth/oauth2-proxy.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
