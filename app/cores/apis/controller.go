package apis

import (
	"example/app/cores/extras"
	"gorm.io/gorm"
)

// ControllerImpl interface for a defined router as extras.RouterImpl and database with gorm.DB
type ControllerImpl interface {
	Router() extras.RouterImpl
	DB() *gorm.DB
}

type Controller struct {
	router extras.RouterImpl
	db     *gorm.DB
}

func NewController(router extras.RouterImpl, db *gorm.DB) ControllerImpl {
	return &Controller{
		router: router,
		db:     db,
	}
}

func (c *Controller) Router() extras.RouterImpl {
	return c.router
}

func (c *Controller) DB() *gorm.DB {
	return c.db
}
