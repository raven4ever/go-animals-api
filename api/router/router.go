package router

import (
	"animalz/api/resource/animal"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	echoServer := echo.New()

	// create the API v1 group
	animals := echoServer.Group("/api/v1")

	// add the animals endpoints
	animals.GET("/animals", animal.GetAnimals)

	return echoServer
}
