package models

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

const (
	// SourceTypeExcel excel files
	SourceTypeExcel = "excel"
	// SourceTypeCSV csv files
	SourceTypeCSV = "csv"
	// SourceTypeGoogleSheets google sheets documents
	SourceTypeGoogleSheets = "google-sheets"
	// SourceTypeGoogleDrive google drive files
	SourceTypeGoogleDrive = "google-drive"
	// SourceTypeDropBox drop box files
	SourceTypeDropBox = "drop-box"
)

// Schema source information for data
type Schema struct {
	ID          int `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `gorm:"index"`
	ResourceID  int
	Name        *string `gorm:"column:name"`
	StartRow    int
	StartColumn int
	SourceType  string `gorm:"type:varchar(10)"`
	SourceURI   string
	UUID        uuid.UUID `gorm:"type:varchar(36);unique"`
}

// BeforeCreate before create operations
func (s *Schema) BeforeCreate(tx *gorm.DB) error {
	s.UUID = uuid.Must(uuid.NewV4())
	return nil
}

func (s Schema) String() string {
	ju, _ := json.MarshalIndent(s, "", " ")
	return string(ju)
}

// SchemaMigration for gorm foreign key
type SchemaMigration struct {
	Schema
	Resource *Resource `gorm:"foreignKey:ResourceID"`
}

// TableName implements tablenameable
func (s SchemaMigration) TableName() string {
	return `schemas`
}
