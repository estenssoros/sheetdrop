package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

// User a user of the application
type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(50)"`
	APIs     []*API
}

// TableName implements tablenameable
func (u User) TableName() string {
	return `user`
}

func (u User) String() string {
	ju, _ := json.MarshalIndent(u, "", " ")
	return string(ju)
}
