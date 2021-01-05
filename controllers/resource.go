package controllers

import (
	"github.com/estenssoros/sheetdrop/models"
	"github.com/pkg/errors"
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

// UsersFromResourceID gets a resource's user
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

// ResourceSchemaCountByID batch resource schema counts
func (c *Controller) ResourceSchemaCountByID(ids []int) ([]int, []error) {
	var counts = []struct {
		ID  int
		Cnt int
	}{}
	err := c.Model(&models.Resource{}).
		Select("resources.id, count(schemas.id) cnt").
		Joins("LEFT JOIN schemas ON schemas.resource_id = resources.id").
		Group("resources.id").Find(&counts).Error
	if err != nil {
		return nil, []error{err}
	}
	lookup := map[int]int{}
	for _, value := range counts {
		lookup[value.ID] = value.Cnt
	}
	out := make([]int, len(ids))
	for i, id := range ids {
		out[i] = lookup[id]
	}
	return out, nil
}

// CreateResourceInput input for create resource
type CreateResourceInput struct {
	OrganizationID int
	ResourceName   string
}

// CreateResouce creates a resource with name and organization id
func (c *Controller) CreateResouce(input *CreateResourceInput) (*models.Resource, error) {
	resource := &models.Resource{
		OrganizationID: input.OrganizationID,
		Name:           &input.ResourceName,
	}
	return resource, c.Create(resource).Error
}

// DeleteResourceByID deletes a resource by id
func (c *Controller) DeleteResourceByID(id int) (*models.Resource, error) {
	resource, err := c.ResourceByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "c.ResourceByID")
	}
	return resource, c.Delete(resource).Error
}
