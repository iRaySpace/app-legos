package card

import (
	"github.com/labstack/echo/v4"
)

func AttachApi(g *echo.Group) {
	g.POST("/cards", createCard)
	g.GET("/cards", getCards)
}

func createCard(c echo.Context) error {
	card := new(Card)
	if err := c.Bind(card); err != nil {
		return c.JSON(400, map[string]interface{}{
			"title":   "Bad Request",
			"status":  400,
			"message": "Please input the correct data",
		})
	}

	if err := c.Validate(card); err != nil {
		return err
	}

	err := CreateCard(card)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"title":   "Bad Request",
			"status":  400,
			"message": "Please input the correct data",
		})
	}

	return c.String(201, "Created")
}

func getCards(c echo.Context) error {
	return c.String(200, "Cards")
}
