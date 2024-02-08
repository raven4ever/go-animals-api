package animal

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// get all animals endpoint
func GetAnimals(c echo.Context) error {
	test := Animals{
		Animal{ID: 1, Name: "Fido", Species: "Dog", Age: 3},
		Animal{ID: 2, Name: "Whiskers", Species: "Cat", Age: 5},
		Animal{ID: 3, Name: "Fluffy", Species: "Rabbit", Age: 2},
	}

	return c.JSON(http.StatusOK, test)
}
