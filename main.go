package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
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

func addFlashMessage(session *sessions.Session, status string, message string, c *gin.Context) {
	session.AddFlash(status, "status")
	session.AddFlash(message, "message")
	session.Save(c.Request, c.Writer)
}
