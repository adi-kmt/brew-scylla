package db

import "log"

const (
	keySpace    = "brewscylla"
	keySpaceCQL = "CREATE KEYSPACE IF NOT EXISTS brewscylla WITH replication = { 'class': 'NetworkTopologyStrategy', 'replication_factor': '3' } AND durable_writes = TRUE AND TABLETS = {'enabled': false};"
)

func InitializeKeySpace() {
	ses, err := Session()
	if err != nil {
		log.Fatalln("session: ", err)
	}
	defer ses.Close()

	if err := ses.Query(keySpaceCQL).Exec(); err != nil {
		log.Fatalln("ensure keyspace exists: ", err)
	}
}

func printKeyspaceMetadata() {
	ses, err := Keyspace()
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
