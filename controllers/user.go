package controllers

import (
	"github.com/estenssoros/sheetdrop/models"
)

type User interface {
	GetUserByName(string) (*models.User, error)
	GetOrCreateUserByName(string) (*models.User, error)
	GetUserByID(uint) (*models.User, error)
}

func (c *Controller) GetUserByName(userName string) (*models.User, error) {
	user := &models.User{}
	return user, c.db.Where("user_name=?", userName).First(user).Error
}

func (c *Controller) GetOrCreateUserByName(userName string) (*models.User, error) {
	user := &models.User{}
	return user, c.db.Where(models.User{UserName: userName}).FirstOrCreate(user).Error
}

func (c *Controller) GetUserByID(userID uint) (*models.User, error) {
	user := &models.User{}
	return user, c.db.Where("id=?", userID).First(user).Error
}
