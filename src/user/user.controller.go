package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveAction(context *gin.Context) {
	var body User
	context.ShouldBind(&body)
	UserRepo.Save(&body)
	context.JSON(http.StatusCreated, gin.H{
		"data": body,
	})
}
