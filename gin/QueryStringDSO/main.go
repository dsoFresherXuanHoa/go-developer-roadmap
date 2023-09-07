package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		name := ctx.Query("username")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": name,
		})
	})

	router.Run()
}
