package server

import (
	"animalz/api/v1/animals"
	"animalz/api/v1/foods"
	"animalz/api/v1/persons"

	"animalz/db"

	"github.com/go-playground/validator/v10"
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

	// add the validator to the echo server
	echoServer.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}

	// start the server
	return echoServer.Start(s.Address)
}
