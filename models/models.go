package models

import (
	"fmt"
	"log"

	"github.com/Anirudh4583/go-gin-template/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func Setup() {

	dbURI := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		setting.Config.DBHost,
		setting.Config.DBPort,
		setting.Config.DBName,
		setting.Config.DBUser,
		setting.Config.DBPassword)

	// Connect to the database
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	defer db.Close()

	fmt.Println("Connected to the database!")
}
