package service

import (
	"log"
	"myapp/user/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context, u dto.CreateUserDto) error {
	log.Println(u)

	return c.String(http.StatusOK, "test")
}