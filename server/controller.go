package server

import (
	"github.com/estenssoros/sheetdrop/constants"
	"github.com/estenssoros/sheetdrop/controllers"
	"github.com/labstack/echo/v4"
)

func ctl(c echo.Context) *controllers.Controller {
	return c.Get(constants.ContextController).(*controllers.Controller)
}

func usr(c echo.Context) string {
	return c.Get(constants.ContextUserName).(string)
}
