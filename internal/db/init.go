package db

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3"
	"github.com/spf13/pflag"
)

var config = struct {
	DB       gocql.ClusterConfig
	Password gocql.PasswordAuthenticator
}{}

func init() {
	config.DB = *gocql.NewCluster()

	config.DB.Consistency = gocql.LocalQuorum

	pflag.StringArrayVar(&config.DB.Hosts, "hosts", []string{"brewscylla-1", "brewscylla-2", "brewscylla-3"}, "cluster nodes address list")
	pflag.DurationVar(&config.DB.Timeout, "timeout", 60*time.Second, "connection timeout")
	pflag.DurationVar(&config.DB.ConnectTimeout, "dial-timeout", 5*time.Second, "initial dial timeout")

	pflag.StringVar(&config.Password.Username, "username", "", "password based authentication username")
	pflag.StringVar(&config.Password.Password, "password", "", "password based authentication password")
}

func Config() gocql.ClusterConfig {
	var t = config.DB
	if config.Password.Username != "" {
		t.Authenticator = config.Password
	}
	return t
}

func Session() (*gocql.Session, error) {
	return gocql.NewSession(Config())
}

func Keyspace() (gocqlx.Session, error) {
	cfg := Config()
	cfg.Keyspace = keySpace
	cfg.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	return gocqlx.WrapSession(gocql.NewSession(cfg))
}
