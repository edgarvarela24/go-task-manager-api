package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/edgarvarela24/task-manager-api/internal/models"
	"github.com/jinzhu/gorm"
)

func TasksHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	switch r.Method {
	case http.MethodPost:
		createTask(w, r, db)
	case http.MethodGet:
		getAllTasks(w, db)
	default:
		http.NotFound(w, r)
	}
}

func TaskHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	taskId := strings.TrimPrefix(r.URL.Path, "/tasks/")

	switch r.Method {
	case http.MethodGet:
		getTaskById(w, taskId, db)
	case http.MethodPut:
		updateTask(w, r, db)
	case http.MethodDelete:
		deleteTask(w, taskId, db)
	default:
		http.NotFound(w, r)
	}
}

func createTask(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Create(&task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getAllTasks(w http.ResponseWriter, db *gorm.DB) {
	var tasks []models.Task
	result := db.Find(&tasks)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getTaskById(w http.ResponseWriter, taskId string, db *gorm.DB) {
	var task models.Task
	result := db.Find(&task, taskId)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func updateTask(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.Save(&task)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteTask(w http.ResponseWriter, taskID string, db *gorm.DB) {
	var task models.Task
	result := db.First(&task, taskID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	result = db.Delete(&task)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
