package server

import (
	"github.com/estenssoros/sheetdrop/constants"
	"github.com/estenssoros/sheetdrop/controllers"
	"github.com/labstack/echo/v4"
)

func extractController(c echo.Context) controllers.Interface {
	return c.Get(constants.ContextController).(controllers.Interface)
}

func extractUserName(c echo.Context) string {
	return c.Get(constants.ContextUserName).(string)
}
