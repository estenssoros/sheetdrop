package controllers

import (
	"path/filepath"

	"github.com/99designs/gqlgen/graphql"
	"github.com/estenssoros/sheetdrop/constants"
	"github.com/estenssoros/sheetdrop/internal/process"
	"github.com/estenssoros/sheetdrop/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// SchemaHeaders gets headers for a schema
func (c *Controller) SchemaHeaders(schemaID int) ([]*models.Header, error) {
	headers := []*models.Header{}
	return headers, c.Where("schema_id=?", schemaID).Order("idx").Find(&headers).Error
}

// SchemaHeadersSet headers set for a schema
func (c *Controller) SchemaHeadersSet(schema *models.Schema) (*models.HeaderSet, error) {
	headers, err := c.SchemaHeaders(schema.ID)
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

// UsersFromSchemaID get users from a schema
func (c *Controller) UsersFromSchemaID(schemaID int) ([]*models.User, error) {
	users := []*models.User{}
	return users, c.Model(&models.User{}).
		Joins("JOIN resources ON resources.owner_id = users.id").
		Joins("JOIN schemas ON schemas.resource_id = resources.id").
		Where("schemas.id=?", schemaID).
		Find(&users).Error
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
// func (c *Controller) SchemaRelations(schemas []*models.Schema) error {
// 	ids := make([]int, len(schemas))
// 	for i := 0; i < len(schemas); i++ {
// 		ids[i] = schemas[i].ID
// 	}
// 	headersMap := map[int][]*models.Header{}
// 	{
// 		headers := []*models.Header{}
// 		if err := c.Where("schema_id IN (?)", ids).Order("idx").Find(&headers).Error; err != nil {
// 			return err
// 		}
// 		for _, h := range headers {
// 			if headers, ok := headersMap[h.SchemaID]; ok {
// 				headers = append(headers, h)
// 				headersMap[h.SchemaID] = headers
// 			} else {
// 				headersMap[h.SchemaID] = []*models.Header{h}
// 			}
// 		}
// 	}
// 	for _, schema := range schemas {
// 		if headers, ok := headersMap[schema.ID]; ok {
// 			schema.Headers = headers
// 		}
// 	}
// 	return nil
// }

// SchemasForResource gets schemas for resource
func (c *Controller) SchemasForResource(resourceID int) ([]*models.Schema, error) {
	schemas := []*models.Schema{}
	return schemas, c.Where("resource_id=?", resourceID).Find(&schemas).Error
}

// ListSchemas lists all schemas
func (c *Controller) ListSchemas() ([]*models.Schema, error) {
	m := []*models.Schema{}
	return m, c.Find(&m).Error
}

// SchemasByIDs group queries for finding schema by ids
func (c *Controller) SchemasByIDs(ids []int) ([]*models.Schema, []error) {
	values := []*models.Schema{}
	if err := c.Where("id in (?)", ids).Find(&values).Error; err != nil {
		return nil, []error{errors.Wrap(err, "find values")}
	}
	lookup := map[int]*models.Schema{}
	for _, value := range values {
		lookup[value.ID] = value
	}
	out := make([]*models.Schema, len(ids))
	for i, id := range ids {
		out[i] = lookup[id]
	}
	return out, nil
}

type CreateSchemaInput struct {
	ResourceID int
	Name       string
	File       *graphql.Upload
}

func (c *Controller) CreateSchema(input *CreateSchemaInput) (*models.Schema, error) {
	schema := &models.Schema{
		ResourceID: input.ResourceID,
		Name:       &input.Name,
		SourceURI:  input.File.Filename,
	}
	ext := filepath.Ext(input.File.Filename)
	var processor = func() (*process.Result, error) { return nil, nil }
	switch ext {
	case constants.ExtensionExcel:
		schema.SourceType = constants.SourceTypeExcel
		processor = func() (*process.Result, error) {
			return process.Excel(schema, input.File.File)
		}
	case constants.ExtensionCSV:
		schema.SourceType = constants.SourceTypeCSV
		processor = func() (*process.Result, error) {
			return process.CSV(schema, input.File.File)
		}
	default:
		return nil, errors.Errorf("extesnion not implemented: %s", ext)
	}
	result, err := processor()
	if err != nil {
		return nil, errors.Wrap(err, "processor")
	}
	{
		schema.StartColumn = result.StartColumn
		schema.StartRow = result.StartRow
	}
	if err := c.Create(schema).Error; err != nil {
		return nil, errors.Wrap(err, "create schema")
	}
	for _, header := range result.Headers {
		header.SchemaID = schema.ID
	}
	return schema, errors.Wrap(c.Create(&result.Headers).Error, "create headers")
}
