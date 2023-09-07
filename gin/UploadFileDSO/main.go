package main

import (
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", func(ctx *gin.Context) {
		if file, err := ctx.FormFile("file"); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Something went wrong! " + err.Error(),
			})
		} else {
			currentDir, _ := os.Getwd()
			filePath := path.Join(currentDir, file.Filename)

			ctx.SaveUploadedFile(file, filePath)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Saved file to path: " + filePath,
			})
		}
	})

	router.POST("/multi-upload", func(ctx *gin.Context) {
		if form, err := ctx.MultipartForm(); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Something went wrong! " + err.Error(),
			})
		} else {
			currentDir, _ := os.Getwd()
			files := form.File["upload[]"]

			for _, file := range files {
				log.Println(file.Filename)
				filePath := path.Join(currentDir, file.Filename)
				ctx.SaveUploadedFile(file, filePath)
			}

			ctx.JSON(http.StatusOK, gin.H{
				"message": "Saved file to path: " + currentDir,
			})
		}
	})

	router.Run()
}
