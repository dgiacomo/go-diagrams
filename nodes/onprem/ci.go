package onprem

import "github.com/blushft/go-diagrams/diagram"

type ciContainer struct {
	path string
	opts []diagram.NodeOption
}

var Ci = &ciContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/ci",
}

func (c *ciContainer) Concourseci(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/ci/concourseci.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *ciContainer) Droneci(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/ci/droneci.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *ciContainer) GithubActions(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/ci/github-actions.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *ciContainer) Gitlabci(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/ci/gitlabci.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *ciContainer) Jenkins(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/ci/jenkins.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *ciContainer) Teamcity(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/ci/teamcity.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *ciContainer) Travisci(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/ci/travisci.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *ciContainer) Circleci(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/ci/circleci.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *ciContainer) Zuulci(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/ci/zuulci.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
