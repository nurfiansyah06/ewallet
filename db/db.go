package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {

	godotenv.Load()

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")

	Db_Url := "host=" + DbHost + " user=" + DbUser + " password=" + DbPassword + " dbname=db_go port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := Db_Url
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	return db

		
}