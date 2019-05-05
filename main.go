package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	initDB()

	router = gin.Default()
	initializeRoutes()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	router.Run()
}
