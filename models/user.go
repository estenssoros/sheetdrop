package models

import (
	"encoding/json"
	"time"
)

// User a user of the application
type User struct {
	ID        int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	UserName  string     `gorm:"type:varchar(50);unique"`
}

func (u User) String() string {
	ju, _ := json.MarshalIndent(u, "", " ")
	return string(ju)
}
