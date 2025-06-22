package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func (s *server) handleSayHello() echo.HandlerFunc {
	return func(c echo.Context) error {
		if requestId := c.Response().Header().Get(echo.HeaderXRequestID); requestId != "" {
			c.Logger().Debug("saying hello to request id: " + requestId)
		} else {
			c.Logger().Debug("saying hello to a stranger")
		}
		return c.String(http.StatusOK, "hello world!")
	}
}
