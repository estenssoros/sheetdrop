package graph

import (
	"github.com/estenssoros/sheetdrop/controllers"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	*controllers.Controller
	sema chan struct{}
}

func NewResolver(db *gorm.DB) *Resolver {
	return &Resolver{
		Controller: controllers.New(db),
		sema:       make(chan struct{}, 10),
	}
}
