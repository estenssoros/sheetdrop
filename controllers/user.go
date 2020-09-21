package controllers

import (
	"github.com/estenssoros/sheetdrop/models"
)

type User interface {
	GetUserByName(string) (*models.User, error)
	GetOrCreateUserByName(string) (*models.User, error)
	GetUserByID(int) (*models.User, error)
	UserOrganizations(*models.User) ([]*models.Organization, error)
}

func (c *Controller) GetUserByName(userName string) (*models.User, error) {
	user := &models.User{}
	return user, c.db.Where("user_name=?", userName).First(user).Error
}

func (c *Controller) GetOrCreateUserByName(userName string) (*models.User, error) {
	user := &models.User{}
	return user, c.db.Where(models.User{UserName: userName}).FirstOrCreate(user).Error
}

func (c *Controller) GetUserByID(userID int) (*models.User, error) {
	user := &models.User{}
	return user, c.db.Where("id=?", userID).First(user).Error
}

func (c *Controller) UserOrganizations(obj *models.User) ([]*models.Organization, error) {
	m := []*models.Organization{}
	return m, c.db.
		Joins("JOIN organization_user ON organization_user.organization_id=organization.id").
		Where("organization_user.user_id=?", obj.ID).
		Find(&m).
		Error
}
