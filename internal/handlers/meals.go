package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mjande/forkful-meal-planner-api/internal/models"
)

// Handles getting a list of meals.
func GetMealsByDate(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	// Call database function to query meals
	meals, err := models.ListMealsByDate(params.Get("start"), params.Get("end"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	// Encode meals in JSON and send as response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(meals)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}

func PostMeal(w http.ResponseWriter, r *http.Request) {
	var meal models.Meal
	err := json.NewDecoder(r.Body).Decode(&meal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	id, err := models.CreateMeal(meal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	meal, err = models.FindMeal(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(meal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println()
	}
}
