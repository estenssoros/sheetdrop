package server

import (
	"net/http"
	"strconv"

	"github.com/estenssoros/sheetdrop/constants"
	"github.com/estenssoros/sheetdrop/controllers"
	"github.com/estenssoros/sheetdrop/internal/models"
	"github.com/estenssoros/sheetdrop/responses"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func routes(e *echo.Group) {
	e.GET("/api", getAPIHandler)
	e.POST("/api", createAPIHandler)
	e.DELETE("/api", deleteAPIHandler)
	e.PATCH("/api", updateAPIHandler)

	e.GET("/schema/:apiID", getSchemaHandler)
	e.PATCH("/schema", updateSchemaHandler)
	e.DELETE("/schema", deleteSchemaHandler)
	e.PATCH("/file-upload", schemaFilePatchHandler)
	e.POST("/file-upload", schemaFileUploadHandler)
}

func createAPIHandler(c echo.Context) error {
	api := &models.API{}
	if err := c.Bind(api); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	db := c.Get(constants.ContextDB).(*gorm.DB)
	user, err := controllers.GetUserByName(db, c.Get(constants.ContextUserName).(string))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	newAPI := &models.API{
		UserID: user.ID,
		Name:   api.Name,
	}
	if err := db.Create(&newAPI).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, newAPI)
}

func deleteAPIHandler(c echo.Context) error {
	api := &models.API{}
	if err := c.Bind(api); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	db := c.Get(constants.ContextDB).(*gorm.DB)
	user, err := controllers.GetUserByID(db, api.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	userName := c.Get(constants.ContextUserName).(string)
	if user.UserName != userName {
		return c.JSON(http.StatusForbidden, "username not valid")
	}
	if err := db.Delete(&api).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func updateAPIHandler(c echo.Context) error {
	api := &models.API{}
	if err := c.Bind(api); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	db := c.Get(constants.ContextDB).(*gorm.DB)
	user, err := controllers.GetUserByID(db, api.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	userName := c.Get(constants.ContextUserName).(string)
	if user.UserName != userName {
		return c.JSON(http.StatusForbidden, "username not valid")
	}
	if err := db.Save(api).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, api)
}

func getAPIHandler(c echo.Context) error {
	userName := c.Get(constants.ContextUserName).(string)
	db := c.Get(constants.ContextDB).(*gorm.DB)
	user, err := controllers.GetOrCreateUserByName(db, userName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	apis, err := controllers.GetUserAPIs(db, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if len(apis) == 0 {
		api, err := controllers.CreateAPIForUser(db, user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		apis = append(apis, api)
	}
	return c.JSON(http.StatusOK, apis)
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

func getSchemaHandler(c echo.Context) error {
	apiID, err := strconv.Atoi(c.Param("apiID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}
	if apiID == 0 {
		return c.JSON(http.StatusBadRequest, "no id sent")
	}
	db := c.Get(constants.ContextDB).(*gorm.DB)
	user, err := controllers.GetUserFromAPIID(db, apiID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if user.UserName != c.Get(constants.ContextUserName).(string) {
		return c.JSON(http.StatusForbidden, "user names do not match")
	}
	schemas, err := controllers.GetSchemasForAPI(db, &models.API{ID: apiID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if len(schemas) > 0 {
		return c.JSON(http.StatusOK, schemas)
	}
	schema, err := controllers.CreateSchemaForAPI(db, &models.API{ID: apiID})
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
