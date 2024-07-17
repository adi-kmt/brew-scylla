package db

const (
	keySpace    = "brewscylla"
	keySpaceCQL = "CREATE KEYSPACE IF NOT EXISTS brewscylla WITH replication = { 'class': 'NetworkTopologyStrategy', 'replication_factor': '3' } AND durable_writes = TRUE AND TABLETS = {'enabled': false};"
)
