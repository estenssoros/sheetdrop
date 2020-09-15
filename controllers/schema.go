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

func GetSchemasForAPI(db *gorm.DB, api *models.API) ([]*models.Schema, error) {
	schemas := []*models.Schema{}
	return schemas, db.Where("api_id=?", api.ID).Find(&schemas).Error
}

func CreateSchemaForAPI(db *gorm.DB, api *models.API) (*models.Schema, error) {
	schema := &models.Schema{
		APIID: api.ID,
	}
	return schema, db.Create(schema).Error
}

type UpdateSchemaInput struct {
	ID   *int
	Name *string
}

func (input *UpdateSchemaInput) Validate() error {
	if input.ID == nil {
		return errors.New("missing id")
	}
	if input.Name == nil {
		return errors.New("missing name")
	}
	return nil
}

func UpdateSchema(db *gorm.DB, input *UpdateSchemaInput) (*models.Schema, error) {
	schema := &models.Schema{}
	if err := db.Where("id=?", *input.ID).First(schema).Error; err != nil {
		return nil, err
	}
	schema.Name = input.Name
	return schema, db.Save(schema).Error
}

func GetUserFromSchemaID(db *gorm.DB, schemaID uint) (*models.User, error) {
	user := &models.User{}
	return user, db.Model(user).
		Joins("JOIN api ON api.user_id = user.id").
		Joins("JOIN schema ON schema.api_id = api.id").
		Where("schema.id=?", schemaID).
		First(user).Error
}

func DeleteSchema(db *gorm.DB, schema *models.Schema) error {
	return db.Delete(schema).Error
}

func SchemaByID(db *gorm.DB, schemaID uint) (*models.Schema, error) {
	schema := &models.Schema{}
	return schema, db.Where("id=?", schemaID).First(schema).Error
}

func GetSchemaRelations(db *gorm.DB, schemas []*models.Schema) error {
	ids := make([]uint, len(schemas))
	for i := 0; i < len(schemas); i++ {
		ids[i] = schemas[i].ID
	}
	headersMap := map[uint][]*models.Header{}
	{
		headers := []*models.Header{}
		if err := db.Where("schema_id IN (?)", ids).Order("idx").Find(&headers).Error; err != nil {
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
