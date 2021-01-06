package server

// func schemaFilePatchHandler(c echo.Context) error {
// 	input := &controllers.ProcessFileInput{}
// 	if err := c.Bind(input); err != nil {
// 		return responses.Error(c, http.StatusBadRequest, errors.Wrap(err, "c.Bind"))
// 	}
// 	input.User = usr(c)
// 	return fileUploadHandler(c, input)
// }

// func schemaFileUploadHandler(c echo.Context) error {
// 	input := &controllers.ProcessFileInput{}
// 	if err := c.Bind(input); err != nil {
// 		return responses.Error(c, http.StatusBadRequest, errors.Wrap(err, "c.Bind"))
// 	}
// 	input.User = usr(c)
// 	input.NewSchema = true
// 	return fileUploadHandler(c, input)
// }

// func fileUploadHandler(c echo.Context, input *controllers.ProcessFileInput) error {
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		return responses.Error(c, http.StatusBadRequest, errors.Wrap(err, "c.FormFile"))
// 	}
// 	input.FileName = file.Filename

// 	multiPart, err := file.Open()
// 	if err != nil {
// 		return responses.Error(c, http.StatusInternalServerError, errors.Wrap(err, "file.Open"))
// 	}

// 	input.File = multiPart
// 	resp, err := ctl(c).ProcessFile(input)
// 	if err != nil {
// 		return responses.Error(c, http.StatusInternalServerError, errors.Wrap(err, "controllers.ProcessFile"))
// 	}
// 	return c.JSON(http.StatusOK, resp)
// }
