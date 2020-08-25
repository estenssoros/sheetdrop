package common

import (
	"github.com/estenssoros/sheetdrop/internal/constants"
	"github.com/pkg/errors"
)

var ErrUnknownExtension = errors.New("unknown extension")

func CheckExtension(ext string) error {
	switch ext {
	case constants.ExtensionCSV, constants.ExtensionExcel:
		return nil
	}
	return errors.Wrap(ErrUnknownExtension, ext)
}
