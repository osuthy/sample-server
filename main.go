package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", hello)
	e.Start("0.0.0.0:80")
}

func hello(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"aho": "ahoooooooo",
		"baka": "bakaなの？",
	})
}
