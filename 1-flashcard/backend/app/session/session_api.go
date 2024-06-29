package session

import "github.com/labstack/echo/v4"

func AttachApi(g *echo.Group) {
	g.POST("/sessions", createSession)
}

func createSession(c echo.Context) error {
	session := new(Session)
	if err := c.Bind(session); err != nil {
		return c.JSON(400, map[string]interface{}{
			"title":   "Bad Request",
			"status":  400,
			"message": "Please input the correct data",
		})
	}

	err := CreateSession(session)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"title":   "Internal Server Error",
			"status":  500,
			"message": "Unable to create session",
		})
	}

	return c.JSON(201, session)
}
