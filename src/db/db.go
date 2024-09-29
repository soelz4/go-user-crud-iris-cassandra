package db

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

// Database (Cassandra) Session
var db_session *gocql.Session

func Connect() {
	cluster := gocql.NewCluster(
		"127.0.0.1",
	) // Use "Cassandra Container Name" if Connecting via Docker Compose Service Name
	cluster.Port = 9042
	cluster.Keyspace = "system" // Built-in Keyspace
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = 2 * time.Second

	// Create a New Session
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Error Creating Cassandra Session: %v", err)
	} else {
		fmt.Println("Connected to Cassandra!")
	}

	db_session = session
}

func GetSession() *gocql.Session {
	return db_session
}
