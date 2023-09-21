package aws

import "github.com/blushft/go-diagrams/diagram"

type generalContainer struct {
	path string
	opts []diagram.NodeOption
}

var General = &generalContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/general",
}

func (c *generalContainer) SslPadlock(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/ssl-padlock.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) TapeStorage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/tape-storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Users(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/users.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) GenericSdk(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/generic-sdk.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Forums(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/forums.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Multimedia(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/multimedia.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Disk(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/disk.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) GenericDatabase(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/generic-database.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) GenericSamlToken(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/generic-saml-token.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) InternetAlt1(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/internet-alt1.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) SamlToken(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/saml-token.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Sdk(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/sdk.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) InternetGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/internet-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Client(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/client.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) GenericFirewall(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/generic-firewall.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) MobileClient(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/mobile-client.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) TradicionalServer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/tradicional-server.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) User(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/user.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) GenericOfficeBuilding(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/generic-office-building.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) InternetAlt2(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/internet-alt2.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Marketplace(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/marketplace.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) OfficeBuilding(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/office-building.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) General(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/general.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Toolkit(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/toolkit.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) TraditionalServer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/general/traditional-server.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
