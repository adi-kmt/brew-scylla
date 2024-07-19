package db

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3"
)

var cluster *gocql.ClusterConfig

func Init() (gocqlx.Session, error) {
	cluster = createCluster(gocql.LocalQuorum, "brew", "127.0.0.1")
	return gocqlx.WrapSession(cluster.CreateSession())
}

func Config() gocql.ClusterConfig {
	return *cluster
}

func createCluster(consistency gocql.Consistency, keyspace string, hosts ...string) *gocql.ClusterConfig {
	retryPolicy := &gocql.ExponentialBackoffRetryPolicy{
		Min:        time.Second,
		Max:        30 * time.Second,
		NumRetries: 5,
	}
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = keyspace
	cluster.Timeout = 2 * time.Minute
	cluster.RetryPolicy = retryPolicy
	cluster.Consistency = consistency
	cluster.ProtoVersion = 4
	cluster.Port = 9042
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "cassandra",
		Password: "cassandra",
	}
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	return cluster
}
