package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ladmakhinima/golang-postgres-gin-repository-pattern/src/database"
	"github.com/ladmakhinima/golang-postgres-gin-repository-pattern/src/post"
	"github.com/ladmakhinima/golang-postgres-gin-repository-pattern/src/user"
)

func main() {
	loadEnv()
	database.ConnectDB()
	server := gin.Default()
	database.DB.AutoMigrate(&user.User{})
	database.DB.AutoMigrate(&post.Post{})
	post.InitialPostRepository()
	user.InitialUserRepository()
	post.InitialPostRoute(server)
	user.InitialUserRoute(server)
	server.Run(":3000")
}

func loadEnv() {
	godotenv.Load("./.env")
}
