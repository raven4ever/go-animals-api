package foods

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// get all foods endpoint
func GetFoods(c echo.Context) error {
	test := []string{"apple", "banana", "orange", "pear", "grape", "strawberry", "blueberry", "raspberry", "blackberry", "kiwi", "mango", "pineapple", "watermelon", "cantaloupe", "honeydew", "peach", "plum", "cherry", "apricot", "nectarine", "pomegranate", "fig", "date", "coconut", "papaya", "lychee", "guava", "passion fruit", "dragon fruit", "star fruit", "persimmon", "kiwano", "breadfruit", "jackfruit", "durian", "rambutan", "longan", "mangosteen", "salak", "soursop", "cherimoya", "custard apple", "sugar apple", "atemoia", "atemoya"}
	return c.JSON(http.StatusOK, test)
}
