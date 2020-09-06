package controllers

import (
	"github.com/estenssoros/sheetdrop/internal/models"
	"gorm.io/gorm"
)

func GetUserAPIs(db *gorm.DB, user *models.User) ([]*models.API, error) {
	apis := []*models.API{}
	return apis, db.Where("user_id=?", user.ID).Find(&apis).Error
}

func CreateAPIForUser(db *gorm.DB, user *models.User) (*models.API, error) {
	api := &models.API{UserID: user.ID}
	return api, db.Create(api).Error
}

func GetAPIByID(db *gorm.DB, id int) (*models.API, error) {
	api := &models.API{}
	return api, db.Where("id=?", id).First(api).Error
}

func GetUserFromAPIID(db *gorm.DB, apiID int) (*models.User, error) {
	user := &models.User{}
	return user, db.Model(&models.API{}).Joins("JOIN api ON api.user_id = user.id").Where("api.id=?", apiID).First(user).Error
}
