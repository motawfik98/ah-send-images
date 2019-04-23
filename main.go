package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	initDB()
	router := gin.Default()

	router.Run()
}
