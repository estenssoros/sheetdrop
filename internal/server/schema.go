package server

import (
	"net/http"
	"strconv"

	"github.com/estenssoros/sheetdrop/constants"
	"github.com/estenssoros/sheetdrop/controllers"
	"github.com/estenssoros/sheetdrop/models"
	"github.com/estenssoros/sheetdrop/responses"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func getSchemaHandler(c echo.Context) error {
	apiID, err := strconv.Atoi(c.Param("apiID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}
	if apiID == 0 {
		return c.JSON(http.StatusBadRequest, "no id sent")
	}
	ctl := c.Get(constants.ContextDB).(controllers.Interface)
	user, err := ctl.GetUserFromAPIID(uint(apiID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if user.UserName != c.Get(constants.ContextUserName).(string) {
		return c.JSON(http.StatusForbidden, "user names do not match")
	}
	schemas, err := ctl.GetSchemasForAPI(&models.API{
		Model: gorm.Model{ID: uint(apiID)},
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if len(schemas) > 0 {
		if err := ctl.GetSchemaRelations(schemas); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, schemas)
	}
	schema, err := ctl.CreateSchemaForAPI(&models.API{
		Model: gorm.Model{ID: uint(apiID)},
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
	schema, err := controllers.UpdateSchema(c.Get(constants.ContextDB).(*gorm.DB), input)
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
	if err := controllers.DeleteSchema(c.Get(constants.ContextDB).(*gorm.DB), schema); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func schemaFilePatchHandler(c echo.Context) error {
	input := &controllers.ProcessFileInput{}
	if err := c.Bind(input); err != nil {
		return responses.Error(c, http.StatusBadRequest, errors.Wrap(err, "c.Bind"))
	}
	input.User = c.Get(constants.ContextUserName).(string)
	return fileUploadHandler(c, input)
}

func schemaFileUploadHandler(c echo.Context) error {
	input := &controllers.ProcessFileInput{}
	if err := c.Bind(input); err != nil {
		return responses.Error(c, http.StatusBadRequest, errors.Wrap(err, "c.Bind"))
	}
	input.User = c.Get(constants.ContextUserName).(string)
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

	resp, err := controllers.ProcessFile(c.Get(constants.ContextDB).(*gorm.DB), input)
	if err != nil {
		return responses.Error(c, http.StatusInternalServerError, errors.Wrap(err, "controllers.ProcessFile"))
	}
	return c.JSON(http.StatusOK, resp)
}
