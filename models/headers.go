package models

import (
	"reflect"
	"time"
)

// HeaderHeader WIP header-header relationship
type HeaderHeader struct {
	ID              int
	HeaderID        int
	ForeignHeaderID int
}

// TableName implements tablenameable
func (h HeaderHeader) TableName() string {
	return `header_header`
}

// Header field information from a data source
type Header struct {
	ID        int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	SchemaID  int
	Name      string `gorm:"type:varchar(50)"`
	Index     int    `gorm:"column:idx"`
	DataType  string `gorm:"type:varchar(10)"`
	IsID      bool
}

// TableName implements tablenameable
func (h Header) TableName() string {
	return `header`
}

// HeaderMigration for gorm foreign key
type HeaderMigration struct {
	Header
	Schema *Schema `gorm:"foreignKey:SchemaID"`
}

// TableName implements tablenameable
func (h HeaderMigration) TableName() string {
	return `header`
}

// HeaderSet for set operations with headers
type HeaderSet struct {
	Headers map[string]*Header
}

// NewHeaderSet creates a new header set
func NewHeaderSet(headers []*Header) *HeaderSet {
	headerMap := map[string]*Header{}
	for _, header := range headers {
		headerMap[header.Name] = header
	}
	return &HeaderSet{
		Headers: headerMap,
	}
}

// Has does header exists in headerset
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
