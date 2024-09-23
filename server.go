package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Todo struct {
	Id          int    `json:"id"`
	Task        string `json:"task"`
	IsCompleted bool   `json:"isCompleted"`
}

func main() {
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// handler
	e.GET("/todo/:count", func(c echo.Context) error {
		count := c.Param("count")
		var content struct {
			Count string `json:"count"`
		}
		content.Count = count
		return c.JSON(http.StatusOK, content)
	})

	// Error Handle
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Logger.Fatal(e.Start(":8080"))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	if err := c.JSON(code, "{}"); err != nil {
		c.Logger().Error(err)
	}
}
