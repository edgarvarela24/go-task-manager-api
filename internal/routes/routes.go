package routes

import (
	"github.com/edgarvarela24/task-manager-api/internal/middleware"
	"net/http"

	"github.com/edgarvarela24/task-manager-api/internal/handlers"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(router *http.ServeMux, db *gorm.DB) {
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, db)
	})
	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterHandler(w, r, db)
	})

	// Following routes require JWT authentication
	router.HandleFunc("/tasks", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.TasksHandler(w, r, db)
	}))
	router.HandleFunc("/tasks/", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.TaskHandler(w, r, db)
	}))
}
