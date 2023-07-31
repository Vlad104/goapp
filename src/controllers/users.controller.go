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

// UsersController представляет контроллер пользователей.
type UsersController struct {
	service *services.UsersService
}

// NewUsersController создает новый экземпляр UserController.
func NewUsersController(service *services.UsersService) *UsersController {
	return &UsersController{service}
}

// FindAll обрабатывает запрос на получение всех пользователей.
func (uc *UsersController) FindAll(w http.ResponseWriter, r *http.Request) {
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
func (uc *UsersController) FindById(w http.ResponseWriter, r *http.Request) {
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
func (uc *UsersController) FindByEmail(w http.ResponseWriter, r *http.Request) {
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
func (uc *UsersController) Create(w http.ResponseWriter, r *http.Request) {
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
func (uc *UsersController) Update(w http.ResponseWriter, r *http.Request) {
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

func (uc *UsersController) FindByMe(w http.ResponseWriter, r *http.Request) {
    authData := r.Context().Value("authData").(*entities.AuthData) // получение authData из контекста запроса

    if authData == nil {
        // Если authData равен nil, значит данные аутентификации не были добавлены в контекст,
        // что может произойти, если middleware AuthMiddleware не был применен или возникли проблемы с аутентификацией.
        // Обрабатываем ошибку, например, отправляем ответ с кодом ошибки HTTP 401 Unauthorized и сообщением об ошибке.
        common.HandleHttpError(w, common.ForbiddenError)
        return
    }

    id, err := common.UUIDFromString(authData.UserId)

    if err != nil {
        // Если произошла ошибка при преобразовании строки идентификатора пользователя в UUID,
        // тогда данные аутентификации содержат некорректный или невалидный идентификатор пользователя.
        // Обрабатываем ошибку, например, отправляем ответ с кодом ошибки HTTP 400 Bad Request и сообщением об ошибке.
        common.HandleHttpError(w, err)
        return
    }

    user, err := uc.service.FindById(id)

    if err != nil {
        // Если произошла ошибка при поиске пользователя в базе данных по идентификатору,
        // тогда обрабатываем ошибку, например, отправляем ответ с кодом ошибки HTTP 500 Internal Server Error и сообщением об ошибке.
        common.HandleHttpError(w, err)
        return
    }

    // Если все прошло успешно и пользователь был найден, то отправляем успешный ответ с данными пользователя.
    // Например, можем использовать код HTTP 200 OK и преобразовать найденного пользователя в JSON и отправить его клиенту.
    // Здесь предполагается, что у вас есть функция для преобразования данных пользователя в JSON, например, json.Marshal(users).
    response, err := json.Marshal(user)
    if err != nil {
        http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
        return
    }

    w.Write(response)
}


// Delete обрабатывает запрос на удаление пользователя по идентификатору.
func (uc *UsersController) Delete(w http.ResponseWriter, r *http.Request) {
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
