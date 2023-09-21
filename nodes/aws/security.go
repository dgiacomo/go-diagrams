package aws

import "github.com/blushft/go-diagrams/diagram"

type securityContainer struct {
	path string
	opts []diagram.NodeOption
}

var Security = &securityContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/security",
}

func (c *securityContainer) IdentityAndAccessManagementIamAccessAnalyzer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/identity-and-access-management-iam-access-analyzer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) IdentityAndAccessManagementIamAwsStsAlternate(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/identity-and-access-management-iam-aws-sts-alternate.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) IdentityAndAccessManagementIamAwsSts(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/identity-and-access-management-iam-aws-sts.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) IdentityAndAccessManagementIamMfaToken(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/identity-and-access-management-iam-mfa-token.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) AdConnector(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/ad-connector.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) IdentityAndAccessManagementIamLongTermSecurityCredential(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/identity-and-access-management-iam-long-term-security-credential.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Inspector(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/inspector.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Macie(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/macie.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) ManagedMicrosoftAd(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/managed-microsoft-ad.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) ResourceAccessManager(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/resource-access-manager.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) ShieldAdvanced(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/shield-advanced.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Waf(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/waf.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Artifact(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/artifact.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) CertificateAuthority(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/certificate-authority.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Cognito(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/cognito.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Detective(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/detective.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) DirectoryService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/directory-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) IdentityAndAccessManagementIamEncryptedData(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/identity-and-access-management-iam-encrypted-data.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) InspectorAgent(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/inspector-agent.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) WafFilteringRule(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/waf-filtering-rule.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) CertificateManager(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/certificate-manager.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Guardduty(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/guardduty.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Cloudhsm(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/cloudhsm.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) IdentityAndAccessManagementIamPermissions(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/identity-and-access-management-iam-permissions.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) SecurityHubFinding(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/security-hub-finding.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) IdentityAndAccessManagementIamDataEncryptionKey(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/identity-and-access-management-iam-data-encryption-key.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) IdentityAndAccessManagementIamRole(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/identity-and-access-management-iam-role.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) KeyManagementService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/key-management-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) SecretsManager(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/secrets-manager.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Shield(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/shield.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) CloudDirectory(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/cloud-directory.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) IdentityAndAccessManagementIamAddOn(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/identity-and-access-management-iam-add-on.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) IdentityAndAccessManagementIam(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/identity-and-access-management-iam.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) SecurityHub(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/security-hub.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) SingleSignOn(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/single-sign-on.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) FirewallManager(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/firewall-manager.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) IdentityAndAccessManagementIamTemporarySecurityCredential(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/identity-and-access-management-iam-temporary-security-credential.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) SecurityIdentityAndCompliance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/security-identity-and-compliance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) SimpleAd(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/security/simple-ad.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
