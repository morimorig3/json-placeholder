package main

import (
	"net/http"
	"strconv"

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
	e.GET("/todo/:count", createTodoHandler)

	// Error Handle
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Logger.Fatal(e.Start(":8080"))
}

func createTodoHandler(c echo.Context) error {
	p := c.Param("count")
	i, err := strconv.Atoi(p)

	// 数値以外のものが渡された場合エラー
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, "{}")
	}

	// 100件以上エラー
	if i > 100 {
		return c.JSON(http.StatusInternalServerError, "{}")
	}
	todoList := createTodoList(i)
	return c.JSON(http.StatusOK, todoList)
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

func createTodoList(cnt int) []Todo {
	todoList := make([]Todo, cnt)
	for i, _ := range todoList {
		newTodo := Todo{
			Id:          i,
			Task:        "タスク No." + strconv.Itoa(i),
			IsCompleted: false,
		}
		todoList[i] = newTodo
	}
	return todoList
}
