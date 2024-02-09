package foods

import (
	"animalz/db/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// get all foods endpoint
func GetFoods(c echo.Context) error {
	test := model.Foods{
		model.Food{ID: uuid.NewString(), Name: "Pizza", Description: "Cheese and Tomato", Qty: 5},
		model.Food{ID: uuid.NewString(), Name: "Burger", Description: "Cheese and Beef", Qty: 10},
		model.Food{ID: uuid.NewString(), Name: "Pasta", Description: "Cheese and Tomato", Qty: 8},
	}

	return c.JSON(http.StatusOK, test)
}
