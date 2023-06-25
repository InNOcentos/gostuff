package http

import (
	"fmt"
	"myapp/config"
	userController "myapp/user/controller"

	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
}

func InitServer() *Server {
	e := &Server{echo: echo.New()}

	e.AddRoutes()

	return e
}

func (s *Server) Bootstrap(cfg config.Config) {
	port := fmt.Sprintf(":%s", cfg.Server.Port)
	s.echo.Start(port)
}

func (s *Server) AddRoutes()  {
	g := s.echo.Group("/users")
	userController.Add(g)
}
