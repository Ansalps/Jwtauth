package database

import (
	""
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/"
	"gorm.io/gorm/driver/postgres"
)

func DBinstance() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PostGres := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open(postgres.Open(PostGres), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return db
}

var Client *gorm.DB = DBinstance()

func openTable(client *gorm.DB, tableNmae string) {

}
