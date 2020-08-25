package models

import "encoding/json"

const (
	DataTypeString = "string"
	DataTypeInt    = "int"
	DataTypeFloat  = "float"
	DataTypeTime   = "time"
)

type DataType struct {
	Idx  int
	Type string
}

func (d DataType) String() string {
	ju, _ := json.Marshal(d)
	return string(ju)
}
