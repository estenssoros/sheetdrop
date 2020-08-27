package server

import (
	"net/http"

	"github.com/estenssoros/sheetdrop/constants"
	"github.com/estenssoros/sheetdrop/controllers"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func routes(e *echo.Group) {
	e.POST("/file-upload", fileUploadHandler)
	e.POST("/login", loginHandler, middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))
	e.GET("/logout", logoutHandler, middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))
}

func fileUploadHandler(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	multiPart, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp, err := controllers.ProcessFile(c.Get(constants.ContextDB).(*gorm.DB), &controllers.ProcessFileInput{
		FileName: file.Filename,
		File:     multiPart,
		User:     c.FormValue("user"),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func loginHandler(c echo.Context) error {
	input := &controllers.LoginInput{}
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := controllers.Login(c.Get(constants.ContextDB).(*gorm.DB), input); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	sess, _ := session.Get("session", c)
	sess.Values[constants.SessionLoggedIn] = true
	sess.Values[constants.SessionUser] = *input.UserName
	sess.Save(c.Request(), c.Response())
	loggedIn, _ := sess.Values[constants.SessionLoggedIn].(bool)
	return c.Render(http.StatusOK, "main", struct{ LoggedIn bool }{loggedIn})
}

func logoutHandler(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Values[constants.SessionLoggedIn] = true
	delete(sess.Values, constants.SessionUser)
	sess.Save(c.Request(), c.Response())
	return c.Render(http.StatusOK, "login", c.Get("csrf").(string))
}
