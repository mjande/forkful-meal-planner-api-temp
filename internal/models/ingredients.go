package models

import (
	_ "github.com/mattn/go-sqlite3"
)

type Ingredient struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	RecipeID int64   `json:"recipeId"`
	Quantity float32 `json:"quantity"`
	Unit     string  `json:"unit"`
}

// Queries the database for all unique ingredients used in any recipe.
func ListIngredients() ([]string, error) {
	query := `SELECT name FROM ingredients`

	// Execute query
	rows, err := db.Query(query)
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()

	// Map database response onto ingredients set
	ingredientsSet := map[string]bool{}
	for rows.Next() {
		var ingredient string

		err = rows.Scan(&ingredient)
		if err != nil {
			return []string{}, err
		}

		ingredientsSet[ingredient] = true
	}

	if err = rows.Err(); err != nil {
		return []string{}, err
	}

	// Extract all unique ingredients in map into slice
	var ingredients []string
	for ingredient := range ingredientsSet {
		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}
