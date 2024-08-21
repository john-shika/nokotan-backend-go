package main

import (
	"example/app/controllers/todo"
	"example/app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var err error
	var db *gorm.DB
	e := echo.New()
	config := &gorm.Config{}
	if db, err = gorm.Open(sqlite.Open("example.sqlite3"), config); err != nil {
		panic("failed to connect database")
	}
	if err = db.AutoMigrate(&models.Todo{}); err != nil {
		return
	}
	db.Create(&models.Todo{
		Title:       "Todo 1",
		Description: "Todo 1 description",
	})
	router := e.Group("/api/v1")
	todo.NewController(router, db)
	e.Logger.Fatal(e.Start(":8080"))
}
