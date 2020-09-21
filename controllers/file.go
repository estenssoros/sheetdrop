package controllers

import (
	"io/ioutil"
	"mime/multipart"
	"path/filepath"

	"github.com/estenssoros/sheetdrop/constants"
	"github.com/estenssoros/sheetdrop/internal/common"
	"github.com/estenssoros/sheetdrop/internal/helpers"
	"github.com/estenssoros/sheetdrop/internal/process"
	"github.com/estenssoros/sheetdrop/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type File interface {
	ProcessFile(*ProcessFileInput) (schema *models.Schema, err error)
}

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

func (c *Controller) ProcessFile(input *ProcessFileInput) (schema *models.Schema, err error) {
	if err := c.Validate(input); err != nil {
		return nil, errors.Wrap(err, "input.Validate")
	}
	if *input.SchemaID != 0 {
		schema, err = c.SchemaByID(*input.SchemaID)
		if err != nil {
			return nil, errors.Wrap(err, "SchemaByID")
		}
		schema.Name = input.Name
	} else {
		schema = &models.Schema{
			APIID: *input.APIID,
			Name:  input.Name,
		}
	}

	schema.SourceURI = input.FileName

	var processor = func() error { return nil }
	data, err := ioutil.ReadAll(input.File)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll")
	}

	switch *input.Extension {
	case constants.ExtensionExcel:
		processor = func() error {
			return process.Excel(schema, data)
		}
	case constants.ExtensionCSV:
		processor = func() error {
			return process.CSV(schema, data)
		}
	default:
		return nil, errors.Wrap(common.ErrUnknownExtension, *input.Extension)
	}
	if err := processor(); err != nil {
		return nil, errors.Wrap(err, *input.Extension)
	}
	headerSet, err := c.GetSchemaHeadersSet(schema)
	{
		headers := headerSet.ToCreate(schema.Headers)
		if len(headers) > 0 {
			if err := c.db.Create(headers).Error; err != nil {
				return nil, errors.Wrap(err, "createHeaders")
			}
		}
	}
	{
		headers := headerSet.ToUpdate(schema.Headers)
		if len(headers) > 0 {
			if err := c.db.Save(headers).Error; err != nil {
				return nil, errors.Wrap(err, "createHeaders")
			}
		}
	}
	{
		headers := headerSet.ToDelete(schema.Headers)
		if len(headers) > 0 {
			if err := c.db.Delete(headers).Error; err != nil {
				return nil, errors.Wrap(err, "createHeaders")
			}
		}
	}

	if err := c.DB().Save(schema).Error; err != nil {
		return nil, errors.Wrap(err, "save schema")
	}
	return schema, nil
}
