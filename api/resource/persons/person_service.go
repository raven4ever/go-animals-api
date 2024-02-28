package persons

import (
	"animalz/db"
	"animalz/db/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type PersonService struct {
	Database *db.Database
}

func NewPersonService(db *db.Database) *PersonService {
	return &PersonService{Database: db}
}

func (s *PersonService) RegisterRoutes(g *echo.Group) {
	g.GET("/persons", s.GetPersons)
}

func (s *PersonService) GetPersons(c echo.Context) error {
	session := s.Database.Driver.NewSession(s.Database.Context, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close(s.Database.Context)

	aRepo := repositories.NewPersonRepo(session, s.Database.Context)

	all, err := aRepo.GetPersons()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, all)
}
