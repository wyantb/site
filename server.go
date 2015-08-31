package main

import (
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

// Example handler
func hello(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func main() {
	e := echo.New()

	e.Use(mw.Logger())
	e.Use(mw.Logger())

	e.Get("/", hello)

	e.Run(":1323")
}
