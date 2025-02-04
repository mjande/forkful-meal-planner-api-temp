package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mjande/forkful-meal-planner-api/internal/handlers"
	"github.com/mjande/forkful-meal-planner-api/internal/models"
)

func main() {
	err := models.InitDB("/Users/mattanderson/Code/cs-post-bacc/361-swe1/forkful-meal-planner-api/sqlite/forkful-meal-planner-db")
	if err != nil {
		log.Panicln(err)
		models.CloseDB()
	}

	defer models.CloseDB()
	log.Println("Database successfully connected")

	router := chi.NewRouter()
	router.Route("/ingredients", func(r chi.Router) {
		r.Get("/", handlers.GetIngredients)
	})

	port := 3001

	log.Printf("Listening on port %d...", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	log.Fatal(err)
}
