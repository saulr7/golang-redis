package handlers

import (
	"fmt"
	m "go-redis/models"
	s "go-redis/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddUserHandler(c echo.Context) error {

	u := new(m.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	us, err := s.AddUser(*u)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	return c.JSON(http.StatusCreated, us)

}

func GetUserHandler(c echo.Context) error {

	id := c.Param("id")

	u, err := s.GetUser(id)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	return c.JSON(http.StatusOK, u)
}

func DeleteUserHanlder(c echo.Context) error {

	id := c.Param("id")

	err := s.DeleteUser(id)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	return c.JSON(http.StatusOK, nil)

}
