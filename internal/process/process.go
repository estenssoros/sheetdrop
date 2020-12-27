package process

import "github.com/estenssoros/sheetdrop/models"

type Result struct {
	Schema  *models.Schema
	Headers []*models.Header
}
