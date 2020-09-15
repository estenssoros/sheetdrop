package server

import (
	"github.com/labstack/echo/v4"
)

func routes(e *echo.Group) {
	e.GET("/apis", getAPIsHandler)
	e.GET("/api/:id", getAPIHandler)
	e.POST("/api", createAPIHandler)
	e.DELETE("/api", deleteAPIHandler)
	e.PATCH("/api", updateAPIHandler)

	e.GET("/orgs", getOrgsHandler)

	e.GET("/schema/:apiID", getSchemaHandler)
	e.PATCH("/schema", updateSchemaHandler)
	e.DELETE("/schema", deleteSchemaHandler)
	e.PATCH("/file-upload", schemaFilePatchHandler)
	e.POST("/file-upload", schemaFileUploadHandler)
}
