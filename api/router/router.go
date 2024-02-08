package router

import (
	"animalz/api/resource/animals"
	"animalz/api/resource/persons"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	echoServer := echo.New()

	// create the API v1 group
	apiV1Group := echoServer.Group("/api/v1")

	// add the animals endpoints
	apiV1Group.GET("/animals", animals.GetAnimals)

	// add the persons endpoints
	apiV1Group.GET("/persons", persons.GetPersons)

	return echoServer
}
