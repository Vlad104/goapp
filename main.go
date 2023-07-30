package main

import (
	"app/src/controllers"
	"app/src/database"
	"app/src/middlewares"
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

	usersService := services.NewUsersServices(
		repositories.NewUserRepositories(
			dataBase,
		),
	)

	// Создаем экземпляр контроллера пользователя
	usersController := controllers.NewUsersController(usersService)

	authController := controllers.NewAuthController(services.NewAuthServices(usersService))

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
		router.Get("/", usersController.FindAll)            // Получить всех пользователей
		router.Get("/{id}", usersController.FindById)       // Получить пользователя по идентификатору
		router.Get("/{email}", usersController.FindByEmail) // Получить пользователя по адресу электронной почты

		// Обработка POST-запроса
		router.Post("/", usersController.Create) // Создать нового пользователя

		// Обработка PUT-запроса
		router.Put("/", usersController.Update) // Обновить информацию о пользователе

		// Обработка DELETE-запроса
		router.Delete("/{id}", usersController.Delete) // Удалить пользователя по идентификатору

		router.With(middlewares.AuthMiddleware).Put("/", usersController.Update) // Обновить информацию о пользователе
		router.With(middlewares.AuthMiddleware).Delete("/{id}", usersController.Delete) // Удалить пользователя по идентификатору
		router.With(middlewares.AuthMiddleware).Get("/me", usersController.FindByMe)
	})

	router.Route("/auth", func(router chi.Router) {
		router.Post("/login", authController.Login)
		
		
	})


	// Запуск HTTP-сервера и обработка запросов с помощью роутера
	log.Fatal(http.ListenAndServe(":80", router))
}
