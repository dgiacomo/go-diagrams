package ibm

import "github.com/blushft/go-diagrams/diagram"

type devopsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Devops = &devopsContainer{
	opts: diagram.OptionSet{diagram.Provider("ibm"), diagram.NodeShape("none")},
	path: "assets/ibm/devops",
}

func (c *devopsContainer) ContinuousDeploy(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/devops/continuous-deploy.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) ContinuousTesting(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/devops/continuous-testing.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) Devops(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/devops/devops.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) ArtifactManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/devops/artifact-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) BuildTest(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/devops/build-test.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) CodeEditor(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/devops/code-editor.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) CollaborativeDevelopment(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/devops/collaborative-development.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) ConfigurationManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/devops/configuration-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) Provision(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/devops/provision.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) ReleaseManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/devops/release-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
