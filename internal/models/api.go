package models

import (
	"encoding/json"
	"time"
)

// API api endpoint for a user
type API struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `json:"-"`
	UserID    int
	Name      *string `gorm:"type:varchar(50)"`
}

// TableName implements tablenameable
func (a API) TableName() string {
	return `api`
}

func (a API) String() string {
	ju, _ := json.MarshalIndent(a, "", " ")
	return string(ju)
}
