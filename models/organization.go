package models

import "time"

const (
	LevelFree = iota
	LevelPaid1
	LevelAdmin = 69
)

type Organization struct {
	ID           int `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `gorm:"index"`
	Name         *string
	AccountLevel int
}

// TableName implements tablenameable
func (o Organization) TableName() string {
	return `organization`
}

type OrganizationUser struct {
	ID             int `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `gorm:"index"`
	OrganizationID int
	UserID         int
}

// TableName implements tablenameable
func (o OrganizationUser) TableName() string {
	return `organization_user`
}
