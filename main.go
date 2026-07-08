package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/up", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "App is Healthy",
		})
	})

	router.Run()
}
