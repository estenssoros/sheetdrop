package server

import (
	"net/http"
	"strconv"

	"github.com/estenssoros/sheetdrop/constants"
	"github.com/estenssoros/sheetdrop/controllers"
	"github.com/estenssoros/sheetdrop/models"
	"github.com/labstack/echo/v4"
)

func getAPIHandler(c echo.Context) error {
	apiID, err := strconv.Atoi(c.Param("id"))
	ctl := c.Get(constants.ContextController).(controllers.Interface)
	user, err := ctl.GetUserFromAPIID(uint(apiID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if user.UserName != c.Get(constants.ContextUserName).(string) {
		return c.JSON(http.StatusForbidden, "user names do not match")
	}
	api, err := ctl.GetAPIByID(uint(apiID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, api)
}

func createAPIHandler(c echo.Context) error {
	api := &models.API{}
	if err := c.Bind(api); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if api.Name == nil {
		return c.JSON(http.StatusBadRequest, "missing api name")
	}
	ctl := c.Get(constants.ContextController).(controllers.Interface)
	user, err := ctl.GetUserByName(c.Get(constants.ContextUserName).(string))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	newAPI := &models.API{
		UserID: user.ID,
		Name:   api.Name,
	}
	if err := ctl.DB().Create(&newAPI).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, newAPI)
}

func deleteAPIHandler(c echo.Context) error {
	api := &models.API{}
	if err := c.Bind(api); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctl := c.Get(constants.ContextController).(controllers.Interface)
	user, err := ctl.GetUserByID(api.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	userName := c.Get(constants.ContextUserName).(string)
	if user.UserName != userName {
		return c.JSON(http.StatusForbidden, "username not valid")
	}
	if err := ctl.DB().Delete(&api).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func updateAPIHandler(c echo.Context) error {
	api := &models.API{}
	if err := c.Bind(api); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctl := c.Get(constants.ContextController).(controllers.Interface)
	user, err := ctl.GetUserByID(api.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	userName := c.Get(constants.ContextUserName).(string)
	if user.UserName != userName {
		return c.JSON(http.StatusForbidden, "username not valid")
	}
	if err := ctl.DB().Save(api).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, api)
}

func getAPIsHandler(c echo.Context) error {
	userName := c.Get(constants.ContextUserName).(string)
	ctl := c.Get(constants.ContextDB).(controllers.Interface)
	user, err := ctl.GetOrCreateUserByName(userName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	apis, err := ctl.GetUserAPIs(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if len(apis) == 0 {
		api, err := ctl.CreateAPIForUser(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		apis = append(apis, api)
	}
	return c.JSON(http.StatusOK, apis)
}
