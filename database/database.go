package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/galifornia/go-fiber-note-taking/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseInt(p, 10, 32)

	if err != nil {
		log.Println("Bad port in .env")
	}

	connection_url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(connection_url))

	if err != nil {
		panic("Error connecting to the database")
	}

	fmt.Println("Connected to the database!")
}
