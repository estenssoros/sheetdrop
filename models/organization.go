package models

import (
	"gorm.io/gorm"
)

const (
	LevelFree = iota
	LevelPaid1
	LevelAdmin = 69
)

type Organization struct {
	gorm.Model
	Name         string
	AccountLevel int
}

// TableName implements tablenameable
func (o Organization) TableName() string {
	return `organization`
}

type OrganizationUser struct {
	gorm.Model
	OrganizationID uint
	UserID         uint
}

// TableName implements tablenameable
func (o OrganizationUser) TableName() string {
	return `organization_user`
}
