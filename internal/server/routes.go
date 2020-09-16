package server

import (
	"github.com/labstack/echo/v4"
)

func routes(e *echo.Group) {
	apiRoutes(e)
	orgRoutes(e)
	schemaRoutes(e)
}

func apiRoutes(e *echo.Group) {
	e.GET("/apis", getAPIsHandler)
	e.GET("/api/:id", getAPIHandler)
	e.POST("/api", createAPIHandler)
	e.DELETE("/api", deleteAPIHandler)
	e.PATCH("/api", updateAPIHandler)
}

func orgRoutes(e *echo.Group) {
	e.GET("/orgs", getOrgsHandler)
	e.POST("/org", createOrgHandler)
	e.PATCH("/org", updateOrgHandler)
	e.DELETE("/org", deleteOrgHandler)
}

func schemaRoutes(e *echo.Group) {
	e.GET("/schema/:apiID", getSchemaHandler)
	e.PATCH("/schema", updateSchemaHandler)
	e.DELETE("/schema", deleteSchemaHandler)
	e.PATCH("/file-upload", schemaFilePatchHandler)
	e.POST("/file-upload", schemaFileUploadHandler)
}
