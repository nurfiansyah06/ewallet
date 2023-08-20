package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Print("cant find the .env")
	}

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	Db_Url := "host=" + DbHost + " user=" + DbUser + " password=" + DbPassword + " dbname=" + DbName + " port=" + DbPort + " sslmode=disable TimeZone=Asia/Shanghai"
	dsn := Db_Url
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	return db

		
}