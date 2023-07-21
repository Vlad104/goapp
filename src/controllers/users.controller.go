package controllers

import (
	"app/src/common"
	"app/src/entities"
	"app/src/services"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// UserController представляет контроллер пользователей.
type UserController struct {
	service *services.UsersService
}

// NewUserController создает новый экземпляр UserController.
func NewUserController(service *services.UsersService) *UserController {
	return &UserController{service}
}

// FindAll обрабатывает запрос на получение всех пользователей.
func (uc *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	users, err := uc.service.FindAll()

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	response, err := json.Marshal(users)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	w.Write(response)
}

// FindById обрабатывает запрос на получение пользователя по идентификатору.
func (uc *UserController) FindById(w http.ResponseWriter, r *http.Request) {
	id, err := common.UUIDFromString(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("could not parse id %v as UUID", chi.URLParam(r, "id"))
		common.HandleHttpError(w, common.LogicError)
		return
	}

	users, err := uc.service.FindById(id)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	response, err := json.Marshal(users)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	w.Write(response)
}

// FindByEmail обрабатывает запрос на получение пользователя по адресу электронной почты.
func (uc *UserController) FindByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	if len(email) == 0 {
		log.Printf("could not found email %v", chi.URLParam(r, "email"))
		common.HandleHttpError(w, common.LogicError)
		return
	}

	users, err := uc.service.FindByEmail(email)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	response, err := json.Marshal(users)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	w.Write(response)
}

// Create обрабатывает запрос на создание нового пользователя.
func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var createUserDto entities.CreateUserDto
	err := json.NewDecoder(r.Body).Decode(&createUserDto)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	user, err := uc.service.Create(&createUserDto)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	response, err := json.Marshal(user)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	w.Write(response)
}

// Update обрабатывает запрос на обновление информации о пользователе.
func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
	var updateUserDto entities.UpdateUserDto
	err := json.NewDecoder(r.Body).Decode(&updateUserDto)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user, err := uc.service.Update(&updateUserDto)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	response, err := json.Marshal(user)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	w.Write(response)
}

// Delete обрабатывает запрос на удаление пользователя по идентификатору.
func (uc *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := common.UUIDFromString(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("could not parse id %v as UUID", chi.URLParam(r, "id"))
		common.HandleHttpError(w, common.LogicError)
		return
	}

	err = uc.service.Delete(id)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
