package controllers

import (
	"github.com/estenssoros/sheetdrop/models"
	"github.com/pkg/errors"
)

func (c *Controller) ListUsers() ([]*models.User, error) {
	users := []*models.User{}
	return users, c.Find(&users).Error
}

// GetUserByName gets user model by username
func (c *Controller) GetUserByName(userName string) (*models.User, error) {
	user := &models.User{}
	return user, c.Where("user_name=?", userName).First(user).Error
}

// GetOrCreateUserByName gets or creates a user
func (c *Controller) GetOrCreateUserByName(userName string) (*models.User, error) {
	user := &models.User{}
	return user, c.Where(models.User{UserName: userName}).FirstOrCreate(user).Error
}

// GetUserByID gets a user by id
func (c *Controller) GetUserByID(userID int) (*models.User, error) {
	user := &models.User{}
	return user, c.Where("id=?", userID).First(user).Error
}

// UserOrganizations gets a user's organizations
func (c *Controller) UserOrganizations(obj *models.User) ([]*models.Organization, error) {
	m := []*models.Organization{}
	return m, c.
		Joins("JOIN organization_user ON organization_user.organization_id=organization.id").
		Where("organization_user.user_id=?", obj.ID).
		Find(&m).
		Error
}

var errNotImplemented = errors.New("not implemented")

func (c *Controller) GetUserOrgsResponse(user *models.User) ([]*models.Organization, error) {
	return nil, errNotImplemented
}
