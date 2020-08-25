package models

type API struct {
	ID     int
	UserID int
	Name   string `gorm:"type:varchar(50)"`
}

// TableName implements tablenameable
func (a API) TableName() string {
	return `api`
}
