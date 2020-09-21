package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/estenssoros/sheetdrop/graph/generated"
	"github.com/estenssoros/sheetdrop/models"
)

func (r *aPIResolver) Organization(ctx context.Context, obj *models.API) (*models.Organization, error) {
	return r.OrganizationByID(obj.OrganizationID)
}

func (r *aPIResolver) AuthToken(ctx context.Context, obj *models.API) (string, error) {
	return obj.AuthToken.String(), nil
}

func (r *aPIResolver) Schemas(ctx context.Context, obj *models.API) ([]*models.Schema, error) {
	return r.APISChemas(obj)
}

func (r *headerResolver) Schema(ctx context.Context, obj *models.Header) (*models.Schema, error) {
	return r.SchemaByID(obj.SchemaID)
}

func (r *organizationResolver) User(ctx context.Context, obj *models.Organization) ([]*models.User, error) {
	return r.OrganizationUsers(obj)
}

func (r *schemaResolver) UUID(ctx context.Context, obj *models.Schema) (string, error) {
	return obj.UUID.String(), nil
}

func (r *schemaResolver) API(ctx context.Context, obj *models.Schema) (*models.API, error) {
	return r.APIByID(obj.APIID)
}

func (r *userResolver) Organizations(ctx context.Context, obj *models.User) ([]*models.Organization, error) {
	return r.UserOrganizations(obj)
}

// API returns generated.APIResolver implementation.
func (r *Resolver) API() generated.APIResolver { return &aPIResolver{r} }

// Header returns generated.HeaderResolver implementation.
func (r *Resolver) Header() generated.HeaderResolver { return &headerResolver{r} }

// Organization returns generated.OrganizationResolver implementation.
func (r *Resolver) Organization() generated.OrganizationResolver { return &organizationResolver{r} }

// Schema returns generated.SchemaResolver implementation.
func (r *Resolver) Schema() generated.SchemaResolver { return &schemaResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type aPIResolver struct{ *Resolver }
type headerResolver struct{ *Resolver }
type organizationResolver struct{ *Resolver }
type schemaResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
