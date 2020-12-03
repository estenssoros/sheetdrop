package controllers

import "gorm.io/gorm"

// Controller does the controlling
type Controller struct {
	*gorm.DB
}

// New creates a new controller
func New(db *gorm.DB) *Controller {
	return &Controller{db}
}

// Validator a thing that can be validated
type Validator interface {
	Validate(db *gorm.DB) error
}

// Validate validates if it is a validator
func (c *Controller) Validate(v interface{}) error {
	valid, ok := v.(Validator)
	if ok {
		return valid.Validate(c.DB)
	}
	return nil
}
