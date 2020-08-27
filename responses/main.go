package responses

import "github.com/estenssoros/sheetdrop/internal/models"

type Main struct {
	LoggedIn bool
	APIs     []*models.API
}
