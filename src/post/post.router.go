package post

import "github.com/gin-gonic/gin"

func InitialPostRoute(server *gin.Engine) {
	postsRoute := server.Group("/api/posts")
	postsRoute.POST("/", SaveAction)
}
