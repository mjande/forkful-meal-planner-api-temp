package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mjande/forkful-meal-planner-api/internal/models"
)

// Handles getting a list of recipes.
func GetRecipes(w http.ResponseWriter, r *http.Request) {
	// Call database function to query recipes
	recipes, err := models.ListRecipes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Encode the recipes in JSON and send as response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(recipes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}
