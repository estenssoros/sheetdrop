package controllers

import (
	"fmt"
	"testing"

	"github.com/estenssoros/sheetdrop/orm"
)

func TestGetUserFromAPIID(t *testing.T) {
	db, err := orm.Connect()
	if err != nil {
		t.Fatal(err)
	}
	user, err := GetUserFromAPIID(db, 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(user)
}
