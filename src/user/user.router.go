package user

import "github.com/gin-gonic/gin"

func InitialUserRoute(server *gin.Engine) {
	usersRoute := server.Group("/api/users")
	usersRoute.POST("/", SaveAction)
}
