package models

import (
	"gorm.io/gorm"
)

// Header field information from a data source
type Header struct {
	gorm.Model
	SchemaID uint
	Name     string `gorm:"type:varchar(50)"`
	Index    int    `gorm:"column:idx"`
	DataType string `gorm:"type:varchar(10)"`
	IsID     bool
}

// TableName implements tablenameable
func (h Header) TableName() string {
	return `header`
}
