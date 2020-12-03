package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	generated1 "github.com/estenssoros/sheetdrop/graph/generated"
	"github.com/estenssoros/sheetdrop/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, userName string) (*models.User, error) {
	return r.GetOrCreateUserByName(userName)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (string, error) {
	return "", r.DeleteUserByID(id)
}

func (r *mutationResolver) CreateOrg(ctx context.Context, userID int, orgName string) (*models.Organization, error) {
	hasOrg, err := r.UserHasOrg(userID, orgName)
	if err != nil {
		return nil, fmt.Errorf("r.UserHasOrg:%v", err)
	}
	if hasOrg {
		return nil, fmt.Errorf("user already has org named: %s", orgName)
	}
	org, err := r.CreateOrgWithName(orgName)
	if err != nil {
		return nil, fmt.Errorf("r.CreateOrgByName: %v", err)
	}
	if _, err := r.CreateOrgUser(org.ID, userID); err != nil {
		return nil, fmt.Errorf("r.CreateOrgUser: %v", err)
	}
	return org, nil
}

func (r *mutationResolver) AddUserToOrg(ctx context.Context, userID int, orgID int) (*models.Organization, error) {
	_, err := r.CreateOrgUser(orgID, userID)
	if err != nil {
		return nil, err
	}
	return r.OrganizationByID(orgID)
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
