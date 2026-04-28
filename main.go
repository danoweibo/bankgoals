package main

import (
	"log"
)

func main() {
	dbase, err := PostgresDBInstance()
	if err != nil {
		log.Fatal("An unknown error occurred while connecting to the database: ", err)
	}

	if err = dbase.Init(); err != nil {
		log.Fatal("Could not initialize database: ", err)
	}

	server := APIServerInstance(":3000", dbase)
	server.Run()
}