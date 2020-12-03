package middle

import (
	"github.com/estenssoros/sheetdrop/constants"
	"github.com/estenssoros/sheetdrop/controllers"
	"github.com/labstack/echo/v4"
)

// DBInjector inserts the DBInjectors onto the context
func DBInjector(ctl *controllers.Controller) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(constants.ContextController, ctl)
			return next(c)
		}
	}
}
