package main

import (
	"log"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	currentDir, _ := os.Getwd()
	router.Static("/public", path.Join(currentDir, "/StaticFilesDSO/public"))

	router.GET("/", func(ctx *gin.Context) {
		log.Println(currentDir)
		filePath := path.Join(currentDir, "/StaticFilesDSO/public")
		ctx.File(filePath)
	})

	router.Run()
}
