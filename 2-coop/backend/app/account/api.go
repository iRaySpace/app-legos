package account

import (
	"github.com/labstack/echo/v4"
	"irayspace.com/coop/app"
)

func AttachApi(g *echo.Group) {
	g.POST("/accounts", createAccount)
	g.GET("/accounts", getAccounts)
}

func createAccount(c echo.Context) error {
	a := new(Account)
	if err := c.Bind(a); err != nil {
		return app.NewBadRequestError(c, "Unable to parse data")
	}

	err := CreateAccount(a)
	if err != nil {
		return app.NewInternalServerError(c, "Unable to create an account")
	}

	return c.JSON(201, a)
}

func getAccounts(c echo.Context) error {
	data := []string{}
	return c.JSON(200, map[string]interface{}{"data": data})
}
