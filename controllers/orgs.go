package controllers

import (
	"github.com/estenssoros/sheetdrop/internal/helpers"
	"github.com/estenssoros/sheetdrop/models"
	"github.com/estenssoros/sheetdrop/responses"
	"github.com/pkg/errors"
)

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
func (c *Controller) CreateOrg(input *CreateOrgInput) error {
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
func (c *Controller) UserCanEditOrg(userID, orgID int) (bool, error) {
	var count int64
	err := c.Model(&models.OrganizationUser{}).
		Where("user_id=?", userID).
		Where("organization_id=?", orgID).
		Count(&count).Error
	return count > 0, err
}

// OrganizationByID gets organization by id
func (c *Controller) OrganizationByID(id int) (*models.Organization, error) {
	org := &models.Organization{}
	return org, c.Where("id=?", id).First(org).Error
}

// OrganizationUsers get users for an organization
func (c *Controller) OrganizationUsers(orgID int) ([]*models.User, error) {
	users := []*models.User{}
	return users, c.
		Joins("JOIN organization_user ON organization_user.user_id = users.id").
		Where("organization_user.organization_id=?", orgID).Find(&users).Error
}

// OrganizationResources gets resources for an organization
func (c *Controller) OrganizationResources(orgID int) ([]*models.Resource, error) {
	resources := []*models.Resource{}
	return resources, c.Where("organization_id=?", orgID).Find(&resources).Error
}

// UserHasOrg checks to see if a user has an org
func (c *Controller) UserHasOrg(userID, orgID int) (bool, error) {
	var count int64
	query := c.Model(&models.OrganizationUser{}).
		Joins("JOIN organization ON organization.id = organization_user.organization_id").
		Where("organization.id = ?", orgID)
	return count > 0, query.Count(&count).Error
}

// CreateOrgWithName creates an org with a given name
func (c *Controller) CreateOrgWithName(orgName string) (*models.Organization, error) {
	org := &models.Organization{
		Name: helpers.StringPtr(orgName),
	}
	return org, c.Create(org).Error
}

// CreateOrgUser creates an org-user relationship
func (c *Controller) CreateOrgUser(orgID, userID int) (*models.OrganizationUser, error) {
	orgUser := &models.OrganizationUser{
		OrganizationID: orgID,
		UserID:         userID,
	}
	return orgUser, c.Create(orgUser).Error
}

// ListOrganizations lists the organizations
func (c *Controller) ListOrganizations() ([]*models.Organization, error) {
	m := []*models.Organization{}
	return m, c.Find(&m).Error
}

// OrganizationsByIDs group queries for finding organization by id
func (c *Controller) OrganizationsByIDs(ids []int) ([]*models.Organization, []error) {
	values := []*models.Organization{}
	if err := c.Where("id in (?)", ids).Find(&values).Error; err != nil {
		return nil, []error{errors.Wrap(err, "find values")}
	}
	lookup := map[int]*models.Organization{}
	for _, value := range values {
		lookup[value.ID] = value
	}
	out := make([]*models.Organization, len(ids))
	for i, id := range ids {
		out[i] = lookup[id]
	}
	return out, nil
}

// DeleteOrgByID deletes an org by id
func (c *Controller) DeleteOrgByID(id int) (*models.Organization, error) {
	org, err := c.OrganizationByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "c.OrganizationByID")
	}
	return org, c.Delete(org).Error
}
