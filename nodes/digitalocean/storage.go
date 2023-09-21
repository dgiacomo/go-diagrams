package digitalocean

import "github.com/blushft/go-diagrams/diagram"

type storageContainer struct {
	path string
	opts []diagram.NodeOption
}

var Storage = &storageContainer{
	opts: diagram.OptionSet{diagram.Provider("digitalocean"), diagram.NodeShape("none")},
	path: "assets/digitalocean/storage",
}

func (c *storageContainer) VolumeSnapshot(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/storage/volume-snapshot.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Volume(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/storage/volume.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Folder(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/storage/folder.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Space(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/digitalocean/storage/space.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
