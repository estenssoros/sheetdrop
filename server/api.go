package server

import (
	"net/http"
	"strconv"

	"github.com/estenssoros/sheetdrop/constants"
	"github.com/estenssoros/sheetdrop/controllers"
	"github.com/estenssoros/sheetdrop/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func getAPIHandler(c echo.Context) error {
	apiID, err := strconv.Atoi(c.Param("id"))
	ctl := extractController(c)
	user, err := ctl.UserFromAPIID(apiID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if user.UserName != c.Get(constants.ContextUserName).(string) {
		return c.JSON(http.StatusForbidden, "user names do not match")
	}
	api, err := ctl.APIByID(apiID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, api)
}

type createAPIRequest struct {
	Name           *string
	OrganizationID *int
}

func (req *createAPIRequest) ValidateAPI(c echo.Context, ctl controllers.Interface) (*models.API, error) {
	if req.Name == nil {
		return nil, errors.New("missing api name")
	}
	if req.OrganizationID == nil {
		return nil, errors.New("missing org id")
	}
	user, err := ctl.GetUserByName(c.Get(constants.ContextUserName).(string))
	if err != nil {
		return nil, errors.Wrap(err, "GetUserByName")
	}
	hasOrg, err := ctl.UserCanEditOrg(user, &models.Organization{
		ID: *req.OrganizationID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "UserHasOrg")
	}
	if !hasOrg {
		return nil, errors.New("user is not org member")
	}
	return &models.API{
		OrganizationID: *req.OrganizationID,
		OwnerID:        user.ID,
		Name:           req.Name,
	}, nil
}

func createAPIHandler(c echo.Context) error {
	req := &createAPIRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctl := extractController(c)
	api, err := req.ValidateAPI(c, ctl)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := ctl.DB().Create(api).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, api)
}

type deleteAPIRequest struct {
	ID *int
}

func deleteAPIHandler(c echo.Context) error {
	req := &deleteAPIRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctl := extractController(c)
	user, err := ctl.GetUserByName(c.Get(constants.ContextUserName).(string))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	api, err := ctl.APIByID(*req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if api.OwnerID != user.ID {
		return c.JSON(http.StatusForbidden, "user cannot edit org")
	}
	// TODO deal with orphaned schemas, headers, etc
	if err := ctl.DB().Delete(&api).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

type updateAPIRequest struct {
	ID   *int
	Name *string
}

func (req updateAPIRequest) Validate() error {
	if req.ID == nil {
		return errors.New("missing id")
	}
	if req.Name == nil {
		return errors.New("missing name")
	}
	return nil
}

func updateAPIHandler(c echo.Context) error {
	req := &updateAPIRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctl := extractController(c)
	api, err := ctl.APIByID(*req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	user, err := ctl.GetUserByID(api.OwnerID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	userName := c.Get(constants.ContextUserName).(string)
	if user.UserName != userName {
		return c.JSON(http.StatusForbidden, "user cannot edit api")
	}
	api.Name = req.Name
	if err := ctl.DB().Save(api).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, api)
}

func getAPIsHandler(c echo.Context) error {
	userName := c.Get(constants.ContextUserName).(string)
	ctl := extractController(c)
	user, err := ctl.GetOrCreateUserByName(userName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	apis, err := ctl.UserAPIs(user)
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
