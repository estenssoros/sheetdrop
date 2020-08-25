package controllers

import (
	"io"
	"path/filepath"

	"github.com/estenssoros/sheetdrop/constants"
	"github.com/estenssoros/sheetdrop/internal/common"
	"github.com/estenssoros/sheetdrop/internal/helpers"
	"github.com/estenssoros/sheetdrop/internal/models"
	"github.com/estenssoros/sheetdrop/internal/process"
	"github.com/estenssoros/sheetdrop/responses"
	"github.com/pkg/errors"
)

type ProcessFileInput struct {
	FileName  string
	Extension *string
	File      io.ReaderAt
	Size      int64
}

func (input *ProcessFileInput) Validate() error {
	if input.Extension == nil {
		input.Extension = helpers.StringPtr(filepath.Ext(input.FileName))
	}
	return nil
}

func ProcessFile(input *ProcessFileInput) (*responses.ProcessFile, error) {
	if err := input.Validate(); err != nil {
		return nil, errors.Wrap(err, "input.Validate")
	}

	var processor = func() (*models.Schema, error) { return nil, nil }

	switch *input.Extension {
	case constants.ExtensionExcel:
		processor = func() (*models.Schema, error) {
			return process.Excel(input.File, input.Size)
		}
	case constants.ExtensionCSV:
		processor = func() (*models.Schema, error) {
			return process.CSV(input.File)
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
