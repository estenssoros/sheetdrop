package controllers

import (
	"testing"

	"github.com/estenssoros/sheetdrop/orm"
)

func ctl(t *testing.T) *Controller {
	db, err := orm.Connect()
	if err != nil {
		t.Fatal(err)
	}
	return New(db)
}
