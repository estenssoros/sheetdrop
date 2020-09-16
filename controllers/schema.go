package controllers

import (
	"github.com/estenssoros/sheetdrop/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Schema interface {
	GetSchemasForAPI(*models.API) ([]*models.Schema, error)
	CreateSchemaForAPI(*models.API) (*models.Schema, error)
	UpdateSchema(*UpdateSchemaInput) (*models.Schema, error)
	GetUserFromSchemaID(schemaID uint) (*models.User, error)
	DeleteSchema(*models.Schema) error
	SchemaByID(uint) (*models.Schema, error)
	GetSchemaRelations([]*models.Schema) error
}

func (c *Controller) GetSchemasForAPI(api *models.API) ([]*models.Schema, error) {
	schemas := []*models.Schema{}
	return schemas, c.db.Where("api_id=?", api.ID).Find(&schemas).Error
}

func (c *Controller) CreateSchemaForAPI(api *models.API) (*models.Schema, error) {
	schema := &models.Schema{
		APIID: api.ID,
	}
	return schema, c.db.Create(schema).Error
}

type UpdateSchemaInput struct {
	ID   *int
	Name *string
}

func (input *UpdateSchemaInput) Validate(db *gorm.DB) error {
	if input.ID == nil {
		return errors.New("missing id")
	}
	if input.Name == nil {
		return errors.New("missing name")
	}
	return nil
}

func (c *Controller) UpdateSchema(input *UpdateSchemaInput) (*models.Schema, error) {
	if err := c.Validate(input); err != nil {
		return nil, errors.Wrap(err, "validate")
	}
	schema := &models.Schema{}
	if err := c.db.Where("id=?", *input.ID).First(schema).Error; err != nil {
		return nil, err
	}
	schema.Name = input.Name
	return schema, c.db.Save(schema).Error
}

func (c *Controller) GetUserFromSchemaID(schemaID uint) (*models.User, error) {
	user := &models.User{}
	return user, c.db.Model(user).
		Joins("JOIN api ON api.user_id = user.id").
		Joins("JOIN schema ON schema.api_id = api.id").
		Where("schema.id=?", schemaID).
		First(user).Error
}

func (c *Controller) DeleteSchema(schema *models.Schema) error {
	return c.db.Delete(schema).Error
}

func (c *Controller) SchemaByID(schemaID uint) (*models.Schema, error) {
	schema := &models.Schema{}
	return schema, c.db.Where("id=?", schemaID).First(schema).Error
}

func (c *Controller) GetSchemaRelations(schemas []*models.Schema) error {
	ids := make([]uint, len(schemas))
	for i := 0; i < len(schemas); i++ {
		ids[i] = schemas[i].ID
	}
	headersMap := map[uint][]*models.Header{}
	{
		headers := []*models.Header{}
		if err := c.db.Where("schema_id IN (?)", ids).Order("idx").Find(&headers).Error; err != nil {
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
