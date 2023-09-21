package aws

import "github.com/blushft/go-diagrams/diagram"

type managementContainer struct {
	path string
	opts []diagram.NodeOption
}

var Management = &managementContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/management",
}

func (c *managementContainer) CloudformationStack(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/cloudformation-stack.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) Cloudtrail(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/cloudtrail.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) Codeguru(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/codeguru.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) OpsworksInstances(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/opsworks-instances.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) OpsworksLayers(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/opsworks-layers.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) SystemsManagerDocuments(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/systems-manager-documents.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) SystemsManagerPatchManager(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/systems-manager-patch-manager.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) Chatbot(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/chatbot.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) TrustedAdvisorChecklistCost(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/trusted-advisor-checklist-cost.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) OpsworksApps(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/opsworks-apps.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) OpsworksDeployments(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/opsworks-deployments.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) CloudwatchAlarm(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/cloudwatch-alarm.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) SystemsManagerParameterStore(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/systems-manager-parameter-store.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) ServiceCatalog(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/service-catalog.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) PersonalHealthDashboard(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/personal-health-dashboard.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) LicenseManager(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/license-manager.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) ManagementConsole(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/management-console.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) TrustedAdvisorChecklistSecurity(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/trusted-advisor-checklist-security.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) CloudwatchEventEventBased(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/cloudwatch-event-event-based.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) ManagedServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/managed-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) OpsworksMonitoring(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/opsworks-monitoring.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) SystemsManagerMaintenanceWindows(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/systems-manager-maintenance-windows.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) WellArchitectedTool(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/well-architected-tool.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) Cloudformation(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/cloudformation.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) CloudformationTemplate(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/cloudformation-template.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) CommandLineInterface(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/command-line-interface.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) Config(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/config.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) SystemsManagerInventory(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/systems-manager-inventory.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) TrustedAdvisorChecklist(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/trusted-advisor-checklist.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) AutoScaling(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/auto-scaling.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) Opsworks(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/opsworks.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) OpsworksPermissions(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/opsworks-permissions.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) Organizations(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/organizations.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) SystemsManagerAutomation(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/systems-manager-automation.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) SystemsManagerOpscenter(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/systems-manager-opscenter.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) CloudformationChangeSet(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/cloudformation-change-set.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) CloudwatchEventTimeBased(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/cloudwatch-event-time-based.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) ManagementAndGovernance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/management-and-governance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) OpsworksResources(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/opsworks-resources.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) OrganizationsAccount(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/organizations-account.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) OrganizationsOrganizationalUnit(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/organizations-organizational-unit.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) SystemsManagerRunCommand(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/systems-manager-run-command.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) SystemsManager(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/systems-manager.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) TrustedAdvisorChecklistFaultTolerant(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/trusted-advisor-checklist-fault-tolerant.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) OpsworksStack(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/opsworks-stack.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) TrustedAdvisorChecklistPerformance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/trusted-advisor-checklist-performance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) Cloudwatch(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/cloudwatch.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) SystemsManagerStateManager(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/systems-manager-state-manager.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) TrustedAdvisor(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/trusted-advisor.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) ControlTower(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/control-tower.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) CloudwatchRule(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/cloudwatch-rule.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
