package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// handler
	e.GET("/todo/:count", createTodoHandler)

	// Error Handle
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Logger.Fatal(e.Start(":10000"))
}

func createTodoHandler(c echo.Context) error {
	p := c.Param("count")
	i, err := strconv.Atoi(p)
	todoList := TodoList{
		TodoList: []Todo{},
	}

	// 数値以外のものが渡された場合エラー
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, todoList.TodoList)
	}

	// 100件以上エラー
	if i > 100 {
		return c.JSON(http.StatusInternalServerError, todoList.TodoList)
	}
	todoList.List(i)
	return c.JSON(http.StatusOK, todoList.TodoList)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	if err := c.JSON(code, TodoList{
		TodoList: []Todo{},
	}.TodoList); err != nil {
		c.Logger().Error(err)
	}
}
