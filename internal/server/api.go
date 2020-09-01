package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/estenssoros/sheetdrop/constants"
	"github.com/estenssoros/sheetdrop/controllers"
	"github.com/estenssoros/sheetdrop/internal/models"
	"github.com/estenssoros/sheetdrop/responses"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func routes(e *echo.Group) {
	e.GET("/api", getAPIHandler)
	e.POST("/api", createAPIHandler)
	e.DELETE("/api", deleteAPIHandler)
	e.PATCH("/api", updateAPIHandler)

	e.GET("/schema/:apiID", getSchemaHandler)
	e.PATCH("/schema", updateSchemaHandler)
	e.POST("/file-upload", fileUploadHandler)
	// e.POST("/login", loginHandler,
	// 	middleware.CSRFWithConfig(middleware.CSRFConfig{
	// 		TokenLookup: "form:csrf",
	// 	}),
	// )
	// e.GET("/logout", logoutHandler, middleware.CSRFWithConfig(middleware.CSRFConfig{
	// 	TokenLookup: "form:csrf",
	// }))
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

func fileUploadHandler(c echo.Context) error {
	fmt.Println(c.Request().Header)
	file, err := c.FormFile("file")
	if err != nil {
		return responses.Error(c, http.StatusBadRequest, errors.Wrap(err, "c.FormFile"))
	}
	multiPart, err := file.Open()
	if err != nil {
		return responses.Error(c, http.StatusInternalServerError, errors.Wrap(err, "file.Open"))
	}

	resp, err := controllers.ProcessFile(c.Get(constants.ContextDB).(*gorm.DB), &controllers.ProcessFileInput{
		FileName: file.Filename,
		File:     multiPart,
		User:     c.Get(constants.ContextUserName).(string),
	})
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
	db := c.Get(constants.ContextDB).(*gorm.DB)
	schema, err := controllers.UpdateSchema(db, input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, schema)
}

// func loginHandler(c echo.Context) error {
// 	input := &controllers.LoginInput{}
// 	if err := c.Bind(input); err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}
// 	if err := controllers.Login(c.Get(constants.ContextDB).(*gorm.DB), input); err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}
// 	sess, _ := session.Get("session", c)
// 	sess.Values[constants.SessionLoggedIn] = true
// 	sess.Values[constants.SessionUser] = *input.UserName
// 	sess.Save(c.Request(), c.Response())
// 	loggedIn, _ := sess.Values[constants.SessionLoggedIn].(bool)
// 	return c.Render(http.StatusOK, "main", struct{ LoggedIn bool }{loggedIn})
// }

// func logoutHandler(c echo.Context) error {
// 	sess, _ := session.Get("session", c)
// 	sess.Values[constants.SessionLoggedIn] = true
// 	delete(sess.Values, constants.SessionUser)
// 	sess.Save(c.Request(), c.Response())
// 	return c.Render(http.StatusOK, "login", c.Get("csrf").(string))
// }
