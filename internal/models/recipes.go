package models

import (
	_ "github.com/mattn/go-sqlite3"
)

type Recipe struct {
	ID           int64        `json:"id"`
	Name         string       `json:"name"`
	CookingTime  string       `json:"cookingTime"`
	Description  string       `json:"description"`
	Instructions string       `json:"instructions"`
	Ingredients  []Ingredient `json:"ingredients"`
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

// Queries the database for an ingredient that has the given id.
func FindRecipe(id int64) (Recipe, error) {
	query := `SELECT id, name, cooking_time, description, instructions FROM recipes WHERE id = ?`

	// Query the database
	result := db.QueryRow(query, id)

	// Scan database result into recipe object
	var recipe Recipe
	err := result.Scan(&recipe.ID, &recipe.Name, &recipe.CookingTime, &recipe.Description, &recipe.Instructions)
	if err != nil {
		return Recipe{}, err
	}

	ingredientsQuery := `SELECT id, name, quantity, unit FROM ingredients WHERE recipe_id = ?`

	// Get all ingredients used in this recipe
	rows, err := db.Query(ingredientsQuery, recipe.ID)
	if err != nil {
		return Recipe{}, err
	}
	defer rows.Close()

	// Map database response onto ingredients slice
	var ingredients []Ingredient
	for rows.Next() {
		var ingredient Ingredient

		err = rows.Scan(&ingredient.ID, &ingredient.Name, &ingredient.Quantity, &ingredient.Unit)
		if err != nil {
			return Recipe{}, err
		}

		ingredients = append(ingredients, ingredient)
	}

	if err = rows.Err(); err != nil {
		return Recipe{}, err
	}

	// Add ingredients slice to recipe object
	recipe.Ingredients = ingredients

	return recipe, nil
}

func CreateRecipe(recipe Recipe) (int64, error) {
	query := `INSERT INTO recipes (name, cooking_time, description, instructions) VALUES (?, ?, ?, ?)`

	// Send query
	result, err := db.Exec(query, recipe.Name, recipe.CookingTime, recipe.Description, recipe.Instructions)
	if err != nil {
		return -1, err
	}

	// Get id of created recipe
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	for i := 0; i < len(recipe.Ingredients); i++ {
		ingredient := recipe.Ingredients[i]

		// Create association between recipe and ingredient
		recipeIngredientQuery := `INSERT INTO ingredients (name, recipe_id, quantity, unit) VALUES (?, ?, ?, ?)`

		_, err = db.Exec(recipeIngredientQuery, ingredient.Name, id, ingredient.Quantity, ingredient.Unit)
		if err != nil {
			return -1, err
		}
	}

	return id, nil
}
