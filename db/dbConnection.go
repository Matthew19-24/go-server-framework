package db

import (
	"fmt"
	"os"
	"testServer/utils"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	// Load environment variables
	utils.Log("loading database environment variables")
	var err error
	godotenv.Load()
	dbhost := os.Getenv("MYSQL_HOST")
	dbuser := os.Getenv("MYSQL_USER")
	dbpassword := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbname)

	// Attempt to connect to database and retry x times
	utils.Log("attempting to connect to database")
	retries := 5
	for retries > 0 {
		DB, err = gorm.Open(mysql.Open(connection), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent), 
		})
		if err == nil {
			break
		}
		utils.Log("database connection failed. Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
		retries--
	}
	if err != nil {
		panic("database connection failed after multiple retries")
	}
	utils.Log("db conneted successfully")
}
