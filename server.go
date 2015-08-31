package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	"github.com/wyantb/site/data"
)

func genericError(c *echo.Context) {
	c.String(http.StatusInternalServerError, "Error in request")
}
func hello(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}
func largeReported(c *echo.Context) error {
	data, err := data.Asset("LargeCommercialBuildingsReported.json")
	if err != nil {
		genericError(c)
		return err
	}
	return c.String(http.StatusOK, string(data))
}
func largeUnreported(c *echo.Context) error {
	data, err := data.Asset("LargeCommercialBuildingsUnreported.json")
	if err != nil {
		genericError(c)
		return err
	}
	return c.String(http.StatusOK, string(data))
}

func main() {
	e := echo.New()

	e.Use(mw.Logger())
	e.Use(mw.Logger())

	e.Get("/", hello)
	e.Get("/data/large-buildings-reported.json", largeReported)
	e.Get("/data/large-buildings-unreported.json", largeUnreported)

	port := 1323
	fmt.Printf("Running on port: %d\n", port)
	portspec := fmt.Sprintf(":%d", port)
	e.Run(portspec)
}
