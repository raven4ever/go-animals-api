package main

import (
	"animalz/api/resource/animal"
	"animalz/config"

	"github.com/labstack/echo/v4"
)

func main() {
	// create the config
	c := config.New()

	e := echo.New()

	gr := e.Group("/api/v1")
	gr.GET("/animals", animal.GetAnimals)

	// Start server
	e.Logger.Fatal(e.Start(c.Server.Addr()))
}
