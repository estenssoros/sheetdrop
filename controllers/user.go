package controllers

import (
	"github.com/estenssoros/sheetdrop/models"
)

type User interface {
	GetUserByName(string) (*models.User, error)
	GetOrCreateUserByName(string) (*models.User, error)
	GetUserByID(uint) (*models.User, error)
}

func (db *Controller) GetUserByName(userName string) (*models.User, error) {
	user := &models.User{}
	return user, db.Where("user_name=?", userName).First(user).Error
}

func (db *Controller) GetOrCreateUserByName(userName string) (*models.User, error) {
	user := &models.User{}
	return user, db.Where(models.User{UserName: userName}).FirstOrCreate(user).Error
}

func (db *Controller) GetUserByID(userID uint) (*models.User, error) {
	user := &models.User{}
	return user, db.Where("id=?", userID).First(user).Error
}
