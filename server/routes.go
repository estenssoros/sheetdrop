package server

import (
	"github.com/labstack/echo/v4"
)

func restRoutes(e *echo.Group) {
	e.GET("/:url", handleRest)
}
