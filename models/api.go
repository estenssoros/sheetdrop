package models

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// API api endpoint for a user
type API struct {
	ID             int `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `gorm:"index"`
	OrganizationID int
	OwnerID        int
	Name           *string   `gorm:"type:varchar(50)"`
	AuthToken      uuid.UUID `gorm:"type:varchar(36);unique"`
}

// TableName implements tablenameable
func (a API) TableName() string {
	return `api`
}

func (a API) String() string {
	ju, _ := json.MarshalIndent(a, "", " ")
	return string(ju)
}

func (a *API) BeforeCreate(tx *gorm.DB) error {
	a.AuthToken = uuid.Must(uuid.NewV4())
	return nil
}
