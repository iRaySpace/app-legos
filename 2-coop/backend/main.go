package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"irayspace.com/coop/app/account"
)

func main() {
	e := echo.New()

	g := e.Group("/api/v1")
	account.AttachApi(g)

	log.Fatal(e.Start(":8000"))
}
