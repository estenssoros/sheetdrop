package controllers

import (
	"github.com/estenssoros/sheetdrop/models"
	"github.com/estenssoros/sheetdrop/responses"
	"github.com/pkg/errors"
)

// Org everything an org must do
type Org interface {
	OrganizationByID(int) (*models.Organization, error)
	OrganizationUsers(*models.Organization) ([]*models.User, error)
	UserOrgsResponse(*models.User) ([]*responses.Organization, error)
	CreateOrg(*CreateOrgInput) error
	UserCanEditOrg(*models.User, *models.Organization) (bool, error)
	UpdateOrg(*models.Organization) error
}

// UserOrgsResponse get orgs for a user
func (c *Controller) UserOrgsResponse(user *models.User) ([]*responses.Organization, error) {
	orgs := []*responses.Organization{}
	query := c.Model(&models.Organization{}).
		Select(`organization.id, organization.name, organization.account_level, count(*) members`).
		Joins("JOIN organization_user ON organization_user.organization_id = organization.id").
		Where("organization_user.user_id=?", user.ID).
		Group("organization.id, organization.name,organization.account_level")
	return orgs, query.
		Find(&orgs).
		Error
}

// CreateOrgInput in put for creating org
type CreateOrgInput struct {
	Org  *models.Organization
	User *models.User
}

// CreateOrg create an org
func (c *Controller) c(input *CreateOrgInput) error {
	if err := c.Validate(input); err != nil {
		return errors.Wrap(err, "Validate")
	}
	if err := c.Create(input.Org).Error; err != nil {
		return errors.Wrap(err, "create org")
	}
	orgUser := &models.OrganizationUser{
		OrganizationID: input.Org.ID,
		UserID:         input.User.ID,
	}
	return errors.Wrap(c.Create(orgUser).Error, "create org user")
}

// UserCanEditOrg can a user edit an org
func (c *Controller) UserCanEditOrg(user *models.User, org *models.Organization) (bool, error) {
	var count int64
	err := c.Model(&models.OrganizationUser{}).
		Where("user_id=?", user.ID).
		Where("organization_id=?", org.ID).
		Count(&count).Error
	return count > 0, err
}

// UpdateOrg saves an org
func (c *Controller) UpdateOrg(org *models.Organization) error {
	return c.Save(org).Error
}

// OrganizationByID gets organization by id
func (c *Controller) OrganizationByID(id int) (*models.Organization, error) {
	org := &models.Organization{}
	return org, c.Where("id=?", id).First(org).Error
}

// OrganizationUsers get users for an organization
func (c *Controller) OrganizationUsers(obj *models.Organization) ([]*models.User, error) {
	users := []*models.User{}
	return users, c.
		Joins("JOIN organization_user ON organization_user.user_id = user.id").
		Where("organization_user.organization_id=?", obj.ID).Find(&users).Error
}
