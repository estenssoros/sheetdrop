package server

import (
	"net/http"

	"github.com/estenssoros/sheetdrop/controllers"
	"github.com/labstack/echo"
)

func routes(e *echo.Group) {
	e.POST("/file-upload", fileUploadHandler)
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
	resp, err := controllers.ProcessFile(&controllers.ProcessFileInput{
		FileName: file.Filename,
		File:     multiPart,
		Size:     file.Size,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}
