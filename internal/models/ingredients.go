package models

import (
	"database/sql"

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

func ListIngredientsByRecipe(recipeId int64) ([]Ingredient, error) {
	query := `SELECT id, name, recipe_id, quantity, unit FROM ingredients WHERE recipe_id = ?`

	rows, err := db.Query(query, recipeId)
	if err != nil && err == sql.ErrNoRows {
		return []Ingredient{}, nil
	} else if err != nil {
		return []Ingredient{}, err
	}

	// Map database response onto recipes slice
	var ingredients []Ingredient
	for rows.Next() {
		var ingredient Ingredient

		err = rows.Scan(&ingredient.ID, &ingredient.Name, &ingredient.RecipeID, &ingredient.Quantity, &ingredient.Unit)
		if err != nil {
			return []Ingredient{}, err
		}

		ingredients = append(ingredients, ingredient)
	}

	if err = rows.Err(); err != nil {
		return []Ingredient{}, err
	}

	return ingredients, nil
}

func FindIngredient(name string, recipeId int64) (Ingredient, error) {
	query := `SELECT id, name, recipe_id, quantity, unit FROM ingredients WHERE name = ? AND recipe_id = ?`

	// Query the database
	result := db.QueryRow(query, name, recipeId)

	// Scan database result into recipe object
	var ingredient Ingredient
	err := result.Scan(&ingredient.ID, &ingredient.Name, &ingredient.RecipeID, &ingredient.Quantity, &ingredient.Unit)
	if err != nil {
		return Ingredient{}, err
	}

	return ingredient, nil
}

func createIngredient(ingredient Ingredient) (int64, error) {
	query := `INSERT INTO ingredients (name, recipe_id, quantity, unit) VALUES (?, ?, ?, ?)`

	result, err := db.Exec(query, ingredient.Name, ingredient.RecipeID, ingredient.Quantity, ingredient.Unit)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}
