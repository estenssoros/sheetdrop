package controllers

import (
	"github.com/estenssoros/sheetdrop/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// SchemaHeaders gets headers for a schema
func (c *Controller) SchemaHeaders(schema *models.Schema) ([]*models.Header, error) {
	headers := []*models.Header{}
	return headers, c.Where("schema_id=?", schema.ID).Find(&headers).Error
}

// SchemaHeadersSet headers set for a schema
func (c *Controller) SchemaHeadersSet(schema *models.Schema) (*models.HeaderSet, error) {
	headers, err := c.SchemaHeaders(schema)
	if err != nil {
		return nil, errors.Wrap(err, "GetSchemaHeaders")
	}
	return models.NewHeaderSet(headers), nil
}

// UpdateSchemaInput input into update schema
type UpdateSchemaInput struct {
	ID   *int
	Name *string
}

// Validate validates inputs
func (input *UpdateSchemaInput) Validate(db *gorm.DB) error {
	if input.ID == nil {
		return errors.New("missing id")
	}
	if input.Name == nil {
		return errors.New("missing name")
	}
	return nil
}

// UpdateSchema updates a schema
func (c *Controller) UpdateSchema(input *UpdateSchemaInput) (*models.Schema, error) {
	if err := c.Validate(input); err != nil {
		return nil, errors.Wrap(err, "validate")
	}
	schema := &models.Schema{}
	if err := c.Where("id=?", *input.ID).First(schema).Error; err != nil {
		return nil, err
	}
	schema.Name = input.Name
	return schema, c.Save(schema).Error
}

// UserFromSchemaID get a user from a schema
func (c *Controller) UserFromSchemaID(schemaID int) (*models.User, error) {
	user := &models.User{}
	return user, c.Model(user).
		Joins("JOIN resources ON resources.owner_id = users.id").
		Joins("JOIN schema ON schema.resource_id = resources.id").
		Where("schemas.id=?", schemaID).
		First(user).Error
}

// DeleteSchema deletes a schema
func (c *Controller) DeleteSchema(schema *models.Schema) error {
	return c.Delete(schema).Error
}

// SchemaByID fetch a schema by id
func (c *Controller) SchemaByID(schemaID int) (*models.Schema, error) {
	schema := &models.Schema{}
	return schema, c.Where("id=?", schemaID).First(schema).Error
}

// SchemaRelations populate schema relations
func (c *Controller) SchemaRelations(schemas []*models.Schema) error {
	ids := make([]int, len(schemas))
	for i := 0; i < len(schemas); i++ {
		ids[i] = schemas[i].ID
	}
	headersMap := map[int][]*models.Header{}
	{
		headers := []*models.Header{}
		if err := c.Where("schema_id IN (?)", ids).Order("idx").Find(&headers).Error; err != nil {
			return err
		}
		for _, h := range headers {
			if headers, ok := headersMap[h.SchemaID]; ok {
				headers = append(headers, h)
				headersMap[h.SchemaID] = headers
			} else {
				headersMap[h.SchemaID] = []*models.Header{h}
			}
		}
	}
	for _, schema := range schemas {
		if headers, ok := headersMap[schema.ID]; ok {
			schema.Headers = headers
		}
	}
	return nil
}

func (c *Controller) SchemasForResources(obj *models.Resource) ([]*models.Schema, error) {
	return nil, errNotImplemented
}
