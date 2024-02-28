package server

import (
	"animalz/api/resource/animals"
	"animalz/api/resource/foods"
	"animalz/api/resource/persons"

	"github.com/labstack/echo/v4"
)

type ApiServer struct {
	Address string
}

func NewApiServer(address string) *ApiServer {
	return &ApiServer{
		Address: address,
	}
}

func (s *ApiServer) Run() error {
	echoServer := echo.New()

	// create the API v1 group
	apiV1Group := echoServer.Group("/api/v1")

	// add the animals endpoints
	apiV1Group.GET("/animals", animals.GetAnimals)

	// add the persons endpoints
	apiV1Group.GET("/persons", persons.GetPersons)

	// add the foods endpoints
	apiV1Group.GET("/foods", foods.GetFoods)

	// start the server
	return echoServer.Start(s.Address)
}
