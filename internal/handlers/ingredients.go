package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mjande/forkful-meal-planner-api/internal/models"
)

// Handles getting a list of ingredients.
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
		return
	}
}

// Handles creating a new ingredient.
func PostIngredient(w http.ResponseWriter, r *http.Request) {
	// Decode JSON data from request
	var ingredient models.Ingredient
	err := json.NewDecoder(r.Body).Decode(&ingredient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	// Use database function to create ingredient
	id, err := models.CreateIngredient(ingredient.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	// Get ingredient from database
	ingredient, err = models.FindIngredient(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Encode ingredient as JSON and send response
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(ingredient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

// Handles updating an ingredient.
func PatchIngredient(w http.ResponseWriter, r *http.Request) {
	// Get ID from request
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	// Check that ingredient exists
	_, err = models.FindIngredient(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Println(err)
		return
	}

	// Decode JSON data from request
	var ingredient models.Ingredient
	err = json.NewDecoder(r.Body).Decode(&ingredient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	// Use database function to create ingredient
	err = models.UpdateIngredient(ingredient.ID, ingredient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	// Get ingredient from database
	ingredient, err = models.FindIngredient(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Encode ingredient as JSON and send response
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(ingredient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
