package persons

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// get all persons endpoint
func GetPersons(c echo.Context) error {
	test := []string{"John", "Paul", "George", "Ringo", "Pete", "Stuart", "Brian", "Billy", "Tony", "George", "Eric", "Jack", "Keith", "Charlie", "Bill", "Mick", "Brian", "Ronnie", "Ian", "Roger", "John", "Paul", "George", "Ringo", "Pete", "Stuart", "Brian", "Billy", "Tony", "George", "Eric", "Jack", "Keith", "Charlie", "Bill", "Mick", "Brian", "Ronnie", "Ian", "Roger", "John", "Paul", "George", "Ringo", "Pete", "Stuart", "Brian", "Billy", "Tony", "George", "Eric", "Jack", "Keith", "Charlie", "Bill", "Mick", "Brian", "Ronnie", "Ian", "Roger", "John", "Paul", "George", "Ringo", "Pete", "Stuart", "Brian", "Billy", "Tony", "George", "Eric", "Jack", "Keith", "Charlie", "Bill", "Mick", "Brian", "Ronnie", "Ian", "Roger", "John", "Paul", "George", "Ringo", "Pete", "Stuart", "Brian", "Billy", "Tony", "George", "Eric", "Jack", "Keith", "Charlie", "Bill", "Mick", "Brian", "Ronnie", "Ian", "Roger", "John", "Paul", "George", "Ringo", "Pete", "Stuart", "Brian", "Billy", "Tony", "George", "Eric", "Jack", "Keith", "Charlie", "Bill", "Mick", "Brian", "Ronnie", "Ian", "Roger"}
	return c.JSON(http.StatusOK, test)
}
