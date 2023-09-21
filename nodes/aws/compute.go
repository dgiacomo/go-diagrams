package aws

import "github.com/blushft/go-diagrams/diagram"

type computeContainer struct {
	path string
	opts []diagram.NodeOption
}

var Compute = &computeContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/compute",
}

func (c *computeContainer) Compute(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/compute.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Ec2ElasticIpAddress(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/ec2-elastic-ip-address.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ThinkboxXmesh(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/thinkbox-xmesh.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Ec2Rescue(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/ec2-rescue.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ElasticContainerService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/elastic-container-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ThinkboxDraftRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/thinkbox-draft-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ThinkboxDraft(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/thinkbox-draft.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) LambdaRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/lambda-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Outposts(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/outposts.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ThinkboxStoke(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/thinkbox-stoke.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Batch(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/batch.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Ec2Rounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/ec2-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ElasticBeanstalkApplication(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/elastic-beanstalk-application.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) FargateRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/fargate-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) LambdaFunction(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/lambda-function.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) LocalZones(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/local-zones.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Ec2SpotInstance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/ec2-spot-instance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ElasticContainerServiceRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/elastic-container-service-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) LightsailRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/lightsail-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ThinkboxSequoia(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/thinkbox-sequoia.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ComputeRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/compute-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ThinkboxKrakatoa(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/thinkbox-krakatoa.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) AppRunner(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/app-runner.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ApplicationAutoScalingRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/application-auto-scaling-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Ec2(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/ec2.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Wavelength(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/wavelength.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ThinkboxDeadlineRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/thinkbox-deadline-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Ec2ContainerRegistryImage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/ec2-container-registry-image.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Ec2ContainerRegistryRegistry(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/ec2-container-registry-registry.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ElasticBeanstalkDeployment(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/elastic-beanstalk-deployment.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ElasticBeanstalkRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/elastic-beanstalk-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ThinkboxDeadline(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/thinkbox-deadline.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ThinkboxKrakatoaRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/thinkbox-krakatoa-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ServerlessApplicationRepositoryRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/serverless-application-repository-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Ec2ContainerRegistryRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/ec2-container-registry-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Ec2ImageBuilder(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/ec2-image-builder.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Lambda(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/lambda.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) OutpostsRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/outposts-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) VmwareCloudOnAwsRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/vmware-cloud-on-aws-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) BatchRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/batch-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Ec2AutoScaling(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/ec2-auto-scaling.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ElasticContainerServiceContainer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/elastic-container-service-container.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ServerlessApplicationRepository(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/serverless-application-repository.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ThinkboxFrost(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/thinkbox-frost.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Ec2Ami(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/ec2-ami.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ElasticKubernetesService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/elastic-kubernetes-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Fargate(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/fargate.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Lightsail(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/lightsail.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ThinkboxSequoiaRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/thinkbox-sequoia-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ThinkboxXmeshRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/thinkbox-xmesh-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ApplicationAutoScaling(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/application-auto-scaling.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Ec2Instances(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/ec2-instances.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ElasticKubernetesServiceRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/elastic-kubernetes-service-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ThinkboxFrostRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/thinkbox-frost-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ComputeOptimizer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/compute-optimizer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ElasticBeanstalk(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/elastic-beanstalk.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ThinkboxStokeRounded(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/thinkbox-stoke-rounded.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Ec2ContainerRegistry(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/ec2-container-registry.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Ec2Instance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/ec2-instance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ElasticContainerServiceService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/elastic-container-service-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) VmwareCloudOnAws(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/compute/vmware-cloud-on-aws.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
