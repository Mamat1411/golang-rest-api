package controllers

import (
	"golang-rest-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	Service *services.PostService
}

func NewPostController(service *services.PostService) *PostController {
	return &PostController{Service: service}
}

func (controller *PostController) GetPosts(ctx *gin.Context) {
	data, err := controller.Service.GetPosts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to Fetch All Posts",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully Fetch All Posts",
		"data":    data,
	})
}
