package controllers

import "gorm.io/gorm"

type Interface interface {
	User
	Org
	API
	Schema
	DB() *gorm.DB
}

type Controller struct {
	*gorm.DB
}
