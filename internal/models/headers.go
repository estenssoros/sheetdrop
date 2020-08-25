package models

type Header struct {
	ID       int
	Name     string `gorm:"type:varchar(50)"`
	Index    int
	DataType string `gorm:"type:varchar(10)"`
}

// TableName implements tablenameable
func (h Header) TableName() string {
	return `header`
}
