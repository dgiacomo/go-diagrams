package aws

import "github.com/blushft/go-diagrams/diagram"

type integrationContainer struct {
	path string
	opts []diagram.NodeOption
}

var Integration = &integrationContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/integration",
}

func (c *integrationContainer) StepFunctions(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/step-functions.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) ApplicationIntegration(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/application-integration.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) ConsoleMobileApplication(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/console-mobile-application.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) EventbridgeSaasPartnerEventBusResource(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/eventbridge-saas-partner-event-bus-resource.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) Eventbridge(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/eventbridge.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) SimpleNotificationServiceSnsEmailNotification(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/simple-notification-service-sns-email-notification.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) SimpleNotificationServiceSnsHttpNotification(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/simple-notification-service-sns-http-notification.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) Appsync(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/appsync.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) SimpleQueueServiceSqs(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/simple-queue-service-sqs.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) EventResource(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/event-resource.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) EventbridgeCustomEventBusResource(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/eventbridge-custom-event-bus-resource.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) SimpleNotificationServiceSnsTopic(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/simple-notification-service-sns-topic.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) SimpleQueueServiceSqsQueue(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/simple-queue-service-sqs-queue.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) EventbridgeDefaultEventBusResource(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/eventbridge-default-event-bus-resource.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) ExpressWorkflows(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/express-workflows.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) Mq(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/mq.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) SimpleNotificationServiceSns(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/simple-notification-service-sns.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) SimpleQueueServiceSqsMessage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/simple-queue-service-sqs-message.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
