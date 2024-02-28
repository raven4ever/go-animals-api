package main

import (
	"animalz/api/router"
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
	srv := router.New(config.EnvConfig.Server)

	// Start server
	srv.Logger.Fatal(srv.Start(config.EnvConfig.Server.Addr()))
}
