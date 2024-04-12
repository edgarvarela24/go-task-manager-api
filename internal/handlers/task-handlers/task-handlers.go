package task_handlers

import (
	"encoding/json"
	"errors"
	"github.com/edgarvarela24/task-manager-api/internal/models"
	"github.com/jinzhu/gorm"
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
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

func GetAllTasks(w http.ResponseWriter, db *gorm.DB) {
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

func GetTaskById(w http.ResponseWriter, taskId string, db *gorm.DB) {
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

func UpdateTask(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
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

func DeleteTask(w http.ResponseWriter, taskID string, db *gorm.DB) {
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
