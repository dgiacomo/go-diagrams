package onprem

import "github.com/blushft/go-diagrams/diagram"

type inmemoryContainer struct {
	path string
	opts []diagram.NodeOption
}

var Inmemory = &inmemoryContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/inmemory",
}

func (c *inmemoryContainer) Hazelcast(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/inmemory/hazelcast.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *inmemoryContainer) Memcached(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/inmemory/memcached.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *inmemoryContainer) Redis(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/inmemory/redis.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *inmemoryContainer) Aerospike(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/inmemory/aerospike.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
