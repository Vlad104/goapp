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

type UserController struct {
	service *services.UsersService
}

func New(service *services.UsersService) *UserController {
	return &UserController{service}
}

func (controller *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	users, err := controller.service.FindAll()

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	response, err := json.Marshal(users)

	w.Write(response)
}

func (controller *UserController) FindById(w http.ResponseWriter, r *http.Request) {
	id, err := common.UUIDFromString(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("could not parse id %v as UUID", chi.URLParam(r, "id"))
		common.HandleHttpError(w, common.LogicError)
		return
	}

	users, err := controller.service.FindById(id)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	response, err := json.Marshal(users)

	w.Write(response)
}

func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var createUserDto entities.CreateUserDto
	err := json.NewDecoder(r.Body).Decode(&createUserDto)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	user, err := controller.service.Create(&createUserDto)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	response, err := json.Marshal(user)

	w.Write(response)
}

func (controller *UserController) Update(w http.ResponseWriter, r *http.Request) {
	var updateUserDto entities.UpdateUserDto
	err := json.NewDecoder(r.Body).Decode(&updateUserDto)

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	user, err := controller.service.Update(&updateUserDto)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	response, err := json.Marshal(user)

	w.Write(response)
}

func (controller *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := common.UUIDFromString(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("could not parse id %v as UUID", chi.URLParam(r, "id"))
		common.HandleHttpError(w, common.LogicError)
		return
	}

	err = controller.service.Delete(id)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	w.WriteHeader(200)
}
