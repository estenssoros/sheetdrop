package server

import (
	"net/http"

	"github.com/estenssoros/sheetdrop/models"
	"github.com/labstack/echo/v4"
)

func getOrgsHandler(c echo.Context) error {
	user, err := ctl(c).GetOrCreateUserByName(usr(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	orgs, err := ctl(c).GetUserOrgsResponse(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, orgs)
}

// func createOrgHandler(c echo.Context) error {
// 	model := &models.Organization{}
// 	if err := c.Bind(model); err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}
// 	if model.Name == nil {
// 		return c.JSON(http.StatusBadRequest, "missing api name")
// 	}
// 	user, err := ctl(c).GetUserByName(usr(c))
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}
// 	input := &controllers.CreateOrgInput{
// 		Org:  model,
// 		User: user,
// 	}
// 	if err := ctl(c).CreateOrg(input); err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, model)
// }

func updateOrgHandler(c echo.Context) error {
	org := &models.Organization{}
	if err := c.Bind(org); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user, err := ctl(c).GetUserByName(usr(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	canEdit, err := ctl(c).UserCanEditOrg(user.ID, org.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if !canEdit {
		return c.JSON(http.StatusForbidden, "user cannot edit org")
	}
	if err := ctl(c).Save(org).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, org)
}

func deleteOrgHandler(c echo.Context) error {
	org := &models.Organization{}
	if err := c.Bind(org); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user, err := ctl(c).GetUserByName(usr(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	canEdit, err := ctl(c).UserCanEditOrg(user.ID, org.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if !canEdit {
		return c.JSON(http.StatusForbidden, "user cannot edit org")
	}
	if err := ctl(c).Delete(org).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
