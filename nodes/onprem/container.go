package onprem

import "github.com/blushft/go-diagrams/diagram"

type containerContainer struct {
	path string
	opts []diagram.NodeOption
}

var Container = &containerContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/container",
}

func (c *containerContainer) K3S(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/container/k3s.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *containerContainer) Lxc(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/container/lxc.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *containerContainer) Rkt(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/container/rkt.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *containerContainer) Containerd(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/container/containerd.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *containerContainer) Crio(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/container/crio.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *containerContainer) Docker(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/container/docker.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *containerContainer) Firecracker(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/container/firecracker.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *containerContainer) Gvisor(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/container/gvisor.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
