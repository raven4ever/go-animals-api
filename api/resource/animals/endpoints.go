package animals

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// get all animals endpoint
func GetAnimals(c echo.Context) error {
	test := []string{"dog", "cat", "bird", "fish", "hamster", "rabbit", "turtle", "snake", "lizard", "guinea pig", "chinchilla", "hedgehog", "ferret", "sugar glider", "parrot", "cockatoo", "macaw", "canary", "finch", "parakeet", "budgerigar", "lovebird", "african	gray", "amazon", "cockatiel", "conure", "pionus", "quaker", "ringneck", "rosella", "lorikeet", "lory", "cuckatoo", "parrakeet", "parraket", "paraket"}
	return c.JSON(http.StatusOK, test)
}
