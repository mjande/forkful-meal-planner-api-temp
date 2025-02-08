package models

import (
	_ "github.com/mattn/go-sqlite3"
)

type Recipe struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	CookingTime  string       `json:"cookingTime"`
	Description  string       `json:"description"`
	Instructions string       `json:"instructions"`
	Ingredients  []Ingredient `json:"ingredients"`
}

type RecipeIngredient struct {
	Ingredient Ingredient `json:"ingredient"`
	Unit       string     `json:"unit"`
	Quantity   float32    `json:"quantity"`
}

// Queries the database for all recipes (while only loading basic
// data for index page)
func ListRecipes() ([]Recipe, error) {
	query := `SELECT id, name, cooking_time, description FROM recipes`

	// Execute query
	rows, err := db.Query(query)
	if err != nil {
		return []Recipe{}, err
	}
	defer rows.Close()

	// Map database response onto recipes slice
	var recipes []Recipe
	for rows.Next() {
		var recipe Recipe

		err = rows.Scan(&recipe.ID, &recipe.Name, &recipe.CookingTime, &recipe.Description)
		if err != nil {
			return []Recipe{}, err
		}

		recipes = append(recipes, recipe)
	}

	if err = rows.Err(); err != nil {
		return []Recipe{}, err
	}

	return recipes, nil
}
