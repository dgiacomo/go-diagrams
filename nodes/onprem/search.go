package onprem

import "github.com/blushft/go-diagrams/diagram"

type searchContainer struct {
	path string
	opts []diagram.NodeOption
}

var Search = &searchContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/search",
}

func (c *searchContainer) Solr(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/search/solr.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
