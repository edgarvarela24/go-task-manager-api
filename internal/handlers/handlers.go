package handlers

import (
	"net/http"
	"strings"

	"github.com/edgarvarela24/task-manager-api/internal/handlers/task-handlers"
	"github.com/edgarvarela24/task-manager-api/internal/handlers/user-handlers"
	"github.com/jinzhu/gorm"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	switch r.Method {
	case http.MethodPost:
		user_handlers.LoginUser(w, r, db)
	default:
		http.NotFound(w, r)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	switch r.Method {
	case http.MethodPost:
		user_handlers.RegisterUser(w, r, db)
	default:
		http.NotFound(w, r)
	}
}

func TasksHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	switch r.Method {
	case http.MethodPost:
		task_handlers.CreateTask(w, r, db)
	case http.MethodGet:
		task_handlers.GetAllTasks(w, db)
	default:
		http.NotFound(w, r)
	}
}

func TaskHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	taskId := strings.TrimPrefix(r.URL.Path, "/tasks/")

	switch r.Method {
	case http.MethodGet:
		task_handlers.GetTaskById(w, taskId, db)
	case http.MethodPut:
		task_handlers.UpdateTask(w, r, db)
	case http.MethodDelete:
		task_handlers.DeleteTask(w, taskId, db)
	default:
		http.NotFound(w, r)
	}
}
