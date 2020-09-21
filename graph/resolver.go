package graph

import (
	"github.com/estenssoros/sheetdrop/controllers"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	controllers.Interface
	sema chan struct{}
}

func NewResolver(db *gorm.DB) *Resolver {
	return &Resolver{
		Interface: controllers.New(db),
		sema:      make(chan struct{}, 10),
	}
}
