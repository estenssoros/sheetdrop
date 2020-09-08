package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

// API api endpoint for a user
type API struct {
	gorm.Model
	UserID  uint
	Name    *string `gorm:"type:varchar(50)"`
	Schemas []*Schema
}

// TableName implements tablenameable
func (a API) TableName() string {
	return `api`
}

func (a API) String() string {
	ju, _ := json.MarshalIndent(a, "", " ")
	return string(ju)
}
