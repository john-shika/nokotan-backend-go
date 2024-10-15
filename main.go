package main

import (
	"crypto/sha256"
	"example/app/controllers/todo"
	"example/app/controllers/user"
	"example/app/cores"
	"example/app/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var err error
	var db *gorm.DB
	viper.SetConfigFile("config.yaml")
	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	e := echo.New()
	config := &gorm.Config{}

	if db, err = gorm.Open(sqlite.Open("example.sqlite3"), config); err != nil {
		panic("failed to connect database")
	}
	if err = db.AutoMigrate(&models.User{}); err != nil {
		panic(fmt.Errorf("failed to migrate database: %w", err))
	}
	if err = db.AutoMigrate(&models.Session{}); err != nil {
		panic(fmt.Errorf("failed to migrate database: %w", err))
	}
	if err = db.AutoMigrate(&models.Todo{}); err != nil {
		panic(fmt.Errorf("failed to migrate database: %w", err))
	}
	password := cores.NewBase64EncodeToString(cores.GetArraySliceSize32(sha256.Sum256([]byte("Admin@1234"))))
	db.Create(&models.User{
		UUID:     cores.NewUuid(),
		Email:    cores.NewNullString("admin@localhost"),
		Username: "admin",
		Password: password,
		Role:     "admin,user",
	})
	db.Create(&models.Todo{
		Title:       "Todo 1",
		Description: "Todo 1 description",
	})
	router := e.Group("/api/v1")

	user.NewController(router, db)
	todo.NewController(router, db)

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		// TODO
		fmt.Println(err)
		e.DefaultHTTPErrorHandler(err, c)
	}

	if err = e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
