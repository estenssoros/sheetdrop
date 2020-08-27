package controllers

import (
	"github.com/estenssoros/sheetdrop/internal/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func GetUserAPIs(db *gorm.DB, userName string) ([]*models.API, error) {
	apis := []*models.API{}
	user, err := GetUserByName(db, userName)
	if err != nil {
		return nil, errors.Wrap(err, "GetUserByName")
	}
	return apis, db.Where("user_id=?", user.ID).Find(&apis).Error
}
