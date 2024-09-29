package main

import (
	"fmt"
	"log"

	"go-user-crud-iris-cassandra/src/cmd/server"
	"go-user-crud-iris-cassandra/src/db"
)

func main() {
	// DataBase
	db.Init()
	db.Connect()
	db_session := db.GetSession()
	// defer db.Close()

	// Server
	server := server.NewServer(fmt.Sprintf(":%s", "9010"), db_session)

	// RUN Server
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
