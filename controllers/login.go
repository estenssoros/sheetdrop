package controllers

import (
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

func (db *Controller) Login(input *LoginInput) error {
	if err := input.Validate(); err != nil {
		return errors.Wrap(err, "input.Validate")
	}
	_, err := db.GetUserByName(*input.UserName)
	return err
}
