package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/mstzn/modanisa_backend/routes"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Port        int
	ServerReady chan bool
}

func (s *Server) Start() {

	appPort := fmt.Sprintf(":%d", s.Port)

	e := echo.New()

	e.POST("/todos", routes.AddNewTodo)
	e.GET("/todos", routes.GetAllToDos)

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"pong": "ok",
		})
	})

	go func() {
		if err := e.Start(appPort); err != nil {
			logrus.Errorf(err.Error())
			logrus.Infof("shutting down the server")
		}
	}()

	if s.ServerReady != nil {
		s.ServerReady <- true
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logrus.Fatalf("failed to gracefully shutdown the server: %s", err)
	}
}
