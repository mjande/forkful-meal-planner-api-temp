package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

// Handles getting a single recipe.
func GetRecipe(w http.ResponseWriter, r *http.Request) {
	// Extract id from request
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	// Call database function to query recipes
	recipe, err := models.FindRecipe(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Println(err)
		return
	}

	// Encode the recipes in JSON and send as response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}

// Handles creating a recipe with ingredients
func PostRecipe(w http.ResponseWriter, r *http.Request) {
	// Decode JSON data from request
	var recipe models.Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	// Use database function to create recipe
	id, err := models.CreateRecipe(recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	// Get recipe from database
	recipe, err = models.FindRecipe(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err, "Could not find recipe ", id)
		return
	}

	// Encode recipe as JSON and send response
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func PatchRecipe(w http.ResponseWriter, r *http.Request) {
	// Extract id from request
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	// Check that recipe exists
	_, err = models.FindRecipe(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Println(err)
		return
	}

	// Decode JSON data from request
	var recipe models.Recipe
	err = json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	// Use database function to create recipe
	id, err = models.UpdateRecipe(id, recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	// Get recipe from database
	recipe, err = models.FindRecipe(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Encode recipe as JSON and send response
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	// Extract id from request
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	err = models.DeleteRecipe(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
