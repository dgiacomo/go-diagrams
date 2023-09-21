package digitalocean

import "github.com/blushft/go-diagrams/diagram"

type computeContainer struct {
	path string
	opts []diagram.NodeOption
}

var Compute = &computeContainer{
	opts: diagram.OptionSet{diagram.Provider("digitalocean"), diagram.NodeShape("none")},
	path: "assets/digitalocean/compute",
}

func (c *computeContainer) K8SNodePool(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/compute/k8s-node-pool.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) K8SNode(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/compute/k8s-node.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Containers(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/compute/containers.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Docker(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/compute/docker.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) DropletConnect(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/compute/droplet-connect.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) DropletSnapshot(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/compute/droplet-snapshot.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Droplet(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/compute/droplet.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) K8SCluster(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/compute/k8s-cluster.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
