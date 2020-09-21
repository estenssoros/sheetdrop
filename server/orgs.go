package server

import (
	"net/http"

	"github.com/estenssoros/sheetdrop/controllers"
	"github.com/estenssoros/sheetdrop/models"
	"github.com/labstack/echo/v4"
)

func getOrgsHandler(c echo.Context) error {
	userName := extractUserName(c)
	ctl := extractController(c)
	user, err := ctl.GetOrCreateUserByName(userName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	orgs, err := ctl.GetUserOrgsResponse(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, orgs)
}

func createOrgHandler(c echo.Context) error {
	model := &models.Organization{}
	if err := c.Bind(model); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if model.Name == nil {
		return c.JSON(http.StatusBadRequest, "missing api name")
	}
	ctl := extractController(c)
	user, err := ctl.GetUserByName(extractUserName(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	input := &controllers.CreateOrgInput{
		Org:  model,
		User: user,
	}
	if err := ctl.CreateOrg(input); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, model)
}

func updateOrgHandler(c echo.Context) error {
	model := &models.Organization{}
	if err := c.Bind(model); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctl := extractController(c)
	user, err := ctl.GetUserByName(extractUserName(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	canEdit, err := ctl.UserCanEditOrg(user, model)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if !canEdit {
		return c.JSON(http.StatusForbidden, "user cannot edit org")
	}
	if err := ctl.UpdateOrg(model); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, model)
}

func deleteOrgHandler(c echo.Context) error {
	model := &models.Organization{}
	if err := c.Bind(model); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctl := extractController(c)
	user, err := ctl.GetUserByName(extractUserName(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	canEdit, err := ctl.UserCanEditOrg(user, model)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if !canEdit {
		return c.JSON(http.StatusForbidden, "user cannot edit org")
	}
	if err := ctl.DB().Delete(model).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
