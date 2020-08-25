package models

type User struct {
	ID       int
	UserName string `gorm:"type:varchar(50)"`
}

// TableName implements tablenameable
func (u User) TableName() string {
	return `user`
}
