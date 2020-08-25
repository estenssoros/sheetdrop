package models

import "time"

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
