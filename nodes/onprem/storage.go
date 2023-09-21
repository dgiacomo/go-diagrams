package onprem

import "github.com/blushft/go-diagrams/diagram"

type storageContainer struct {
	path string
	opts []diagram.NodeOption
}

var Storage = &storageContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/storage",
}

func (c *storageContainer) CephOsd(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/storage/ceph-osd.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Ceph(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/storage/ceph.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Glusterfs(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/storage/glusterfs.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Portworx(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/storage/portworx.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
