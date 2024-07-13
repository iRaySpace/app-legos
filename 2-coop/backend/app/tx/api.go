package tx

import (
	"github.com/labstack/echo/v4"
	"irayspace.com/coop/app"
)

func AttachApi(g *echo.Group) {
	g.POST("/txs", createTx)
	g.GET("/txs", getTxs)
}

func createTx(c echo.Context) error {
	t := new(CreateTxDTO)
	if err := c.Bind(t); err != nil {
		return app.NewBadRequestError(c, "Unable to parse data")
	}

	tx, err := CreateTx(t)
	if err != nil {
		app.LogError(err)
		return app.NewInternalServerError(c, "Unable to create a transaction")
	}

	return c.JSON(201, tx)
}

func getTxs(c echo.Context) error {
	data := []string{}
	return c.JSON(200, map[string]interface{}{"data": data})
}
