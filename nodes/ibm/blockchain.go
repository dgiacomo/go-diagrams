package ibm

import "github.com/blushft/go-diagrams/diagram"

type blockchainContainer struct {
	path string
	opts []diagram.NodeOption
}

var Blockchain = &blockchainContainer{
	opts: diagram.OptionSet{diagram.Provider("ibm"), diagram.NodeShape("none")},
	path: "assets/ibm/blockchain",
}

func (c *blockchainContainer) Blockchain(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/blockchain.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) ClientApplication(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/client-application.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) Communication(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/communication.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) HyperledgerFabric(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/hyperledger-fabric.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) MessageBus(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/message-bus.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) SmartContract(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/smart-contract.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) Consensus(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/consensus.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) EventListener(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/event-listener.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) Event(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/event.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) ExistingEnterpriseSystems(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/existing-enterprise-systems.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) MembershipServicesProviderApi(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/membership-services-provider-api.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) Membership(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/membership.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) TransactionManager(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/transaction-manager.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) BlockchainDeveloper(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/blockchain-developer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) CertificateAuthority(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/certificate-authority.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) Ledger(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/ledger.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) Services(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) KeyManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/key-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) Node(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/node.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) Wallet(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/ibm/blockchain/wallet.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
