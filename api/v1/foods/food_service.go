package foods

import (
	"animalz/db"
	"animalz/db/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type FoodService struct {
	Database *db.Database
}

func NewFoodService(db *db.Database) *FoodService {
	return &FoodService{Database: db}
}

func (s *FoodService) RegisterRoutes(g *echo.Group) {
	g.GET("/foods", s.GetFoods)
}

func (s *FoodService) GetFoods(c echo.Context) error {
	session := s.Database.Driver.NewSession(s.Database.Context, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close(s.Database.Context)

	aRepo := repositories.NewFoodRepo(session, s.Database.Context)

	all, err := aRepo.GetFoods()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, all)
}
