package controllers

import (
	"testing"

	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
)

func TestListUsers(t *testing.T) {
	out, err := ctl(t).ListUsers()
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}
func TestGetUserByName(t *testing.T) {
	out, err := ctl(t).GetUserByName("sebastian")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
	assert.Equal(t, "sebastian", out.UserName)
}
func TestGetOrCreateUserByName(t *testing.T) {
	out, err := ctl(t).GetOrCreateUserByName("sebastian")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
	assert.Equal(t, "sebastian", out.UserName)
}
func TestGetUserByID(t *testing.T) {
	out, err := ctl(t).GetUserByID(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}
func TestUserOrganizations(t *testing.T) {
	out, err := ctl(t).UserOrganizations(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}
func TestUsersByIds(t *testing.T) {
	out, err := ctl(t).UsersByIds([]int{1})
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}

func TestCreateDeleteUser(t *testing.T) {
	user, err := ctl(t).CreateUser(uuid.NewV4().String())
	if err != nil {
		t.Fatal(err)
	}
	_, err = ctl(t).DeleteUserByID(user.ID)
	if err != nil {
		t.Fatal(err)
	}
}
