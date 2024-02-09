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

	// create and parse the config
	c := config.New()

	log.Println("Connecting to the Neo4j Database...")
	// create the database driver
	ctx := context.Background()
	driver, err := db.New(c.Neo4j)
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
	if c.Neo4j.DemoData {
		log.Println("Loading demo data...")
		db.LoadDemoData(ctx, driver)
	}

	// create the server
	srv := router.New(c.Server)

	// Start server
	srv.Logger.Fatal(srv.Start(c.Server.Addr()))
}
