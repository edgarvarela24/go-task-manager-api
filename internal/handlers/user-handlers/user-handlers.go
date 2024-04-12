package user_handlers

import (
	"encoding/json"
	"errors"
	"github.com/edgarvarela24/task-manager-api/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"regexp"
	"time"
)

func RegisterUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Validate data model
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validateUser(&user, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Verify data model
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve user from db based on provided username and email
	var storedUser models.User
	db.Where("username = ? OR email = ?", user.Username, user.Email).First(&storedUser)
	if storedUser.ID == 0 {
		http.Error(w, "user not found", http.StatusUnauthorized)
		return
	}

	// Compare stored password to provided password
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": storedUser.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	// Get the secret key from environment variable
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		http.Error(w, "missing JWT secret key", http.StatusInternalServerError)
		return
	}

	// Sign token with secret
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	// Return the token in the response
	response := map[string]string{
		"token": tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func validateUser(user *models.User, db *gorm.DB) error {
	// Validate fields are present
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return errors.New("username, email and password are required")
	}

	// Validate email format
	if !isValidEmail(user.Email) {
		return errors.New("invalid email format")
	}

	// Validate password strength
	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	// Check if username or email already exists
	var existingUser models.User
	db.Where("username = ? OR email = ?", user.Username, user.Email).First(&existingUser)
	if existingUser.ID != 0 {
		return errors.New("username or email already exists")
	}

	return nil
}

func isValidEmail(email string) bool {
	// Simple email format validation using regular expression
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
