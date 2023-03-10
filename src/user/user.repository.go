package user

import (
	"github.com/ladmakhinima/golang-postgres-gin-repository-pattern/src/database"
	"github.com/ladmakhinima/golang-postgres-gin-repository-pattern/src/repo"
)

var UserRepo repo.IRepo[User]

func InitialUserRepository() {
	UserRepo = repo.InitialRepo[User](database.DB)
}
