package controllers

import (
	"github.com/estenssoros/sheetdrop/models"
	"gorm.io/gorm"
)

func GetUserByName(db *gorm.DB, userName string) (*models.User, error) {
	user := &models.User{}
	return user, db.Where("user_name=?", userName).First(user).Error
}

func GetOrCreateUserByName(db *gorm.DB, userName string) (*models.User, error) {
	user := &models.User{}
	return user, db.Where(models.User{UserName: userName}).FirstOrCreate(user).Error
}

func GetUserByID(db *gorm.DB, userID uint) (*models.User, error) {
	user := &models.User{}
	return user, db.Where("id=?", userID).First(user).Error
}
