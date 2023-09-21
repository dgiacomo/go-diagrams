package ibm

import "github.com/blushft/go-diagrams/diagram"

type storageContainer struct {
	path string
	opts []diagram.NodeOption
}

var Storage = &storageContainer{
	opts: diagram.OptionSet{diagram.Provider("ibm"), diagram.NodeShape("none")},
	path: "assets/ibm/storage",
}

func (c *storageContainer) BlockStorage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/storage/block-storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) ObjectStorage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/storage/object-storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
