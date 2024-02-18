package echos

import (
	"github.com/labstack/echo/v4"
	"github.com/uzdemirdl/ready2play/src/repository"
)

type HTTPServer struct {
	engine *echo.Echo
	repos  repository.Repository
}

func New() *HTTPServer {
	return &HTTPServer{
		engine: echo.New(),
	}
}

func (s HTTPServer) ECHO() *echo.Echo {
	return s.engine
}

func (s *HTTPServer) Init(repo repository.Repository) {
	s.repos = repo
	s.applyHandlers()
}

func (s *HTTPServer) applyHandlers() {
	s.engine.GET("/", getRoot(s.repos))
	s.engine.GET("/maps/:uuid", getMap(s.repos))
	s.engine.GET("/map_points/:map_uuid", getMapPoints(s.repos))
}
