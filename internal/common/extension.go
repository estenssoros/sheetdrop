package common

import (
	"github.com/estenssoros/sheetdrop/constants"
	"github.com/pkg/errors"
)

var ErrUnknownExtension = errors.New("unknown extension")

func ValidExtension(ext string) bool {
	switch ext {
	case constants.ExtensionCSV, constants.ExtensionExcel:
		return true
	}
	return false
}
