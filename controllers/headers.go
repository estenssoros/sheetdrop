package controllers

import "github.com/estenssoros/sheetdrop/models"

// HeaderForeignKeys finds headers with matching foreign keys
func (c *Controller) HeaderForeignKeys(headerID int) ([]*models.Header, error) {
	headers := []*models.Header{}
	query := c.Model(&models.HeaderHeader{}).
		Select("foreign_header_id").
		Where("header_id=?", headerID)
	return headers, c.Where("id IN (?)", query).Find(&headers).Error
}
