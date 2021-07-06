package handlers

import (
	"github.com/labstack/echo/v4"
)

func Routes() *echo.Echo {
	e := echo.New()

	u := e.Group("user")

	u.GET("/:id", GetUserHandler)
	u.POST("/", AddUserHandler)

	return e
}
