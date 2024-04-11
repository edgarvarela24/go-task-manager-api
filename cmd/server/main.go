package main

import (
	"log"
	"net/http"

	"github.com/edgarvarela24/task-manager-api/internal/database"
	"github.com/edgarvarela24/task-manager-api/internal/routes"
)

func main() {
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	router := http.NewServeMux()
	routes.SetupRoutes(router, db)

	log.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
