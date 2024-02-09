package foods

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// get all foods endpoint
func GetFoods(c echo.Context) error {
	test := Foods{
		Food{ID: 1, Name: "Pizza", Description: "Cheese and Tomato", Qty: 5},
		Food{ID: 2, Name: "Burger", Description: "Cheese and Beef", Qty: 10},
		Food{ID: 3, Name: "Pasta", Description: "Cheese and Tomato", Qty: 8},
	}

	return c.JSON(http.StatusOK, test)
}
