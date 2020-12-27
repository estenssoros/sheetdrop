package dataloader

import (
	"context"
	"time"

	"github.com/estenssoros/sheetdrop/controllers"
	"github.com/estenssoros/sheetdrop/models"
	"github.com/labstack/echo/v4"
)

const loadersKey = "dataloaders"

type Loaders struct {
	UserById                UserLoader
	HeaderByID              HeaderLoader
	SchemaHeadersBySchemaID SchemaHeaderLoader
	SchemaByID              SchemaLoader
}

func Middleware(ctl *controllers.Controller) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.WithValue(context.Background(), loadersKey, &Loaders{
				UserById: UserLoader{
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
			})
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
