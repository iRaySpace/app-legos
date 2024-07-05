package app

import "github.com/labstack/echo/v4"

func NewBadRequestError(c echo.Context, m string) error {
	return c.JSON(400, map[string]interface{}{
		"title":   "Bad Request",
		"status":  400,
		"message": m,
	})
}

func NewInternalServerError(c echo.Context, m string) error {
	return c.JSON(500, map[string]interface{}{
		"title":   "Internal Server Error",
		"status":  500,
		"message": m,
	})
}
