package onprem

import "github.com/blushft/go-diagrams/diagram"

type gitopsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Gitops = &gitopsContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/gitops",
}

func (c *gitopsContainer) Argocd(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/gitops/argocd.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *gitopsContainer) Flagger(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/gitops/flagger.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *gitopsContainer) Flux(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/gitops/flux.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
