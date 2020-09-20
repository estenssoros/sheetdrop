package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func handleRest(c echo.Context) error {
	url := c.Param("url")
	fmt.Println(url)
	return nil
}
