package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mjande/forkful-meal-planner-api/internal/models"
)

// Handles getting a unique list of ingredients used in other recipes.
func GetIngredients(w http.ResponseWriter, r *http.Request) {
	// Call database function to query ingredients
	ingredients, err := models.ListIngredients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Encode the ingredients in JSON and send as response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(ingredients)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}
