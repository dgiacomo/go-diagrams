package aws

import "github.com/blushft/go-diagrams/diagram"

type databaseContainer struct {
	path string
	opts []diagram.NodeOption
}

var Database = &databaseContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/database",
}

func (c *databaseContainer) Neptune(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/neptune.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DynamodbAttribute(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/dynamodb-attribute.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DynamodbAttributes(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/dynamodb-attributes.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DocumentdbMongodbCompatibility(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/documentdb-mongodb-compatibility.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DynamodbDax(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/dynamodb-dax.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) RdsMysqlInstance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/rds-mysql-instance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Timestream(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/timestream.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DatabaseMigrationService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/database-migration-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Database(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/database.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) ElasticacheCacheNode(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/elasticache-cache-node.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) ElasticacheForRedis(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/elasticache-for-redis.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) KeyspacesManagedApacheCassandraService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/keyspaces-managed-apache-cassandra-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) RdsMariadbInstance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/rds-mariadb-instance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Rds(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/rds.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) RedshiftDenseStorageNode(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/redshift-dense-storage-node.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Aurora(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/aurora.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DynamodbItems(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/dynamodb-items.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Redshift(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/redshift.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) RdsInstance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/rds-instance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) RdsOnVmware(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/rds-on-vmware.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) QuantumLedgerDatabaseQldb(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/quantum-ledger-database-qldb.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) RdsSqlServerInstance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/rds-sql-server-instance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) RedshiftDenseComputeNode(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/redshift-dense-compute-node.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DatabaseMigrationServiceDatabaseMigrationWorkflow(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/database-migration-service-database-migration-workflow.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) ElasticacheForMemcached(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/elasticache-for-memcached.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DynamodbTable(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/dynamodb-table.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Dynamodb(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/dynamodb.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Elasticache(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/elasticache.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) RdsOracleInstance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/rds-oracle-instance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) RdsPostgresqlInstance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/rds-postgresql-instance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) AuroraInstance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/aurora-instance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DynamodbGlobalSecondaryIndex(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/dynamodb-global-secondary-index.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DynamodbItem(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/database/dynamodb-item.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
