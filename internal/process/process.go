package process

import "github.com/estenssoros/sheetdrop/models"

type Result struct {
	StartRow    int
	StartColumn int
	Headers     []*models.Header
}
