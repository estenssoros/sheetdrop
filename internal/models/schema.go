package models

import "encoding/json"

type Schema struct {
	StartRow    int
	StartColumn int
	Headers     map[string]int
	DataTypes   []*DataType
}

func (s Schema) String() string {
	ju, _ := json.MarshalIndent(s, "", " ")
	return string(ju)
}
