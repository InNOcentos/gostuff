package controller

import (
	"myapp/user/dto"
	"myapp/user/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func createNewUser(c echo.Context) error {
	var user dto.CreateUserDto
	err := c.Bind(&user)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity,err)
	}

	return service.Create(c, user)
}

func Add(g *echo.Group) {
	g.POST("/new", createNewUser)
}