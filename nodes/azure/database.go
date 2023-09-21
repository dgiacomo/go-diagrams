package azure

import "github.com/blushft/go-diagrams/diagram"

type databaseContainer struct {
	path string
	opts []diagram.NodeOption
}

var Database = &databaseContainer{
	opts: diagram.OptionSet{diagram.Provider("azure"), diagram.NodeShape("none")},
	path: "assets/azure/database",
}

func (c *databaseContainer) VirtualDatacenter(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/virtual-datacenter.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) SqlManagedInstances(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/sql-managed-instances.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) CosmosDb(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/cosmos-db.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DataExplorerClusters(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/data-explorer-clusters.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DatabaseForMysqlServers(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/database-for-mysql-servers.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) SynapseAnalytics(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/synapse-analytics.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) VirtualClusters(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/virtual-clusters.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) BlobStorage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/blob-storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) ElasticJobAgents(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/elastic-job-agents.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) SqlDatabases(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/sql-databases.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) SqlServerStretchDatabases(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/sql-server-stretch-databases.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) SqlVm(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/sql-vm.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Sql(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/sql.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) SsisLiftAndShiftIr(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/ssis-lift-and-shift-ir.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DataFactory(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/data-factory.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DataLake(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/data-lake.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DatabaseForMariadbServers(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/database-for-mariadb-servers.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DatabaseForPostgresqlServers(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/database-for-postgresql-servers.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) ElasticDatabasePools(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/elastic-database-pools.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) InstancePools(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/instance-pools.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) ManagedDatabases(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/managed-databases.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) SqlDatawarehouse(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/sql-datawarehouse.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) CacheForRedis(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/cache-for-redis.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) SqlServers(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/database/sql-servers.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
