package controllers

import (
	"path/filepath"

	"github.com/99designs/gqlgen/graphql"
	"github.com/estenssoros/sheetdrop/constants"
	"github.com/estenssoros/sheetdrop/internal/process"
	"github.com/estenssoros/sheetdrop/models"
	"github.com/pkg/errors"
)

// SchemaHeaders gets headers for a schema
func (c *Controller) SchemaHeaders(schemaID int) ([]*models.Header, error) {
	headers := []*models.Header{}
	return headers, c.Where("schema_id=?", schemaID).Order("idx").Find(&headers).Error
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

// SchemaByID fetch a schema by id
func (c *Controller) SchemaByID(schemaID int) (*models.Schema, error) {
	schema := &models.Schema{}
	return schema, c.Where("id=?", schemaID).First(schema).Error
}

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

// CreateSchemaInput input to create schema
type CreateSchemaInput struct {
	ResourceID int
	Name       string
	File       *graphql.Upload
}

// CreateSchema creates a schema with a resource, name and file
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

func (c *Controller) DeleteSchemaByID(id int) (*models.Schema, error) {
	schema, err := c.SchemaByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "c.SchemaByID")
	}
	return schema, c.Delete(schema).Error
}
