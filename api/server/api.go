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

	// animals handlers
	animalsService := animals.NewAnimalService(s.Database)
	animalsService.RegisterRoutes(apiV1Group)

	// foods handlers
	foodsService := foods.NewFoodService(s.Database)
	foodsService.RegisterRoutes(apiV1Group)

	// persons handlers
	personsService := persons.NewPersonService(s.Database)
	personsService.RegisterRoutes(apiV1Group)

	// start the server
	return echoServer.Start(s.Address)
}
