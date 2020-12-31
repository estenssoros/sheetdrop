package server

import (
	"net/http"
	"strconv"

	"github.com/estenssoros/sheetdrop/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func getResourceHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	_, err = ctl(c).UsersFromResourceID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// TODO: create func to user has resource
	// if user.UserName != usr(c) {
	// 	return c.JSON(http.StatusForbidden, "user names do not match")
	// }
	resource, err := ctl(c).ResourceByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resource)
}

type createResourceRequest struct {
	Name           *string
	OrganizationID *int
}

func (req *createResourceRequest) Validate(c echo.Context) (*models.Resource, error) {
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
	hasOrg, err := ctl(c).UserCanEditOrg(user.ID, *req.OrganizationID)
	if err != nil {
		return nil, errors.Wrap(err, "UserHasOrg")
	}
	if !hasOrg {
		return nil, errors.New("user is not org member")
	}
	return &models.Resource{
		OrganizationID: *req.OrganizationID,
		OwnerID:        user.ID,
		Name:           req.Name,
	}, nil
}

func createResourceHandler(c echo.Context) error {
	req := &createResourceRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	api, err := req.Validate(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := ctl(c).Create(api).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, api)
}

type deleteResourceRequest struct {
	ID *int
}

func deleteResourceHandler(c echo.Context) error {
	req := &deleteResourceRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user, err := ctl(c).GetUserByName(usr(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	resource, err := ctl(c).ResourceByID(*req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if resource.OwnerID != user.ID {
		return c.JSON(http.StatusForbidden, "user cannot edit org")
	}
	// TODO deal with orphaned schemas, headers, etc
	if err := ctl(c).Delete(&resource).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

type updateResourceRequest struct {
	ID   *int
	Name *string
}

func (req updateResourceRequest) Validate() error {
	if req.ID == nil {
		return errors.New("missing id")
	}
	if req.Name == nil {
		return errors.New("missing name")
	}
	return nil
}

func updateResourceHandler(c echo.Context) error {
	req := &updateResourceRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	resource, err := ctl(c).ResourceByID(*req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	userModel, err := ctl(c).GetUserByID(resource.OwnerID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if userModel.UserName != usr(c) {
		return c.JSON(http.StatusForbidden, "user cannot edit api")
	}
	resource.Name = req.Name
	if err := ctl(c).Save(resource).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resource)
}

func getResourcesHandler(c echo.Context) error {
	user, err := ctl(c).GetOrCreateUserByName(usr(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	resources, err := ctl(c).UserResources(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if len(resources) == 0 {
		resource, err := ctl(c).CreateResourceForUser(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		resources = append(resources, resource)
	}
	return c.JSON(http.StatusOK, resources)
}
