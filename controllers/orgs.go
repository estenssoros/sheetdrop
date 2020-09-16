package controllers

import (
	"github.com/estenssoros/sheetdrop/models"
	"github.com/estenssoros/sheetdrop/responses"
	"github.com/pkg/errors"
)

type Org interface {
	GetUserOrgsResponse(*models.User) ([]*responses.Organization, error)
	CreateOrg(*CreateOrgInput) error
	UserCanEditOrg(*models.User, *models.Organization) (bool, error)
	UpdateOrg(*models.Organization) error
}

// GetUserOrgsResponse get orgs for a user
func (c *Controller) GetUserOrgsResponse(user *models.User) ([]*responses.Organization, error) {
	orgs := []*responses.Organization{}
	query := c.db.Model(&models.Organization{}).
		Select(`organization.id, organization.name, organization.account_level, count(*) members`).
		Joins("JOIN organization_user ON organization_user.organization_id = organization.id").
		Where("organization_user.user_id=?", user.ID).
		Group("organization.id, organization.name,organization.account_level")
	return orgs, query.
		Find(&orgs).
		Error
}

type CreateOrgInput struct {
	Org  *models.Organization
	User *models.User
}

func (c *Controller) CreateOrg(input *CreateOrgInput) error {
	if err := c.Validate(input); err != nil {
		return errors.Wrap(err, "Validate")
	}
	if err := c.db.Create(input.Org).Error; err != nil {
		return errors.Wrap(err, "create org")
	}
	orgUser := &models.OrganizationUser{
		OrganizationID: input.Org.ID,
		UserID:         input.User.ID,
	}
	return errors.Wrap(c.db.Create(orgUser).Error, "create org user")
}

func (c *Controller) UserCanEditOrg(user *models.User, org *models.Organization) (bool, error) {
	var count int64
	err := c.db.Model(&models.OrganizationUser{}).
		Where("user_id=?", user.ID).
		Where("organization_id=?", org.ID).
		Count(&count).Error
	return count > 0, err
}

func (c *Controller) UpdateOrg(org *models.Organization) error {
	return c.db.Save(org).Error
}
