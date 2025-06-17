package handlers

import (
	"encoding/json"
	"net/http"
	"broker-backend/db"
	"broker-backend/models"
	"broker-backend/utils"
	"database/sql"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	_, err := db.DB.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", req.Email, string(hash))
	if err != nil {
		http.Error(w, "User exists or DB error", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	var user models.User
	err := db.DB.QueryRow("SELECT id, password FROM users WHERE email=$1", req.Email).
		Scan(&user.ID, &user.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	access, _ := utils.GenerateToken(req.Email, 10*time.Minute)
	refresh, _ := utils.GenerateToken(req.Email, 24*time.Hour)

	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	var body map[string]string
	json.NewDecoder(r.Body).Decode(&body)

	token, err := utils.ParseToken(body["refresh_token"])
	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	access, _ := utils.GenerateToken(email, 10*time.Minute)
	json.NewEncoder(w).Encode(map[string]string{"access_token": access})
}
