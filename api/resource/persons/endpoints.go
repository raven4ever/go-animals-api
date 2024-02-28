package persons

import (
	"animalz/db/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// get all persons endpoint
func GetPersons(c echo.Context) error {
	ctx := context.Background()
	session := db.NewSession(ctx, driver, neo4j.AccessModeWrite)
	defer session.Close(ctx)
	return c.JSON(http.StatusOK, test)
}
