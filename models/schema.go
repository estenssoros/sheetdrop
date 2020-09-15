package models

import (
	"encoding/json"

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
	gorm.Model
	APIID       uint    `gorm:"column:api_id" json:"api_id"`
	Name        *string `gorm:"column:name"`
	StartRow    int
	StartColumn int
	Headers     []*Header
	SourceType  string `gorm:"type:varchar(10)"`
	SourceURI   string
	AuthToken   uuid.UUID `gorm:"type:varchar(36);unique"`
}

func (s *Schema) BeforeCreate(tx *gorm.DB) error {
	s.AuthToken = uuid.Must(uuid.NewV4())
	return nil
}

func (s Schema) String() string {
	ju, _ := json.MarshalIndent(s, "", " ")
	return string(ju)
}

// TableName implements tablenameable
func (s Schema) TableName() string {
	return `api_schema`
}
