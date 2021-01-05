package models

import "time"

const (
	LevelFree = iota
	LevelPaid
	LevelAdmin = 69
)

// Organization groups users
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

// OrganizationUser models the many to many relationship between organizations and users
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

// OrganizationUserMigration for gorm foreign key
type OrganizationUserMigration struct {
	OrganizationUser
	Organization *Organization `gorm:"foreignKey:OrganizationID"`
	User         *User         `gorm:"foreignKey:UserID"`
}

// TableName implements tablenameable
func (o OrganizationUserMigration) TableName() string {
	return `organization_user`
}
