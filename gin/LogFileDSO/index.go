package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	logPath, _ := os.Create("gin.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(logPath)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Pong")
	})

	router.Run()
}
