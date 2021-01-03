package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/estenssoros/sheetdrop/graph/dataloader"
	"github.com/estenssoros/sheetdrop/graph/generated"
	"github.com/estenssoros/sheetdrop/models"
)

func (r *headerResolver) Schema(ctx context.Context, obj *models.Header) (*models.Schema, error) {
	return dataloader.For(ctx).SchemaByID.Load(obj.ID)
}

func (r *headerResolver) ForeignKeys(ctx context.Context, obj *models.Header) ([]*models.Header, error) {
	return r.HeaderForeignKeys(obj.ID)
}

func (r *organizationResolver) Users(ctx context.Context, obj *models.Organization) ([]*models.User, error) {
	return r.OrganizationUsers(obj.ID)
}

func (r *organizationResolver) Resources(ctx context.Context, obj *models.Organization) ([]*models.Resource, error) {
	return r.OrganizationResources(obj.ID)
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return r.ListUsers()
}

func (r *queryResolver) Organizations(ctx context.Context) ([]*models.Organization, error) {
	return r.ListOrganizations()
}

func (r *queryResolver) Organization(ctx context.Context, id int) (*models.Organization, error) {
	return dataloader.For(ctx).OrganizationByID.Load(id)
}

func (r *queryResolver) Resource(ctx context.Context, id int) (*models.Resource, error) {
	return r.ResourceByID(id)
}

func (r *queryResolver) Resources(ctx context.Context, organizationID int) ([]*models.Resource, error) {
	return r.OrganizationResources(organizationID)
}

func (r *queryResolver) Schemas(ctx context.Context, resourceID int) ([]*models.Schema, error) {
	return r.ListSchemas()
}

func (r *queryResolver) Schema(ctx context.Context, id int) (*models.Schema, error) {
	return dataloader.For(ctx).SchemaByID.Load(id)
}

func (r *queryResolver) Headers(ctx context.Context, schemaID int) ([]*models.Header, error) {
	return r.SchemaHeaders(schemaID)
}

func (r *resourceResolver) Organization(ctx context.Context, obj *models.Resource) (*models.Organization, error) {
	return dataloader.For(ctx).OrganizationByID.Load(obj.OrganizationID)
}

func (r *resourceResolver) AuthToken(ctx context.Context, obj *models.Resource) (string, error) {
	return obj.AuthToken.String(), nil
}

func (r *resourceResolver) SchemaCount(ctx context.Context, obj *models.Resource) (int, error) {
	return dataloader.For(ctx).ResourceSchemaCountByID.Load(obj.ID)
}

func (r *resourceResolver) Schemas(ctx context.Context, obj *models.Resource) ([]*models.Schema, error) {
	return r.SchemasForResource(obj.ID)
}

func (r *schemaResolver) UUID(ctx context.Context, obj *models.Schema) (string, error) {
	return obj.UUID.String(), nil
}

func (r *schemaResolver) Resource(ctx context.Context, obj *models.Schema) (*models.Resource, error) {
	return r.ResourceByID(obj.ResourceID)
}

func (r *schemaResolver) Headers(ctx context.Context, obj *models.Schema) ([]*models.Header, error) {
	return dataloader.For(ctx).SchemaHeadersBySchemaID.Load(obj.ID)
}

func (r *userResolver) Organizations(ctx context.Context, obj *models.User) ([]*models.Organization, error) {
	return r.UserOrganizations(obj.ID)
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
