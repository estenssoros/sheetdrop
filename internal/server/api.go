package server

import (
	"net/http"
	"path/filepath"

	"github.com/estenssoros/sheetdrop/internal/common"
	"github.com/estenssoros/sheetdrop/internal/constants"
	"github.com/estenssoros/sheetdrop/internal/process"
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
	ext := filepath.Ext(file.Filename)
	if err := common.CheckExtension(ext); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	multiPart, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	switch ext {
	case constants.ExtensionExcel:
		resp, err := process.Excel(multiPart, file.Size)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, resp)
	case constants.ExtensionCSV:
		resp, err := process.CSV(multiPart)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, resp)
	default:
		return c.JSON(http.StatusBadRequest, common.ErrUnknownExtension)
	}
}
