package controllers

import (
	"github.com/estenssoros/sheetdrop/internal/models"
	"github.com/jinzhu/gorm"
)

func GetUserByName(db *gorm.DB, userName string) (*models.User, error) {
	user := &models.User{}
	return user, db.Where("user_name=?", userName).First(user).Error
}
