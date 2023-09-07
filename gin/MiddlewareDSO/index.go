package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AccessLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("Has an access!")
		ctx.Next()
	}
}

func main() {
	router := gin.New()

	router.Use(AccessLogger())

	router.Use(gin.Logger())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "GET",
		})
	})

	router.Run()
}
