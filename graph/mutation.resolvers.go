package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/estenssoros/sheetdrop/graph/generated"
	"github.com/estenssoros/sheetdrop/graph/model"
	"github.com/estenssoros/sheetdrop/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, userName string) (*models.User, error) {
	return r.GetOrCreateUserByName(userName)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (string, error) {
	return "", r.DeleteUserByID(id)
}

func (r *mutationResolver) CreateOrg(ctx context.Context, input model.CreateOrgInput) (*models.Organization, error) {
	hasOrg, err := r.UserHasOrg(input.UserID, input.OrgName)
	if err != nil {
		return nil, fmt.Errorf("r.UserHasOrg:%v", err)
	}
	if hasOrg {
		return nil, fmt.Errorf("user already has org named: %s", input.OrgName)
	}
	org, err := r.CreateOrgWithName(input.OrgName)
	if err != nil {
		return nil, fmt.Errorf("r.CreateOrgByName: %v", err)
	}
	if _, err := r.CreateOrgUser(org.ID, input.UserID); err != nil {
		return nil, fmt.Errorf("r.CreateOrgUser: %v", err)
	}
	return org, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
