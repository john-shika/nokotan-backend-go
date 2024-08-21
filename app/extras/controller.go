package extras

import (
	"gorm.io/gorm"
)

// ControllerImpl interface for a defined router as extras.RouterImpl and database with gorm.DB
type ControllerImpl interface {
	Router() RouterImpl
	DB() *gorm.DB
}

type Controller struct {
	router RouterImpl
	db     *gorm.DB
}

func NewController(router RouterImpl, db *gorm.DB) ControllerImpl {
	return &Controller{
		router: router,
		db:     db,
	}
}

func (c *Controller) Router() RouterImpl {
	return c.router
}

func (c *Controller) DB() *gorm.DB {
	return c.db
}
