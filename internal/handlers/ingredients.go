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
