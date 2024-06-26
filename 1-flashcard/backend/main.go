package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"irayspace.com/flashcard/app/card"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		errs := err.(validator.ValidationErrors)
		violations := []string{}
		for _, e := range errs {
			violations = append(violations, e.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"title":      "Bad Request",
			"status":     400,
			"violations": violations,
		})
	}
	return nil
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Validator = &CustomValidator{validator: validator.New()}

	g := e.Group("/api/v1")
	card.AttachApi(g)

	log.Fatal(e.Start(":8000"))
}
