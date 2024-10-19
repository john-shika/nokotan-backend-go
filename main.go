package main

import (
	"example/app/cores"
	"example/app/globals"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	var err error
	cores.KeepVoid(err)

	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	globals.GlobalJwtConfigInit()

	fmt.Println(cores.ShikaYamlEncodePreview(globals.ConfigDefaults))

	//var db *gorm.DB
	//cores.KeepVoid(db)
	//
	//e := echo.New()
	//config := &gorm.Config{}
	//
	//cores.NoErr(cores.EnsureDirAndFile("migrations/dev.db"))
	//
	//if db, err = gorm.Open(sqlite.Open("migrations/dev.db"), config); err != nil {
	//	panic("failed to connect database")
	//}
	//
	//if err = db.AutoMigrate(&models.User{}); err != nil {
	//	panic(fmt.Errorf("failed to migrate database: %w", err))
	//}
	//if err = db.AutoMigrate(&models.Session{}); err != nil {
	//	panic(fmt.Errorf("failed to migrate database: %w", err))
	//}
	//if err = db.AutoMigrate(&models.Todo{}); err != nil {
	//	panic(fmt.Errorf("failed to migrate database: %w", err))
	//}
	//
	//hash := sha256.Sum256([]byte("Admin@1234"))
	//password := cores.NewBase64EncodeToString(hash[:])
	//db.Create(&models.User{
	//	UUID:     cores.NewUuid(),
	//	Email:    cores.NewNullString("admin@localhost"),
	//	Username: "admin",
	//	Password: password,
	//	Role:     "admin,user",
	//})
	//db.Create(&models.Todo{
	//	Title:       "Todo 1",
	//	Description: "Todo 1 description",
	//})
	//
	//router := e.Group("/api/v1")
	//cores.KeepVoid(router)
	//
	//e.HTTPErrorHandler = func(err error, c echo.Context) {
	//	// TODO
	//	fmt.Println(err)
	//	e.DefaultHTTPErrorHandler(err, c)
	//}
	//
	//if err = e.Start(":8080"); err != nil {
	//	e.Logger.Fatal(err)
	//}
}
