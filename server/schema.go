package server

import (
	"net/http"
	"strconv"

	"github.com/estenssoros/sheetdrop/controllers"
	"github.com/estenssoros/sheetdrop/models"
	"github.com/estenssoros/sheetdrop/responses"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func getSchemaHandler(c echo.Context) error {
	apiID, err := strconv.Atoi(c.Param("apiID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}
	if apiID == 0 {
		return c.JSON(http.StatusBadRequest, "no id sent")
	}
	user, err := ctl(c).UserFromAPIID(apiID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if user.UserName != usr(c) {
		return c.JSON(http.StatusForbidden, "user names do not match")
	}
	schemas, err := ctl(c).SchemasForAPI(&models.API{
		ID: apiID,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if len(schemas) > 0 {
		if err := ctl(c).SchemaRelations(schemas); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, schemas)
	}
	schema, err := ctl(c).CreateSchemaForAPI(&models.API{
		ID: apiID,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, []*models.Schema{schema})
}

func updateSchemaHandler(c echo.Context) error {
	input := &controllers.UpdateSchemaInput{}
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	schema, err := ctl(c).UpdateSchema(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, schema)
}

func deleteSchemaHandler(c echo.Context) error {
	schema := &models.Schema{}
	if err := c.Bind(schema); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := ctl(c).DeleteSchema(schema); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func schemaFilePatchHandler(c echo.Context) error {
	input := &controllers.ProcessFileInput{}
	if err := c.Bind(input); err != nil {
		return responses.Error(c, http.StatusBadRequest, errors.Wrap(err, "c.Bind"))
	}
	input.User = usr(c)
	return fileUploadHandler(c, input)
}

func schemaFileUploadHandler(c echo.Context) error {
	input := &controllers.ProcessFileInput{}
	if err := c.Bind(input); err != nil {
		return responses.Error(c, http.StatusBadRequest, errors.Wrap(err, "c.Bind"))
	}
	input.User = usr(c)
	input.NewSchema = true
	return fileUploadHandler(c, input)
}

func fileUploadHandler(c echo.Context, input *controllers.ProcessFileInput) error {
	file, err := c.FormFile("file")
	if err != nil {
		return responses.Error(c, http.StatusBadRequest, errors.Wrap(err, "c.FormFile"))
	}
	input.FileName = file.Filename

	multiPart, err := file.Open()
	if err != nil {
		return responses.Error(c, http.StatusInternalServerError, errors.Wrap(err, "file.Open"))
	}

	input.File = multiPart
	resp, err := ctl(c).ProcessFile(input)
	if err != nil {
		return responses.Error(c, http.StatusInternalServerError, errors.Wrap(err, "controllers.ProcessFile"))
	}
	return c.JSON(http.StatusOK, resp)
}