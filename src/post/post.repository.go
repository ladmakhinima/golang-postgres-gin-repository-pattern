package post

import (
	"github.com/ladmakhinima/golang-postgres-gin-repository-pattern/src/database"
	"github.com/ladmakhinima/golang-postgres-gin-repository-pattern/src/repo"
)

var PostRepo repo.IRepo[Post]

func InitialPostRepository() {
	PostRepo = repo.InitialRepo[Post](database.DB)
}
