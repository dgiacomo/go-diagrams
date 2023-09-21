package onprem

import "github.com/blushft/go-diagrams/diagram"

type vcsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Vcs = &vcsContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/vcs",
}

func (c *vcsContainer) Gitea(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/vcs/gitea.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *vcsContainer) Github(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/vcs/github.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *vcsContainer) Gitlab(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/vcs/gitlab.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *vcsContainer) Svn(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/vcs/svn.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *vcsContainer) Git(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/vcs/git.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
