package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	initDB()

	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
	initializeRoutes()

	router.Run()
}
