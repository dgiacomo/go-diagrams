package ibm

import "github.com/blushft/go-diagrams/diagram"

type userContainer struct {
	path string
	opts []diagram.NodeOption
}

var User = &userContainer{
	opts: diagram.OptionSet{diagram.Provider("ibm"), diagram.NodeShape("none")},
	path: "assets/ibm/user",
}

func (c *userContainer) PhysicalEntity(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/user/physical-entity.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *userContainer) Sensor(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/user/sensor.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *userContainer) User(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/user/user.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *userContainer) Browser(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/user/browser.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *userContainer) Device(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/user/device.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *userContainer) IntegratedDigitalExperiences(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/user/integrated-digital-experiences.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
