package server

import (
	"net/http"
	"strconv"

	"github.com/estenssoros/sheetdrop/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func getAPIHandler(c echo.Context) error {
	apiID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	user, err := ctl(c).UserFromAPIID(apiID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if user.UserName != usr(c) {
		return c.JSON(http.StatusForbidden, "user names do not match")
	}
	api, err := ctl(c).APIByID(apiID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, api)
}

type createAPIRequest struct {
	Name           *string
	OrganizationID *int
}

func (req *createAPIRequest) ValidateAPI(c echo.Context) (*models.API, error) {
	if req.Name == nil {
		return nil, errors.New("missing api name")
	}
	if req.OrganizationID == nil {
		return nil, errors.New("missing org id")
	}
	user, err := ctl(c).GetUserByName(usr(c))
	if err != nil {
		return nil, errors.Wrap(err, "GetUserByName")
	}
	hasOrg, err := ctl(c).UserCanEditOrg(user, &models.Organization{
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
	api, err := req.ValidateAPI(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := ctl(c).Create(api).Error; err != nil {
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
	user, err := ctl(c).GetUserByName(usr(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	api, err := ctl(c).APIByID(*req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if api.OwnerID != user.ID {
		return c.JSON(http.StatusForbidden, "user cannot edit org")
	}
	// TODO deal with orphaned schemas, headers, etc
	if err := ctl(c).Delete(&api).Error; err != nil {
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
	api, err := ctl(c).APIByID(*req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	userModel, err := ctl(c).GetUserByID(api.OwnerID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if userModel.UserName != usr(c) {
		return c.JSON(http.StatusForbidden, "user cannot edit api")
	}
	api.Name = req.Name
	if err := ctl(c).Save(api).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, api)
}

func getAPIsHandler(c echo.Context) error {
	user, err := ctl(c).GetOrCreateUserByName(usr(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	apis, err := ctl(c).UserAPIs(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if len(apis) == 0 {
		api, err := ctl(c).CreateAPIForUser(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		apis = append(apis, api)
	}
	return c.JSON(http.StatusOK, apis)
}
