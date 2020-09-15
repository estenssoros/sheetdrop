package controllers

import (
	"github.com/estenssoros/sheetdrop/models"
)

type Org interface {
	GetUserOrgs(*models.User) ([]*models.Organization, error)
}

// GetUserOrgs get orgs for a user
func (db *Controller) GetUserOrgs(user *models.User) ([]*models.Organization, error) {
	orgs := []*models.Organization{}
	return orgs, db.
		Joins("JOIN organization_user ON organization_user.organization_id = organization.id").
		Where("organization_user.user_id=?", user.ID).
		Find(&orgs).
		Error
}
