package controllers

import (
	"github.com/estenssoros/sheetdrop/models"
	"github.com/pkg/errors"
)

type Headers interface {
	GetSChemaHeaders(*models.Schema) ([]*models.Header, error)
	GetSchemaHeadersMap(*models.Schema) (map[string]*models.Header, error)
}

func (c *Controller) GetSchemaHeaders(schema *models.Schema) ([]*models.Header, error) {
	headers := []*models.Header{}
	return headers, c.db.Where("schema_id=?", schema.ID).Find(&headers).Error
}

func (c *Controller) GetSchemaHeadersSet(schema *models.Schema) (*models.HeaderSet, error) {
	headers, err := c.GetSchemaHeaders(schema)
	if err != nil {
		return nil, errors.Wrap(err, "GetSChemaHeaders")
	}
	return models.NewHeaderSet(headers), nil
}
