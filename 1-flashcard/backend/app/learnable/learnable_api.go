package learnable

import (
	"github.com/labstack/echo/v4"
	"irayspace.com/flashcard/app/card"
)

func AttachApi(g *echo.Group) {
	g.POST("/learnables", createLearnable)
	g.GET("/learnables", getLearnables)
}

func createLearnable(c echo.Context) error {
	newCard := new(card.Card)
	if err := c.Bind(newCard); err != nil {
		return c.JSON(400, map[string]interface{}{
			"title":   "Bad Request",
			"status":  400,
			"message": "Please input the correct data",
		})
	}

	if err := c.Validate(newCard); err != nil {
		return err
	}

	// TODO: think if remove or not? if not, move to service
	err := card.CreateCard(newCard)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"title":   "Internal Server Error",
			"status":  500,
			"message": "Unable to create card",
		})
	}

	learnable, err := CreateLearnable(newCard)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"title":   "Internal Server Error",
			"status":  500,
			"message": "Unable to create learnable",
		})
	}

	return c.JSON(201, learnable)
}

func getLearnables(c echo.Context) error {
	learnables, err := GetLearnables()
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"title":   "Internal Server Error",
			"status":  500,
			"message": "Unable to get learnables",
		})
	}
	return c.JSON(200, map[string]interface{}{"data": learnables})
}
