package animals

import (
	"animalz/db/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// get all animals endpoint
func GetAnimals(c echo.Context) error {
	test := model.Animals{
		model.Animal{ID: uuid.NewString(), Name: "Fido", Species: "Dog", Age: 3},
		model.Animal{ID: uuid.NewString(), Name: "Whiskers", Species: "Cat", Age: 5},
		model.Animal{ID: uuid.NewString(), Name: "Fluffy", Species: "Rabbit", Age: 2},
	}

	return c.JSON(http.StatusOK, test)
}
