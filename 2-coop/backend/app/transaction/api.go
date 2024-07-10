package transaction

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"irayspace.com/coop/app"
)

func AttachApi(g *echo.Group) {
	g.POST("/transactions", createTransaction)
	g.GET("/transactions", getTransactions)
}

func createTransaction(c echo.Context) error {
	t := new(Transaction)
	if err := c.Bind(t); err != nil {
		return app.NewBadRequestError(c, "Unable to parse data")
	}

	err := CreateTransaction(t)
	if err != nil {
		fmt.Println(err)
		return app.NewInternalServerError(c, "Unable to create a transaction")
	}

	return c.JSON(201, t)
}

func getTransactions(c echo.Context) error {
	data := []string{}
	return c.JSON(200, map[string]interface{}{"data": data})
}
