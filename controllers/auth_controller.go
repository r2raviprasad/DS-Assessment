package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"find-pairs/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New() // initializing validator
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type LoginResponse struct {
	Token string `json:"token"`
}

func generateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 2400).Unix(), // set expiry of jwt token to 2400 hours
	})

	return token.SignedString([]byte(config.JwtSecret))
}
func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req) // bind request body to LoginRequest struct
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// validate request parameters
	err = validate.Struct(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if CredentialAuthenticate(req.Username, req.Password) {
		token, err := generateJWT(req.Username)
		if err != nil {
			http.Error(w, "Could not generate token", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(LoginResponse{Token: token})
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}
func CredentialAuthenticate(username, password string) bool {
	return username == config.Username && password == config.Password
}
