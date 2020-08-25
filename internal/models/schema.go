package models

import "encoding/json"

type Schema struct {
	ID          int
	APIID       int `gorm:"column:api_id"`
	StartRow    int
	StartColumn int
	Headers     []*Header
}

func (s Schema) String() string {
	ju, _ := json.MarshalIndent(s, "", " ")
	return string(ju)
}

// TableName implements tablenameable
func (s Schema) TableName() string {
	return `schema`
}
