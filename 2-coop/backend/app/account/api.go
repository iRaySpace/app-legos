package account

import (
	"github.com/labstack/echo/v4"
	"irayspace.com/coop/app"
)

func AttachApi(g *echo.Group) {
	g.POST("/accounts", createAccount)
	g.GET("/accounts/:number", getAccount)
}

func createAccount(c echo.Context) error {
	dto := new(CreateAccountDTO)
	if err := c.Bind(dto); err != nil {
		return app.NewBadRequestError(c, "Unable to parse data")
	}

	a, err := CreateAccount(dto)
	if err != nil {
		app.LogError(err)
		return app.NewInternalServerError(c, "Unable to create an account")
	}

	return c.JSON(201, a)
}

func getAccount(c echo.Context) error {
	number := c.Param("number")
	a, err := GetAccount(number)
	if err != nil {
		app.LogError(err)
		return app.NewNotFoundError(c, "Unable to get account")
	}
	return c.JSON(200, a)
}
