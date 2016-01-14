package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	"github.com/wyantb/site/data"
)

func genericError(c *echo.Context) {
	c.String(http.StatusInternalServerError, "Error in request")
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

	e.Get("/data/large-buildings-reported.json", largeReported)
	e.Get("/data/large-buildings-unreported.json", largeUnreported)

	e.Static("/js/", "assets/js")
	e.Static("/css/", "assets/css")

	e.ServeFile("/about", "assets/about.html")
	e.ServeFile("/", "assets/index.html")
	e.ServeFile("/index", "assets/index.html")

	e.ServeFile("/chrome-crash", "assets/ChromeCrash.html")
	e.ServeFile("/25mb", "25mb")
	e.ServeFile("/100mb", "100mb")
	e.ServeFile("/250mb", "250mb")
	e.ServeFile("/500mb", "500mb")
	e.ServeFile("/1000mb", "1000mb")
	e.ServeFile("/2000mb", "2000mb")

	args := os.Args[1:]
	port := 80
	if len(args) >= 1 {
		parsed, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		port = parsed
	}

	fmt.Printf("Running on port: %d\n", port)
	portspec := fmt.Sprintf(":%d", port)
	e.Run(portspec)
}
