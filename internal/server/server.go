package server

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Start(port string, app *fiber.App) error {

	fiberServer := app.Server()

	fiberServer.ReadTimeout = 10 * time.Second
	fiberServer.WriteTimeout = 10 * time.Second
	fiberServer.Handler = app.Handler()

	return fiberServer.ListenAndServe(port)
}
