package middle

import (
	"github.com/estenssoros/sheetdrop/constants"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DBInjector(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(constants.ContextDB, db)
			return next(c)
		}
	}
}
