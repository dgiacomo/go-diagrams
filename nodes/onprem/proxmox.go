package onprem

import "github.com/blushft/go-diagrams/diagram"

type proxmoxContainer struct {
	path string
	opts []diagram.NodeOption
}

var Proxmox = &proxmoxContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/proxmox",
}

func (c *proxmoxContainer) Pve(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/proxmox/pve.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
