package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/themadman159/go-tutor/config"
	"gorm.io/gorm"
)

type echoServer struct {
	app  *echo.Echo
	db   *gorm.DB
	conf *config.Config
}

var (
	once   sync.Once
	server *echoServer
)

func NewEchoServer(conf *config.Config, db *gorm.DB) *echoServer {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	once.Do(func() {
		server = &echoServer{
			app:  echoApp,
			db:   db,
			conf: conf,
		}
	})

	return server
}

func (s *echoServer) Start() {

	s.app.GET("/v1/health", s.healthCheck)

	s.httpListening()

}

func (s *echoServer) httpListening() {

	serverURL := fmt.Sprintf(":%d", s.conf.Server.Port)

	if err := s.app.Start(serverURL); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}

}

func (s *echoServer) healthCheck(c echo.Context) error {

	return c.String(http.StatusOK, "OK")

}
