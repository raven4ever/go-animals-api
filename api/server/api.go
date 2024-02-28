package server

import (
	"animalz/api/resource/animals"
	"animalz/api/resource/foods"
	"animalz/api/resource/persons"
	"animalz/db"

	"github.com/labstack/echo/v4"
)

type ApiServer struct {
	Address  string
	Database *db.Database
}

func NewApiServer(address string, db *db.Database) *ApiServer {
	return &ApiServer{
		Address:  address,
		Database: db,
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
