package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	generated1 "github.com/estenssoros/sheetdrop/graph/generated"
	"github.com/estenssoros/sheetdrop/models"
)

func (r *mutationResolver) GetUser(ctx context.Context, id int) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, userName string) (*models.User, error) {
	return r.GetOrCreateUserByName(userName)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*models.User, error) {
	return nil, r.DeleteUserByID(id)
}

func (r *mutationResolver) GetOrg(ctx context.Context, id int) (*models.Organization, error) {
	panic(fmt.Errorf("not implemented"))
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

func (r *mutationResolver) DeleteOrg(ctx context.Context, id int) (*models.Organization, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddUserToOrg(ctx context.Context, userID int, orgID int) (*models.Organization, error) {
	_, err := r.CreateOrgUser(orgID, userID)
	if err != nil {
		return nil, err
	}
	return r.OrganizationByID(orgID)
}

func (r *mutationResolver) RemoveUserFromOrg(ctx context.Context, userID int, orgID int) (*models.Organization, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) GetResource(ctx context.Context, id int) (*models.Resource, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateResource(ctx context.Context, orgID int, resourceName string) (*models.Resource, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteResource(ctx context.Context, id int) (*models.Resource, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateResource(ctx context.Context, id int, resourceName string) (*models.Resource, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) GetSchema(ctx context.Context, id int) (*models.Schema, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) GetSchemas(ctx context.Context, resourceID int) ([]*models.Schema, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteSchema(ctx context.Context, id int) (*models.Schema, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSchema(ctx context.Context, name string) (*models.Schema, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSchemaWithFile(ctx context.Context, name string, file graphql.Upload) (*models.Schema, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSchemaName(ctx context.Context, id int, name string) (*models.Schema, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSchemaFile(ctx context.Context, id int, file graphql.Upload) (*models.Schema, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
