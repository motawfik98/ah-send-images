package main

import (
	"./configurations"
	"./handlers"
	"github.com/labstack/echo"
)

func main() {
	db, _ := configurations.InitDB()

	e := echo.New()
	myDb := handlers.MyDB{GormDB: db}
	handlers.InitializeRoutes(e, &myDb)

	e.Logger.Fatal(e.Start(":8080"))
}
