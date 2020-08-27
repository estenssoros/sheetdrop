package middle

import (
	"github.com/estenssoros/sheetdrop/constants"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func DBInjector(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(constants.ContextDB, db)
			return next(c)
		}
	}
}
