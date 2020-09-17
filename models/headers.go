package models

import (
	"reflect"

	"gorm.io/gorm"
)

// Header field information from a data source
type Header struct {
	gorm.Model
	SchemaID uint
	Name     string `gorm:"type:varchar(50)"`
	Index    uint   `gorm:"column:idx"`
	DataType string `gorm:"type:varchar(10)"`
	IsID     bool
}

// TableName implements tablenameable
func (h Header) TableName() string {
	return `header`
}

type HeaderSet struct {
	Headers map[string]*Header
}

func NewHeaderSet(headers []*Header) *HeaderSet {
	headerMap := map[string]*Header{}
	for _, header := range headers {
		headerMap[header.Name] = header
	}
	return &HeaderSet{
		Headers: headerMap,
	}
}
func (h *HeaderSet) Has(other *Header) bool {
	_, ok := h.Headers[other.Name]
	return ok
}

func (h *HeaderSet) HasEqual(other *Header) bool {
	header, ok := h.Headers[other.Name]
	if !ok {
		return false
	}
	return reflect.DeepEqual(header, other)
}

func (h *HeaderSet) ToCreate(other []*Header) (headers []*Header) {
	for _, header := range other {
		if !h.Has(header) {
			headers = append(headers, header)
		}
	}
	return headers
}

func (h *HeaderSet) ToUpdate(other []*Header) (headers []*Header) {
	for _, header := range other {
		if h.Has(header) && !h.HasEqual(header) {
			header.ID = h.Headers[header.Name].ID
			headers = append(headers, header)
		}
	}
	return headers
}

func (h *HeaderSet) ToDelete(other []*Header) (headers []*Header) {
	headerSet := NewHeaderSet(other)
	for _, header := range h.Headers {
		if !headerSet.Has(header) {
			headers = append(headers, header)
		}
	}
	return headers
}
