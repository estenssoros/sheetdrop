package controllers

import "gorm.io/gorm"

type Interface interface {
	User
	Org
	API
	Schema
	File
	DB() *gorm.DB
}

type Controller struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Controller {
	return &Controller{db}
}

func (c *Controller) DB() *gorm.DB {
	return c.db
}

type Validator interface {
	Validate(db *gorm.DB) error
}

func (c *Controller) Validate(v interface{}) error {
	valid, ok := v.(Validator)
	if ok {
		return valid.Validate(c.db)
	}
	return nil
}
