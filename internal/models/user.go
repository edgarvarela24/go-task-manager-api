package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"not null;unique"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
	Role     string `json:"role" gorm:"default:'user'"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}
