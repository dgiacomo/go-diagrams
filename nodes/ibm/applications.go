package ibm

import "github.com/blushft/go-diagrams/diagram"

type applicationsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Applications = &applicationsContainer{
	opts: diagram.OptionSet{diagram.Provider("ibm"), diagram.NodeShape("none")},
	path: "assets/ibm/applications",
}

func (c *applicationsContainer) ApplicationLogic(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/applications/application-logic.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationsContainer) Microservice(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/applications/microservice.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationsContainer) RuntimeServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/applications/runtime-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationsContainer) SpeechToText(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/applications/speech-to-text.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationsContainer) ApiPolyglotRuntimes(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/applications/api-polyglot-runtimes.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationsContainer) AppServer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/applications/app-server.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationsContainer) ActionableInsight(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/applications/actionable-insight.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationsContainer) ApiDeveloperPortal(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/applications/api-developer-portal.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationsContainer) EnterpriseApplications(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/applications/enterprise-applications.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationsContainer) IotApplication(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/applications/iot-application.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationsContainer) Ontology(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/applications/ontology.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationsContainer) OpenSourceTools(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/applications/open-source-tools.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationsContainer) SaasApplications(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/applications/saas-applications.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationsContainer) ServiceBroker(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/applications/service-broker.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationsContainer) Annotate(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/applications/annotate.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationsContainer) Index(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/applications/index.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationsContainer) MobileApp(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/applications/mobile-app.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationsContainer) VisualRecognition(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/applications/visual-recognition.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationsContainer) Visualization(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/applications/visualization.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
