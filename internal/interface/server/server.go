package server

import (
	"fmt"
	"net/http"

	"github.com/djoleary/go_ddd_template/internal/infrastructure/environ"
	"github.com/labstack/echo"
)

type server struct {
	env       environ.Getenver
	webserver *echo.Echo
}

func NewServer(e environ.Getenver, ws *echo.Echo) *server {
	s := &server{
		env:       e,
		webserver: ws,
	}
	s.routes()
	return s
}

func (s *server) Serve() error {
	addr := s.env.Getenv("APP_HTTP_ADDRESS")

	port := s.env.Getenv("APP_HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	hs := http.Server{
		Addr:    fmt.Sprintf("%s:%s", addr, port),
		Handler: s,
	}

	return hs.ListenAndServe()
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.webserver.ServeHTTP(w, r)
}
