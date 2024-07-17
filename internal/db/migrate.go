package db

import (
	"context"
	"log"
	"os"

	"github.com/scylladb/gocqlx/v3/migrate"
)

func migrateKeyspace() {
	ses, err := Keyspace()
	if err != nil {
		log.Fatalln("session: ", err)
	}
	defer ses.Close()

	if err := migrate.FromFS(context.Background(), ses, os.DirFS("db/cql")); err != nil {
		log.Fatalln("migrate: ", err)
	}
}
