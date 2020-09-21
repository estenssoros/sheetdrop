package controllers

import "gorm.io/gorm"

// Interface all the things a controller must do
type Interface interface {
	User
	Org
	API
	Schema
	File
	DB() *gorm.DB
	Delete(v interface{}) error
}

// Controller does the controlling
type Controller struct {
	db *gorm.DB
}

// New creates a new controller
func New(db *gorm.DB) *Controller {
	return &Controller{db}
}

// DB return the db connection
func (c *Controller) DB() *gorm.DB {
	return c.db
}

// Validator a thing that can be validated
type Validator interface {
	Validate(db *gorm.DB) error
}

// Validate validates if it is a validator
func (c *Controller) Validate(v interface{}) error {
	valid, ok := v.(Validator)
	if ok {
		return valid.Validate(c.db)
	}
	return nil
}

// Delete deletes a model
func (c *Controller) Delete(v interface{}) error {
	return c.db.Delete(v).Error
}
