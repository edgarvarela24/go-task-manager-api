package task_manager_api

import (
	"net/http"
	"strings"
)

func SetupRoutes() {
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", taskHandler)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createTask(w, r)
	case http.MethodGet:
		getAllTasks(w)
	default:
		http.NotFound(w, r)
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	taskId := strings.TrimPrefix(r.URL.Path, "/tasks/")

	switch r.Method {
	case http.MethodGet:
		getTaskById(w, taskId)
	case http.MethodPut:
		updateTask(w, r)
	case http.MethodDelete:
		deleteTask(w, taskId)
	default:
		http.NotFound(w, r)
	}
}
