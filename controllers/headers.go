package controllers

import (
	"github.com/estenssoros/sheetdrop/models"
	"github.com/pkg/errors"
)

// HeaderForeignKeys finds headers with matching foreign keys
func (c *Controller) HeaderForeignKeys(headerID int) ([]*models.Header, error) {
	headers := []*models.Header{}
	query := c.Model(&models.HeaderHeader{}).
		Select("foreign_header_id").
		Where("header_id=?", headerID)
	return headers, c.Where("id IN (?)", query).Find(&headers).Error
}

// HeadersByIDs find headers by ids
func (c *Controller) HeadersByIDs(ids []int) ([]*models.Header, []error) {
	values := []*models.Header{}
	if err := c.Where("id in (?)", ids).Find(&values).Error; err != nil {
		return nil, []error{errors.Wrap(err, "find values")}
	}
	lookup := map[int]*models.Header{}
	for _, value := range values {
		lookup[value.ID] = value
	}
	out := make([]*models.Header, len(ids))
	for i, id := range ids {
		out[i] = lookup[id]
	}
	return out, nil
}

// HeadersBySchemaIDs group queries for finding headers for schemas
func (c *Controller) HeadersBySchemaIDs(schemaIds []int) ([][]*models.Header, []error) {
	values := []*models.Header{}
	if err := c.Where("schema_id in (?)", schemaIds).Find(&values).Error; err != nil {
		return nil, []error{errors.Wrap(err, "find values")}
	}
	lookup := map[int][]*models.Header{}
	for _, value := range values {
		lookup[value.SchemaID] = append(lookup[value.SchemaID], value)
	}
	out := make([][]*models.Header, len(schemaIds))
	for i, id := range schemaIds {
		out[i] = lookup[id]
	}
	return out, nil
}
