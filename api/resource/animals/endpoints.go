package animals

import (
	"animalz/db/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// get all animals endpoint
func GetAnimals(c echo.Context) error {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	aRepo := repositories.NewAnimalRepo(session, ctx)
	return c.JSON(http.StatusOK, test)
}
