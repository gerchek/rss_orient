package gormclient

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewClient
func NewClient() (client *gorm.DB, err error) {

	err = godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")
	//
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_password, db_host, db_name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalln(err)
	}

	// db.AutoMigrate(&model.Post{})

	return db, err
}

//CloseDatabaseConnection method is closing a connection between app and database
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
