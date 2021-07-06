package handlers

import (
	"fmt"
	m "go-redis/models"
	s "go-redis/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddTodoHandler(c echo.Context) error {

	t := new(m.Todo)
	if err := c.Bind(t); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	todo, err := s.AddTodo(*t)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	return c.JSON(http.StatusCreated, todo)

}

func GetTodosHandler(c echo.Context) error {

	todos, err := s.GetTodos()

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	return c.JSON(http.StatusCreated, todos)
}
