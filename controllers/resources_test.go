package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceByID(t *testing.T) {
	out, err := ctl(t).ResourceByID(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}

func TestUsersFromResourceID(t *testing.T) {
	out, err := ctl(t).UsersFromResourceID(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}

func TestResourceSchemas(t *testing.T) {
	out, err := ctl(t).ResourceSchemas(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}

func TestListResources(t *testing.T) {
	out, err := ctl(t).ListResources()
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}

func TestUserResources(t *testing.T) {
	out, err := ctl(t).UserResources(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}

func TestUsersFromSchemaID(t *testing.T) {
	out, err := ctl(t).UsersFromSchemaID(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}

func TestResourceSchemaCountByID(t *testing.T) {
	out, err := ctl(t).ResourceSchemaCountByID([]int{1})
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}
