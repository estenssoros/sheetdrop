package server

import (
	"github.com/labstack/echo/v4"
)

func apiRoutes(e *echo.Group) {
	resourceRoutes(e)
	orgRoutes(e)
	schemaRoutes(e)
}

func resourceRoutes(e *echo.Group) {
	e.GET("/resources", getResourcesHandler)
	e.GET("/resource/:id", getResourceHandler)
	e.POST("/resource", createResourceHandler)
	e.DELETE("/resource", deleteResourceHandler)
	e.PATCH("/resource", updateResourceHandler)
}

func orgRoutes(e *echo.Group) {
	e.GET("/orgs", getOrgsHandler)
	e.POST("/org", createOrgHandler)
	e.PATCH("/org", updateOrgHandler)
	e.DELETE("/org", deleteOrgHandler)
}

func schemaRoutes(e *echo.Group) {
	e.GET("/schema/:resourceID", getSchemaHandler)
	e.PATCH("/schema", updateSchemaHandler)
	e.DELETE("/schema", deleteSchemaHandler)
	e.PATCH("/schema/file-upload", schemaFilePatchHandler)
	e.POST("/schema/file-upload", schemaFileUploadHandler)
}

func restRoutes(e *echo.Group) {
	e.GET("/:url", handleRest)
}
