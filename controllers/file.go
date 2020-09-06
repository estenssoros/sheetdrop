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
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ProcessFileInput struct {
	User      string
	SchemaID  *int    `form:"id"`
	APIID     *int    `json:"api_id" form:"api_id"`
	Name      *string `form:"name"`
	FileName  string
	Extension *string
	File      multipart.File
	NewSchema bool
}

func (input *ProcessFileInput) Validate(db *gorm.DB) error {
	if input.Extension == nil {
		input.Extension = helpers.StringPtr(filepath.Ext(input.FileName))
	}
	if !common.ValidExtension(*input.Extension) {
		return errors.Errorf("not valid extension: %s", *input.Extension)
	}
	if input.APIID == nil {
		return errors.New("schema missing api_id")
	}
	if !input.NewSchema && input.SchemaID == nil {
		return errors.New("schema missing id")
	}
	if input.Name == nil {
		return errors.New("schema missing name")
	}
	return nil
}
func (input *ProcessFileInput) AddDataToModel(schema *models.Schema) {
	if !input.NewSchema {
		schema.ID = *input.SchemaID
	}
	schema.Name = input.Name
	schema.APIID = *input.APIID
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
	input.AddDataToModel(schema)

	if err := db.Save(schema).Error; err != nil {
		return nil, errors.Wrap(err, "save schema")
	}
	return &responses.ProcessFile{
		Schema: schema,
	}, nil
}
