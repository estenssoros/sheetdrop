package models

import (
	"encoding/json"
	"time"
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
	ID          int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	APIID       int     `gorm:"column:api_id" json:"api_id"`
	Name        *string `gorm:"column:name"`
	StartRow    int
	StartColumn int
	Headers     []*Header
	SourceType  string `gorm:"varchar(10)"`
	SourceURI   string
}

func (s Schema) String() string {
	ju, _ := json.MarshalIndent(s, "", " ")
	return string(ju)
}

// TableName implements tablenameable
func (s Schema) TableName() string {
	return `api_schema`
}
