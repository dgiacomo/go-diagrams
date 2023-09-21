package onprem

import "github.com/blushft/go-diagrams/diagram"

type queueContainer struct {
	path string
	opts []diagram.NodeOption
}

var Queue = &queueContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/queue",
}

func (c *queueContainer) Rabbitmq(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/queue/rabbitmq.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *queueContainer) Zeromq(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/queue/zeromq.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *queueContainer) Activemq(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/queue/activemq.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *queueContainer) Celery(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/queue/celery.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *queueContainer) Emqx(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/queue/emqx.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *queueContainer) Kafka(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/queue/kafka.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *queueContainer) Nats(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/queue/nats.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
