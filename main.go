package main

import (
	"app/src/controllers"
	"app/src/database"
	"app/src/repositories"
	"app/src/services"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	database, err := database.New()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	userController := controllers.New(
		services.New(
			repositories.New(
				database,
			),
		),
	)

	router := chi.NewRouter()

	router.Use(middleware.SetHeader("Content-Type", "application/json"))
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)

	router.Route("/users", func(router chi.Router) {
		router.Get("/", userController.FindAll)
		router.Get("/{id}", userController.FindById)
		router.Post("/", userController.Create)
		router.Put("/", userController.Update)
		router.Delete("/{id}", userController.Delete)
	})

	http.ListenAndServe(":80", router)
}
