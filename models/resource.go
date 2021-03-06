package models

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Resource api endpoint for a user
type Resource struct {
	ID             int `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `gorm:"index"`
	OrganizationID int
	OwnerID        int
	Name           *string   `gorm:"type:varchar(50)"`
	AuthToken      uuid.UUID `gorm:"type:varchar(36);unique"`
}

func (r Resource) String() string {
	ju, _ := json.MarshalIndent(r, "", " ")
	return string(ju)
}

// BeforeCreate before creeate operations
func (r *Resource) BeforeCreate(tx *gorm.DB) error {
	r.AuthToken = uuid.Must(uuid.NewV4())
	return nil
}

// ResourceMigration for gorm foreign key
type ResourceMigration struct {
	Resource
	Organization *Organization `gorm:"foreignKey:OrganizationID"`
}

// TableName implements tablenameable
func (r ResourceMigration) TableName() string {
	return `resources`
}
