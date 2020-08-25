package models

import "time"

// Header field information from a data source
type Header struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"type:varchar(50)"`
	Index     int
	DataType  string `gorm:"type:varchar(10)"`
	IsID      bool
}

// TableName implements tablenameable
func (h Header) TableName() string {
	return `header`
}
