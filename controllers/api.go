package controllers

import (
	"github.com/estenssoros/sheetdrop/models"
)

type API interface {
	UserAPIs(*models.User) ([]*models.API, error)
	CreateAPIForUser(*models.User) (*models.API, error)
	APIByID(int) (*models.API, error)
	UserFromAPIID(int) (*models.User, error)
	APISChemas(*models.API) ([]*models.Schema, error)
}

func (c *Controller) UserAPIs(user *models.User) ([]*models.API, error) {
	apis := []*models.API{}
	return apis, c.db.Where("user_id=?", user.ID).Find(&apis).Error
}

func (c *Controller) CreateAPIForUser(user *models.User) (*models.API, error) {
	api := &models.API{OwnerID: user.ID}
	return api, c.db.Create(api).Error
}

func (c *Controller) APIByID(id int) (*models.API, error) {
	api := &models.API{}
	return api, c.db.Where("id=?", id).First(api).Error
}

func (c *Controller) UserFromAPIID(apiID int) (*models.User, error) {
	user := &models.User{}
	return user, c.db.Joins("JOIN api ON api.user_id = user.id").Where("api.id=?", apiID).First(user).Error
}

func (c *Controller) APISChemas(obj *models.API) ([]*models.Schema, error) {
	m := []*models.Schema{}
	return m, c.db.Where("api_id=?", obj.ID).Find(&m).Error
}
