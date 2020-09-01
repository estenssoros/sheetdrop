package models

import (
	"encoding/json"
	"time"
)

// User a user of the application
type User struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	UserName  string `gorm:"type:varchar(50)"`
}

// TableName implements tablenameable
func (u User) TableName() string {
	return `user`
}

func (u User) String() string {
	ju, _ := json.MarshalIndent(u, "", " ")
	return string(ju)
}
