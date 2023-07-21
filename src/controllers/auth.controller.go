package controllers

import (
	"app/src/entities"
	"app/src/services"
	"encoding/json"
	"net/http"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService}
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var loginDto *entities.LoginDto

	err := json.NewDecoder(r.Body).Decode(&loginDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authDto, err := ac.service.Login(loginDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(authDto)
}
