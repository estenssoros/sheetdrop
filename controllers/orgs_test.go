package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationByID(t *testing.T) {
	out, err := ctl(t).OrganizationByID(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}

func TestOrganizationUsers(t *testing.T) {
	out, err := ctl(t).OrganizationUsers(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}

func TestOrganizationResources(t *testing.T) {
	out, err := ctl(t).OrganizationResources(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}
func TestListOrganizations(t *testing.T) {
	out, err := ctl(t).ListOrganizations()
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)

}
func TestUserCanEditOrg(t *testing.T) {
	out, err := ctl(t).UserCanEditOrg(1, 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, out)
}
func TestUserHasOrg(t *testing.T) {
	out, err := ctl(t).UserHasOrg(1, 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, out)
}
