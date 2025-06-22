package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func (s *server) handleHealthcheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}
}
