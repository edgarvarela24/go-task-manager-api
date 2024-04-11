package database

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func ConnectDatabase() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var db *gorm.DB
	var err error

	// Retry connecting to the database with a delay
	for i := 0; i < 5; i++ {
		db, err = gorm.Open("postgres", connectionString)
		if err == nil {
			break
		}
		fmt.Printf("Failed to connect to the database. Retrying in 5 seconds... (Attempt %d)\n", i+1)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
