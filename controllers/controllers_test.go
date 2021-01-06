package controllers

import (
	"testing"

	"github.com/estenssoros/sheetdrop/orm"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func ctl(t *testing.T) *Controller {
	db, err := orm.Connect()
	if err != nil {
		t.Fatal(err)
	}
	return New(db)
}

type validator struct{}

func (v validator) Validate(db *gorm.DB) error {
	return nil
}
func TestValidator(t *testing.T) {
	assert.Nil(t, ctl(t).Validate(&validator{}))
	assert.Nil(t, ctl(t).Validate(7))
}
