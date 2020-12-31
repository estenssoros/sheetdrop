package controllers

import (
	"github.com/estenssoros/sheetdrop/models"
)

// UserResources gets a user's resources
func (c *Controller) UserResources(userID int) ([]*models.Resource, error) {
	m := []*models.Resource{}
	return m, c.
		Where("owner_id=?", userID).
		Find(&m).Error
}

// CreateResourceForUser creates an resource for a user
func (c *Controller) CreateResourceForUser(user *models.User) (*models.Resource, error) {
	m := &models.Resource{OwnerID: user.ID}
	return m, c.Create(m).Error
}

// ResourceByID gets a resource by id
func (c *Controller) ResourceByID(id int) (*models.Resource, error) {
	m := &models.Resource{}
	return m, c.
		Where("id=?", id).
		First(m).Error
}

// UserFromResourceID gets a resource's user
func (c *Controller) UsersFromResourceID(id int) ([]*models.User, error) {
	m := []*models.User{}
	return m, c.
		Joins("JOIN resources ON resources.owner_id = users.id").
		Where("resources.id=?", id).
		Find(&m).Error
}

// ResourceSchemas gets schemas for an resource
func (c *Controller) ResourceSchemas(resourceID int) ([]*models.Schema, error) {
	m := []*models.Schema{}
	return m, c.Where("resource_id=?", resourceID).Find(&m).Error
}

// CreateSchemaForResource creates a schema for an resource
func (c *Controller) CreateSchemaForResource(resource *models.Resource) (*models.Schema, error) {
	schema := &models.Schema{
		ResourceID: resource.ID,
	}
	return schema, c.Create(schema).Error
}

// ListResources lists resources
func (c *Controller) ListResources() ([]*models.Resource, error) {
	m := []*models.Resource{}
	return m, c.Find(&m).Error
}
