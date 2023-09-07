package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/users/:id", func(ctx *gin.Context) {
		if id, err := strconv.Atoi(ctx.Param("id")); err != nil {
			log.Println("Something went wrong!")
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": id,
			})
		}
	})

	router.Run()
}
