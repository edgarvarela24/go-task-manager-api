package routes

import (
	"net/http"

	"github.com/edgarvarela24/task-manager-api/internal/handlers"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(router *http.ServeMux, db *gorm.DB) {
	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		handlers.TasksHandler(w, r, db)
	})
	router.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		handlers.TaskHandler(w, r, db)
	})
}
