package main

import (
	"animalz/api/router"
	"animalz/config"
)

func main() {
	// create the config
	c := config.New()

	srv := router.New()

	// Start server
	srv.Logger.Fatal(srv.Start(c.Server.Addr()))
}
