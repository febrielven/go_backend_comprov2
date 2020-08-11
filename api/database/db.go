package database

import (
	"fmt"
	"log"

	"github.com/febrielven/go_backend_comprov2/api/database/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func Connect() *gorm.DB {
	var err error
	err = godotenv.Load()

	if err != nil {
		log.Fatalf("Error getting env, not comming througn %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	db, err := gorm.Open("postgres", config.BuildDSN())

	if err != nil {
		log.Fatal("err connect db: ", err)
	}

	return db

}
