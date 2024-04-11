package models

type Task struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	Completed   bool   `json:"completed" gorm:"default:false"`
}
