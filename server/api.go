package server

import (
	"github.com/estenssoros/sheetdrop/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

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
