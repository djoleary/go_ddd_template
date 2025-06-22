package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func (s *server) handleGreeting() echo.HandlerFunc {
	return func(c echo.Context) error {
		if requestId := c.Response().Header().Get(echo.HeaderXRequestID); requestId != "" {
			c.Logger().Debug("saying hello to request id: " + requestId)
		} else {
			c.Logger().Debug("saying hello to a stranger")
		}
		return c.String(http.StatusOK, "hello world!")
	}
}

func (s *server) handleGreetingByName() echo.HandlerFunc {
	return func(c echo.Context) error {
		n := c.Param("name")
		if n == "" {
			c.Logger().Debug("ignoring a stranger")
			return c.NoContent(http.StatusOK)
		}
		c.Logger().Debug("saying hello to: " + n)
		return c.String(http.StatusOK, "hello "+n+"!")
	}
}
