package task_manager_api

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=admin dbname=dbname password=password sslmode=disable")

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Task{})
	DB = db
}
