package persons

import (
	"animalz/db/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// get all persons endpoint
func GetPersons(c echo.Context) error {
	test := model.Persons{
		model.Person{ID: uuid.NewString(), FirstName: "John", LastName: "Doe", NickName: "JD", Age: 25},
		model.Person{ID: uuid.NewString(), FirstName: "Tracy", LastName: "Reagan", NickName: "TR", Age: 44},
		model.Person{ID: uuid.NewString(), FirstName: "Adrian", LastName: "Junior", NickName: "AJ", Age: 31},
	}

	return c.JSON(http.StatusOK, test)
}
