package dataloader

import (
	"context"
	"time"

	"github.com/estenssoros/sheetdrop/controllers"
	"github.com/estenssoros/sheetdrop/models"
	"github.com/labstack/echo/v4"
)

const loadersKey = "dataloaders"

// Loaders wrapper for all data loaders
type Loaders struct {
	UserByID                UserLoader
	HeaderByID              HeaderLoader
	SchemaHeadersBySchemaID SchemaHeaderLoader
	SchemaByID              SchemaLoader
	OrganizationByID        OrganizationLoader
	ResourceSchemaCountByID ResourceSchemaCount
}

// Middleware injects dataloaders onto context
func Middleware(ctl *controllers.Controller) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.WithValue(context.Background(), loadersKey, &Loaders{
				UserByID: UserLoader{
					maxBatch: 100,
					wait:     1 * time.Millisecond,
					fetch: func(ids []int) ([]*models.User, []error) {
						return ctl.UsersByIds(ids)
					},
				},
				HeaderByID: HeaderLoader{
					maxBatch: 100,
					wait:     1 * time.Millisecond,
					fetch: func(ids []int) ([]*models.Header, []error) {
						return ctl.HeadersByIDs(ids)
					},
				},
				SchemaHeadersBySchemaID: SchemaHeaderLoader{
					maxBatch: 100,
					wait:     1 * time.Millisecond,
					fetch: func(ids []int) ([][]*models.Header, []error) {
						return ctl.HeadersBySchemaIDs(ids)
					},
				},
				SchemaByID: SchemaLoader{
					maxBatch: 100,
					wait:     1 * time.Millisecond,
					fetch: func(ids []int) ([]*models.Schema, []error) {
						return ctl.SchemasByIDs(ids)
					},
				},
				OrganizationByID: OrganizationLoader{
					maxBatch: 100,
					wait:     1 * time.Millisecond,
					fetch: func(ids []int) ([]*models.Organization, []error) {
						return ctl.OrganizationsByIDs(ids)
					},
				},
				ResourceSchemaCountByID: ResourceSchemaCount{
					maxBatch: 100,
					wait:     1 * time.Millisecond,
					fetch: func(ids []int) ([]int, []error) {
						return ctl.ResourceSchemaCountByID(ids)
					},
				},
			})
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}

// For retrieves dataloaders from context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
