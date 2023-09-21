package onprem

import "github.com/blushft/go-diagrams/diagram"

type mlopsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Mlops = &mlopsContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/mlops",
}

func (c *mlopsContainer) Mlflow(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/mlops/mlflow.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlopsContainer) Polyaxon(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/mlops/polyaxon.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
