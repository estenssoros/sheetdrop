package controllers

import (
	"github.com/estenssoros/sheetdrop/models"
)

// UserAPIs gets a user's apis
func (c *Controller) UserAPIs(user *models.User) ([]*models.API, error) {
	apis := []*models.API{}
	return apis, c.
		Where("owner_id=?", user.ID).
		Find(&apis).Error
}

// CreateAPIForUser creates an ap for a user
func (c *Controller) CreateAPIForUser(user *models.User) (*models.API, error) {
	api := &models.API{OwnerID: user.ID}
	return api, c.Create(api).Error
}

// APIByID gets an api by id
func (c *Controller) APIByID(id int) (*models.API, error) {
	api := &models.API{}
	return api, c.
		Where("id=?", id).
		First(api).Error
}

// UserFromAPIID gets an api's user
func (c *Controller) UserFromAPIID(apiID int) (*models.User, error) {
	user := &models.User{}
	return user, c.
		Joins("JOIN api ON api.owner_id = users.id").
		Where("api.id=?", apiID).
		First(user).Error
}

// APISChemas gets schemas for an api
func (c *Controller) APISChemas(obj *models.API) ([]*models.Schema, error) {
	m := []*models.Schema{}
	return m, c.Where("api_id=?", obj.ID).Find(&m).Error
}

// CreateSchemaForAPI creates a schema for an api
func (c *Controller) CreateSchemaForAPI(api *models.API) (*models.Schema, error) {
	schema := &models.Schema{
		APIID: api.ID,
	}
	return schema, c.Create(schema).Error
}
