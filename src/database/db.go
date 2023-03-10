package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	name := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	host := os.Getenv("DB_HOST")

	connection := fmt.Sprintf("dbname=%s user=%s password=%s port=%s host=%s sslmode=disable",
		name, user, password, port, host,
	)

	DB, err = gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		_ = fmt.Errorf("database not connect")
	}
}
