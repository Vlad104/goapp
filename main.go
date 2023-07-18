package main

import (
	"app/src/controllers"
	"app/src/database"
	"app/src/repositories"
	"app/src/services"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// user -> controller -> services -> repositories (сущность)

func main() {
	// Создаем новое подключение к базе данных
	dataBase, err := database.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	// Создаем экземпляр контроллера пользователя
	userController := controllers.New(
		services.New(
			repositories.New(
				dataBase,
			),
		),
	)

	// Создаем новый роутер Chi
	router := chi.NewRouter()

	// Используем промежуточное ПО для обработки запросов
	router.Use(middleware.SetHeader("Content-Type", "application/json"))
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)

	// Настраиваем маршруты для роутера
	router.Route("/users", func(router chi.Router) {
		// Обработка GET-запросов
		router.Get("/", userController.FindAll)          // Получить всех пользователей
		router.Get("/{id}", userController.FindById)    // Получить пользователя по идентификатору
		router.Get("/{email}", userController.FindByEmail) // Получить пользователя по адресу электронной почты

		// Обработка POST-запроса
		router.Post("/", userController.Create)          // Создать нового пользователя

		// Обработка PUT-запроса
		router.Put("/", userController.Update)           // Обновить информацию о пользователе

		// Обработка DELETE-запроса
		router.Delete("/{id}", userController.Delete)    // Удалить пользователя по идентификатору
	})

	// Запуск HTTP-сервера и обработка запросов с помощью роутера
	log.Fatal(http.ListenAndServe(":80", router))
}
