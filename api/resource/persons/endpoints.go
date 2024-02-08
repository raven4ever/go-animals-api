package persons

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// get all persons endpoint
func GetPersons(c echo.Context) error {
	test := Persons{
		Person{ID: 1, FirstName: "John", LastName: "Doe", NickName: "JD", Age: 25},
		Person{ID: 2, FirstName: "Tracy", LastName: "Reagan", NickName: "TR", Age: 44},
		Person{ID: 3, FirstName: "Adrian", LastName: "Junior", NickName: "AJ", Age: 31},
	}

	return c.JSON(http.StatusOK, test)
}
