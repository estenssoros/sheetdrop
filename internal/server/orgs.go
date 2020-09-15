package server

import (
	"net/http"

	"github.com/estenssoros/sheetdrop/constants"
	"github.com/estenssoros/sheetdrop/controllers"
	"github.com/labstack/echo/v4"
)

func getOrgsHandler(c echo.Context) error {
	userName := c.Get(constants.ContextUserName).(string)
	ctl := c.Get(constants.ContextDB).(controllers.Interface)
	user, err := ctl.GetOrCreateUserByName(userName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	orgs, err := ctl.GetUserOrgs(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, orgs)
}
