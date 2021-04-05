package main

import (
	"github/labstack/echo"
	"net/http"
)

func handleHome(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.Logger.Fatal(e.Start(":1323"))
}