package programming

import "github.com/blushft/go-diagrams/diagram"

type flowchartContainer struct {
	path string
	opts []diagram.NodeOption
}

var Flowchart = &flowchartContainer{
	opts: diagram.OptionSet{diagram.Provider("programming"), diagram.NodeShape("none")},
	path: "assets/programming/flowchart",
}

func (c *flowchartContainer) LoopLimit(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/loop-limit.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) OffPageConnectorLeft(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/off-page-connector-left.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) OffPageConnectorRight(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/off-page-connector-right.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) Action(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/action.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) Collate(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/collate.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) Decision(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/decision.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) Inspection(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/inspection.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) InternalStorage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/internal-storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) Sort(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/sort.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) StartEnd(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/start-end.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) SummingJunction(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/summing-junction.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) Document(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/document.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) ManualInput(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/manual-input.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) ManualLoop(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/manual-loop.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) Merge(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/merge.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) MultipleDocuments(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/multiple-documents.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) Database(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/database.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) InputOutput(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/input-output.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) Or(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/or.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) Delay(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/delay.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) Display(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/display.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) PredefinedProcess(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/predefined-process.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) Preparation(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/preparation.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *flowchartContainer) StoredData(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/flowchart/stored-data.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
