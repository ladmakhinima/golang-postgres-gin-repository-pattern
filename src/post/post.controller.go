package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveAction(context *gin.Context) {
	var body Post
	context.ShouldBind(&body)
	PostRepo.Save(&body)
	context.JSON(http.StatusCreated, gin.H{
		"data": body,
	})
}
