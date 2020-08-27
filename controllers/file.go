package controllers

import (
	"io/ioutil"
	"mime/multipart"
	"path/filepath"

	"github.com/estenssoros/sheetdrop/constants"
	"github.com/estenssoros/sheetdrop/internal/common"
	"github.com/estenssoros/sheetdrop/internal/helpers"
	"github.com/estenssoros/sheetdrop/internal/models"
	"github.com/estenssoros/sheetdrop/internal/process"
	"github.com/estenssoros/sheetdrop/responses"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type ProcessFileInput struct {
	User      string
	APIID     *int `json:"api_id"`
	FileName  string
	Extension *string
	File      multipart.File
}

func (input *ProcessFileInput) Validate(db *gorm.DB) error {
	if input.Extension == nil {
		input.Extension = helpers.StringPtr(filepath.Ext(input.FileName))
	}
	if input.APIID != nil {
		return nil
	}
	user, err := GetUserByName(db, input.User)
	if err != nil {
		return errors.Wrap(err, "GetUserByName")
	}
	if err := db.Create(&models.API{UserID: user.ID}).Error; err != nil {
		return errors.Wrap(err, "create api")
	}
	return nil
}

func ProcessFile(db *gorm.DB, input *ProcessFileInput) (*responses.ProcessFile, error) {

	if err := input.Validate(db); err != nil {
		return nil, errors.Wrap(err, "input.Validate")
	}

	var processor = func() (*models.Schema, error) { return nil, nil }
	data, err := ioutil.ReadAll(input.File)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll")
	}
	switch *input.Extension {
	case constants.ExtensionExcel:
		processor = func() (*models.Schema, error) {
			return process.Excel(data)
		}
	case constants.ExtensionCSV:
		processor = func() (*models.Schema, error) {
			return process.CSV(data)
		}
	default:
		return nil, errors.Wrap(common.ErrUnknownExtension, *input.Extension)
	}
	schema, err := processor()
	if err != nil {
		return nil, errors.Wrap(err, *input.Extension)
	}
	return &responses.ProcessFile{
		Schema: schema,
	}, nil
}
