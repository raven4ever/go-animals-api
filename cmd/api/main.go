package main

import (
	"animalz/api/server"
	"animalz/config"
	"animalz/db"
	"log"
)

func main() {
	log.Println("Starting the activity...")

	log.Println("Connecting to the Neo4j Database...")
	// create the database
	db, err := db.NewDatabase(config.EnvConfig.Neo4j)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	log.Println("Verifying DB connectivity...")
	// verify DB connection
	if err := db.VerifyConnectivity(); err != nil {
		log.Fatal(err)
	}

	// load demo data id applicable
	if config.EnvConfig.Neo4j.DemoData {
		log.Println("Loading demo data...")
		db.LoadDemoData()
	}

	// create the server
	srv := server.NewApiServer(config.EnvConfig.Server.Addr(), db)

	// Start server
	log.Println("Starting the server...")
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
