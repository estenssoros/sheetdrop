package controllers

import (
	"github.com/estenssoros/sheetdrop/models"
)

// API interface for api operations
type API interface {
	UserAPIs(*models.User) ([]*models.API, error)
	CreateAPIForUser(*models.User) (*models.API, error)
	APIByID(int) (*models.API, error)
	UserFromAPIID(int) (*models.User, error)
	APISChemas(*models.API) ([]*models.Schema, error)
	CreateSchemaForAPI(*models.API) (*models.Schema, error)
}

// UserAPIs gets a user's apis
func (c *Controller) UserAPIs(user *models.User) ([]*models.API, error) {
	apis := []*models.API{}
	return apis, c.db.Where("owner_id=?", user.ID).Find(&apis).Error
}

// CreateAPIForUser creates an ap for a user
func (c *Controller) CreateAPIForUser(user *models.User) (*models.API, error) {
	api := &models.API{OwnerID: user.ID}
	return api, c.db.Create(api).Error
}

// APIByID gets an api by id
func (c *Controller) APIByID(id int) (*models.API, error) {
	api := &models.API{}
	return api, c.db.Where("id=?", id).First(api).Error
}

// UserFromAPIID gets an api's user
func (c *Controller) UserFromAPIID(apiID int) (*models.User, error) {
	user := &models.User{}
	return user, c.db.Joins("JOIN api ON api.owner_id = user.id").Where("api.id=?", apiID).First(user).Error
}

// APISChemas gets schemas for an api
func (c *Controller) APISChemas(obj *models.API) ([]*models.Schema, error) {
	m := []*models.Schema{}
	return m, c.db.Where("api_id=?", obj.ID).Find(&m).Error
}

func (c *Controller) CreateSchemaForAPI(api *models.API) (*models.Schema, error) {
	schema := &models.Schema{
		APIID: api.ID,
	}
	return schema, c.db.Create(schema).Error
}
