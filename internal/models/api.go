package models

import "time"

// API api endpoint for a user
type API struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int
	Name      string `gorm:"type:varchar(50)"`
}

// TableName implements tablenameable
func (a API) TableName() string {
	return `api`
}