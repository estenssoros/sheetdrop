package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/estenssoros/sheetdrop/graph/generated"
	"github.com/estenssoros/sheetdrop/models"
)

func (r *headerResolver) Schema(ctx context.Context, obj *models.Header) (*models.Schema, error) {
	return r.SchemaByID(obj.SchemaID)
}

func (r *organizationResolver) User(ctx context.Context, obj *models.Organization) ([]*models.User, error) {
	return r.OrganizationUsers(obj)
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return r.ListUsers()
}

func (r *resourceResolver) Organization(ctx context.Context, obj *models.Resource) (*models.Organization, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *resourceResolver) AuthToken(ctx context.Context, obj *models.Resource) (string, error) {
	return obj.AuthToken.String(), nil
}

func (r *resourceResolver) Schemas(ctx context.Context, obj *models.Resource) ([]*models.Schema, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *schemaResolver) UUID(ctx context.Context, obj *models.Schema) (string, error) {
	return obj.UUID.String(), nil
}

func (r *schemaResolver) Resource(ctx context.Context, obj *models.Schema) (*models.Resource, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Organizations(ctx context.Context, obj *models.User) ([]*models.Organization, error) {
	return r.UserOrganizations(obj)
}

// Header returns generated.HeaderResolver implementation.
func (r *Resolver) Header() generated.HeaderResolver { return &headerResolver{r} }

// Organization returns generated.OrganizationResolver implementation.
func (r *Resolver) Organization() generated.OrganizationResolver { return &organizationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Resource returns generated.ResourceResolver implementation.
func (r *Resolver) Resource() generated.ResourceResolver { return &resourceResolver{r} }

// Schema returns generated.SchemaResolver implementation.
func (r *Resolver) Schema() generated.SchemaResolver { return &schemaResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type headerResolver struct{ *Resolver }
type organizationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type resourceResolver struct{ *Resolver }
type schemaResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
