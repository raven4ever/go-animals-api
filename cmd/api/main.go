package main

import (
	"animalz/api/server"
	"animalz/config"
	"animalz/db"
	"context"
	"log"
)

func main() {
	log.Println("Starting the activity...")

	log.Println("Connecting to the Neo4j Database...")
	// create the database driver
	ctx := context.Background()
	driver, err := db.New(config.EnvConfig.Neo4j)
	if err != nil {
		log.Fatal(err)
	}

	defer driver.Close(ctx)

	log.Println("Verifying DB connectivity...")
	// verify DB connection
	if err := driver.VerifyConnectivity(ctx); err != nil {
		log.Fatal(err)
	}

	// load demo data id applicable
	if config.EnvConfig.Neo4j.DemoData {
		log.Println("Loading demo data...")
		db.LoadDemoData(driver)
	}

	// create the server
	srv := server.NewApiServer(config.EnvConfig.Server.Addr())

	// Start server
	log.Println("Starting the server...")
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
