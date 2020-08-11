package config

import (
	"fmt"
	"os"
)

func BuildDSN() string {
	host := os.Getenv("DB_HOST")

	port := os.Getenv("DB_PORT")

	user := os.Getenv("DB_USER")

	pass := os.Getenv("DB_PASSWORD")

	dbname := os.Getenv("DB_NAME")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, pass)

	return DBURL
}
