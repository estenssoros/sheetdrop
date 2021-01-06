package controllers

import (
	"testing"

	"github.com/satori/uuid"
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

func TestOrganizationsByIDs(t *testing.T) {
	out, err := ctl(t).OrganizationsByIDs([]int{1})
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}

func TestCreateDeleteOrg(t *testing.T) {
	org, err := ctl(t).CreateOrgWithName(uuid.NewV4().String())
	if err != nil {
		t.Fatal(err)
	}
	_, err = ctl(t).CreateOrgUser(org.ID, 1)
	if err != nil {
		t.Fatal(err)
	}
	_, err = ctl(t).DeleteOrgByID(org.ID)
	if err != nil {
		t.Fatal(err)
	}
}
