package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"math/big"
	"net/url"
	"nokowebapi/console"
	"nokowebapi/cores"
	"nokowebapi/globals"
	"time"
)

func main() {
	var err error
	cores.KeepVoid(err)

	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigFile("app/config.yaml")

	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	//cores.ApplyMainFunc(app.Main)
	//return

	console.Dir(globals.Globals().GetJwtConfig())
	console.Log(fmt.Sprintf("JWT config expires in = %f / hour \n", cores.Unwrap(time.ParseDuration(globals.Globals().GetJwtConfig().ExpiresIn)).Hours()))
	return

	console.Warn("This is warning message.", zap.Int("EXIT_CODE", cores.ExitCodeFailure))
	//console.Fatal("This is error message.", zap.Int("EXIT_CODE", cores.ExitCodeFailure))

	f1, _, _ := big.ParseFloat("540.235273912", 10, 0, big.ToNearestEven)
	f2, _, _ := big.ParseFloat("728.825273912", 10, 0, big.ToNearestEven)
	f1.Add(f1, f2)
	fmt.Println(f1.String())

	console.Log("Try fetch url.")

	cores.TryFetchUrlWaitForAlive(cores.Unwrap(url.Parse("http://localhost")), 12, time.Second)

	buff := cores.Unwrap(cores.HashPassword("Admin@1234"))
	fmt.Println("Compare Password =", cores.CompareHashPassword(buff, "Admin@1234"))
	fmt.Println("Password =", buff)

	console.Log("Done.")

	cores.Unwrap(cores.SetWorkingDir(func(workingDir cores.WorkingDirImpl) {
		console.Log(workingDir.GetSourceRootDir())
		console.Log(cores.Unwrap(cores.GetSourceRootDir()))
		console.Log(workingDir.GetCurrentWorkingDir())
		console.Log(cores.Unwrap(cores.GetCurrentWorkingDir()))
	}))

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
	//password := cores.Base64Encode(hash[:])
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
