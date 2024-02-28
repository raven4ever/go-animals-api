package animals

import (
	"animalz/db"
	"animalz/db/repositories"
	"net/http"

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
