package handlers

import (
	"github.com/labstack/echo/v4"
)

func Routes() *echo.Echo {
	e := echo.New()

	u := e.Group("user")
	t := e.Group("todo")

	u.GET("/:id", GetUserHandler)
	u.POST("/", AddUserHandler)
	u.DELETE("/:id", DeleteUserHanlder)

	t.GET("/", GetTodosHandler)
	t.POST("/", AddTodoHandler)

	return e
}
