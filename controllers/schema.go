package controllers

import (
	"github.com/estenssoros/sheetdrop/internal/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

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

func GetUserFromSchemaID(db *gorm.DB, schemaID int) (*models.User, error) {
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
