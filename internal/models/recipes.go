package models

import (
	_ "github.com/mattn/go-sqlite3"
)

type Recipe struct {
	ID           int                `json:"id"`
	Name         string             `json:"name"`
	CookingTime  string             `json:"cookingTime"`
	Description  string             `json:"description"`
	Instructions string             `json:"instructions"`
	Ingredients  []RecipeIngredient `json:"ingredients"`
}

type RecipeIngredient struct {
	ID           int     `json:"id"`
	IngredientID int     `json:"ingredientId"`
	Name         string  `json:"name"`
	Unit         string  `json:"unit"`
	Quantity     float32 `json:"quantity"`
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
func FindRecipe(id int) (Recipe, error) {
	query := `SELECT id, name, cooking_time, description, instructions FROM recipes WHERE id = ?`

	// Query the database
	result := db.QueryRow(query, id)

	// Scan database result into recipe object
	var recipe Recipe
	err := result.Scan(&recipe.ID, &recipe.Name, &recipe.CookingTime, &recipe.Description, &recipe.Instructions)
	if err != nil {
		return Recipe{}, err
	}

	ingredientsQuery := `
		SELECT ri.id, i.id, i.name, ri.unit, ri.quantity 
    		FROM recipe_ingredients as ri 
		LEFT JOIN ingredients AS i ON i.id = ri.ingredient_id
    		WHERE ri.recipe_id = ?
	`

	// Get all ingredients used in this recipe
	rows, err := db.Query(ingredientsQuery, recipe.ID)
	if err != nil {
		return Recipe{}, err
	}
	defer rows.Close()

	// Map database response onto ingredientss slice
	var ingredients []RecipeIngredient
	for rows.Next() {
		var ingredient RecipeIngredient

		err = rows.Scan(&ingredient.ID, &ingredient.IngredientID, &ingredient.Name, &ingredient.Unit, &ingredient.Quantity)
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
