package models

import (
	"fmt"
	"log"

	"github.com/Anirudh4583/go-gin-template/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Test struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);unique_index" json:"username"`
	Password string `gorm:"type:varchar(100)" json:"password"`
}

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

	db.Debug().AutoMigrate(&Test{})

	fmt.Println("[info] Connected to the database!")
}
