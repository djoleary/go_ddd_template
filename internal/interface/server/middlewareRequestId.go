package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func middlewareAddRequestId() echo.MiddlewareFunc {
	return middleware.RequestID()
}
