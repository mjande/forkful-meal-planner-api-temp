package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mjande/forkful-meal-planner-api/internal/models"
)

func GetIngredients(w http.ResponseWriter, r *http.Request) {
	ingredients, err := models.ListIngredients()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	body, err := json.Marshal(ingredients)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Write(body)
}

func PostIngredient(w http.ResponseWriter, r *http.Request) {
	var ingredient models.Ingredient
	err := json.NewDecoder(r.Body).Decode(&ingredient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	id, err := models.CreateIngredient(ingredient.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	ingredient, err = models.FindIngredient(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	body, err := json.Marshal(ingredient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}
