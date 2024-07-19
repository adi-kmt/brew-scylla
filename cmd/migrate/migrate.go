package main

import (
	"context"
	"log"
	"os"

	"github.com/adi-kmt/brew-scylla/internal/db"
	"github.com/scylladb/gocqlx/v3/migrate"
	"github.com/spf13/pflag"
)

const (
	keySpace    = "brew"
	keySpaceCQL = "CREATE KEYSPACE IF NOT EXISTS brew WITH replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 } AND durable_writes = TRUE AND TABLETS = {'enabled': false};"
)

var verbose = pflag.Bool("verbose", false, "output more info")

func main() {
	pflag.Parse()

	log.Println("Bootstrap database...")

	if *verbose {
		log.Printf("Configuration = %+v\n", db.Config())
	}

	migrateKeyspace()
	log.Println("Migrations settled!")
	printKeyspaceMetadata()
	log.Println("Now you're ready to use the application! Make sure to test the whole app: sensor, server and loadtest.")

}

func migrateKeyspace() {
	ses, err := db.Init()
	if err != nil {
		log.Fatalln("session: ", err)
	}
	defer ses.Close()

	if err := migrate.FromFS(context.Background(), ses, os.DirFS("internal/db/cql")); err != nil {
		log.Fatalln("migrate: ", err)
	}
}

func printKeyspaceMetadata() {
	ses, err := db.Init()
	if err != nil {
		log.Fatalln("session: ", err)
	}
	defer ses.Close()

	m, err := ses.KeyspaceMetadata(keySpace)
	if err != nil {
		log.Fatalln("keyspace metadata: ", err)
	}

	log.Printf("Keyspace metadata = %+v\n", *m)
}
