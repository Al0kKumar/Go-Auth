package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Al0kKumar/Go-Auth.git/db"
	"github.com/Al0kKumar/Go-Auth.git/models"
	"github.com/Al0kKumar/Go-Auth.git/utils"
	"gorm.io/gorm"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Inavlid input", http.StatusBadRequest)
		return
	}

	var existingUser models.User
	if err := db.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		http.Error(w, "Email already registered", http.StatusConflict)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "error hasing password", http.StatusInternalServerError)
		return
	}

	user.Password = hashedPassword

	if err := db.DB.Create(&user).Error; err != nil {
		http.Error(w, "error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user created successfully"})

}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var founduser models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := db.DB.Where("email = ?", user.Email).First(&founduser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "invalid email or password", http.StatusUnauthorized)
			return
		}
		http.Error(w, "error fetcing user", http.StatusInternalServerError)
		return
	}

	if !utils.CheckPassword(founduser.Password, user.Password) {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	tokenstring, err := utils.GenerateJWT(founduser.Email)
	if err != nil {
		http.Error(w, "error during token genration", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenstring})

}
