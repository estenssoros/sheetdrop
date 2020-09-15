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

func (db *Controller) GetUserAPIs(user *models.User) ([]*models.API, error) {
	apis := []*models.API{}
	return apis, db.Where("user_id=?", user.ID).Find(&apis).Error
}

func (db *Controller) CreateAPIForUser(user *models.User) (*models.API, error) {
	api := &models.API{UserID: user.ID}
	return api, db.Create(api).Error
}

func (db *Controller) GetAPIByID(id uint) (*models.API, error) {
	api := &models.API{}
	return api, db.Where("id=?", id).First(api).Error
}

func (db *Controller) GetUserFromAPIID(apiID uint) (*models.User, error) {
	user := &models.User{}
	return user, db.Joins("JOIN api ON api.user_id = user.id").Where("api.id=?", apiID).First(user).Error
}
