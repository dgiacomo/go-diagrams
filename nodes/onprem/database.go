package onprem

import "github.com/blushft/go-diagrams/diagram"

type databaseContainer struct {
	path string
	opts []diagram.NodeOption
}

var Database = &databaseContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/database",
}

func (c *databaseContainer) Scylla(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/database/scylla.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Cockroachdb(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/database/cockroachdb.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Dgraph(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/database/dgraph.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Hbase(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/database/hbase.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Influxdb(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/database/influxdb.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Janusgraph(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/database/janusgraph.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Mongodb(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/database/mongodb.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Mssql(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/database/mssql.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Druid(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/database/druid.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Cassandra(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/database/cassandra.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Clickhouse(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/database/clickhouse.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Couchbase(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/database/couchbase.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Couchdb(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/database/couchdb.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Neo4J(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/database/neo4j.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Oracle(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/database/oracle.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Postgresql(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/database/postgresql.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Mariadb(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/database/mariadb.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Mysql(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/database/mysql.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
