package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/mjande/forkful-meal-planner-api/internal/handlers"
	"github.com/mjande/forkful-meal-planner-api/internal/models"
)

func main() {
	err := models.InitDB("/Users/mattanderson/Code/cs-post-bacc/361-swe1/forkful-meal-planner-api/sqlite/forkful-meal-planner-db?_fk=true")
	if err != nil {
		log.Panicln(err)
		models.CloseDB()
	}

	defer models.CloseDB()
	log.Println("Database successfully connected")

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Use(middleware.Logger)

	router.Route("/ingredients", func(r chi.Router) {
		r.Get("/", handlers.GetIngredients)
		r.Post("/", handlers.PostIngredient)
		r.Patch("/{id}", handlers.PatchIngredient)
		r.Delete("/{id}", handlers.DeleteIngredient)
	})

	router.Route("/recipes", func(r chi.Router) {
		r.Get("/", handlers.GetRecipes)
	})

	port := 3001

	log.Printf("Listening on port %d...", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	log.Fatal(err)
}
