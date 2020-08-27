package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type LoginInput struct {
	UserName *string `json:"userName"`
}

func (i *LoginInput) Validate() error {
	if i.UserName == nil {
		return errors.New("no username")
	}
	return nil
}

func Login(db *gorm.DB, input *LoginInput) error {
	if err := input.Validate(); err != nil {
		return errors.Wrap(err, "input.Validate")
	}
	_, err := GetUserByName(db, *input.UserName)
	return err
}
