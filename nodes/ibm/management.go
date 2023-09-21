package ibm

import "github.com/blushft/go-diagrams/diagram"

type managementContainer struct {
	path string
	opts []diagram.NodeOption
}

var Management = &managementContainer{
	opts: diagram.OptionSet{diagram.Provider("ibm"), diagram.NodeShape("none")},
	path: "assets/ibm/management",
}

func (c *managementContainer) CloudManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/management/cloud-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) DataServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/management/data-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) ItServiceManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/management/it-service-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) PushNotifications(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/management/push-notifications.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) ContentManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/management/content-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) DeviceManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/management/device-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) InformationGovernance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/management/information-governance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) Management(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/management/management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) ProcessManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/management/process-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) ProviderCloudPortalService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/management/provider-cloud-portal-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) AlertNotification(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/management/alert-notification.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) ApiManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/management/api-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) ClusterManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/management/cluster-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) MonitoringMetrics(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/management/monitoring-metrics.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) ServiceManagementTools(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/management/service-management-tools.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
