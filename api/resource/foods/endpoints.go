package foods

import (
	"animalz/db/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// get all foods endpoint
func GetFoods(c echo.Context) error {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	fRepo := repositories.NewFoodRepo(session, ctx)
	return c.JSON(http.StatusOK, test)
}
