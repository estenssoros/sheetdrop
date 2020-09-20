package controllers

import (
	"github.com/estenssoros/sheetdrop/models"
)

type API interface {
	GetUserAPIs(*models.User) ([]*models.API, error)
	CreateAPIForUser(*models.User) (*models.API, error)
	GetAPIByID(uint) (*models.API, error)
	GetUserFromAPIID(uint) (*models.User, error)
}

func (c *Controller) GetUserAPIs(user *models.User) ([]*models.API, error) {
	apis := []*models.API{}
	return apis, c.db.Where("user_id=?", user.ID).Find(&apis).Error
}

func (c *Controller) CreateAPIForUser(user *models.User) (*models.API, error) {
	api := &models.API{OwnerID: user.ID}
	return api, c.db.Create(api).Error
}

func (c *Controller) GetAPIByID(id uint) (*models.API, error) {
	api := &models.API{}
	return api, c.db.Where("id=?", id).First(api).Error
}

func (c *Controller) GetUserFromAPIID(apiID uint) (*models.User, error) {
	user := &models.User{}
	return user, c.db.Joins("JOIN api ON api.user_id = user.id").Where("api.id=?", apiID).First(user).Error
}
