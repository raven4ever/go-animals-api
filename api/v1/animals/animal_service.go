package animals

import (
	"animalz/db"
	"animalz/db/model"
	"animalz/db/repositories"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type AnimalService struct {
	Database *db.Database
}

func NewAnimalService(db *db.Database) *AnimalService {
	return &AnimalService{Database: db}
}

func (s *AnimalService) RegisterRoutes(g *echo.Group) {
	g.GET("/animals", s.GetAnimals)
	g.GET("/animals/:id", s.GetAnimal)
	g.POST("/animals", s.CreateAnimal)
	g.DELETE("/animals/:id", s.DeleteAnimal)
}

// get all animals endpoint
func (s *AnimalService) GetAnimals(c echo.Context) error {
	session := s.Database.Driver.NewSession(s.Database.Context, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close(s.Database.Context)

	aRepo := repositories.NewAnimalRepo(session, s.Database.Context)

	all, err := aRepo.GetAnimals()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, all)
}

func (s *AnimalService) GetAnimal(c echo.Context) error {
	session := s.Database.Driver.NewSession(s.Database.Context, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close(s.Database.Context)

	aRepo := repositories.NewAnimalRepo(session, s.Database.Context)

	animal := new(AnimalById)
	if err := c.Bind(animal); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(animal); err != nil {
		return err
	}

	found, err := aRepo.GetAnimal(animal.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if found.ID == "" {
		return c.JSON(http.StatusNotFound, errors.New("animal not found"))
	}

	return c.JSON(http.StatusOK, found)
}

func (s *AnimalService) CreateAnimal(c echo.Context) error {
	session := s.Database.Driver.NewSession(s.Database.Context, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close(s.Database.Context)

	aRepo := repositories.NewAnimalRepo(session, s.Database.Context)

	animal := new(NewAnimal)
	if err := c.Bind(animal); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(animal); err != nil {
		return err
	}

	created, err := aRepo.CreateAnimal(model.Animal{
		ID:      uuid.NewString(),
		Name:    animal.Name,
		Species: animal.Species,
		Age:     animal.Age,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, created)
}

func (s *AnimalService) DeleteAnimal(c echo.Context) error {
	session := s.Database.Driver.NewSession(s.Database.Context, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close(s.Database.Context)

	aRepo := repositories.NewAnimalRepo(session, s.Database.Context)

	animal := new(DeleteAnimalById)
	if err := c.Bind(animal); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(animal); err != nil {
		return err
	}

	_, err := aRepo.DeleteAnimal(animal.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}
