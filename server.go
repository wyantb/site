package main

import (
	"fmt"
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

	port := 1323
	fmt.Printf("Running on port: %d\n", port)
	portspec := fmt.Sprintf(":%d", port)
	e.Run(portspec)
}
