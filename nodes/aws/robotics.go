package aws

import "github.com/blushft/go-diagrams/diagram"

type roboticsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Robotics = &roboticsContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/robotics",
}

func (c *roboticsContainer) RobomakerCloudExtensionRos(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/robotics/robomaker-cloud-extension-ros.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *roboticsContainer) RobomakerDevelopmentEnvironment(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/robotics/robomaker-development-environment.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *roboticsContainer) RobomakerFleetManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/robotics/robomaker-fleet-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *roboticsContainer) RobomakerSimulator(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/robotics/robomaker-simulator.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *roboticsContainer) Robomaker(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/robotics/robomaker.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *roboticsContainer) Robotics(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/robotics/robotics.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
