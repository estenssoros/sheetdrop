package responses

import "github.com/labstack/echo/v4"

func Error(c echo.Context, statusCode int, err error) error {
	return c.JSON(statusCode, err.Error())
}
