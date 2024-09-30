package db

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

// Database (Cassandra) Session
var db_session *gocql.Session

func Init() *gocql.ClusterConfig {
	cluster := gocql.NewCluster(
		"cassandra",
	) // Use "Cassandra Container Name" if Connecting via Docker Compose Service Name
	cluster.Port = 9042
	cluster.Keyspace = "userdb" // Keyspace
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = 2 * time.Second

	// Return Cluster
	return cluster
}

func Connect() {
	cluster := Init()
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

func Close() {
	db_session.Close()
}
