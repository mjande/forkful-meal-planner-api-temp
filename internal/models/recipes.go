package models

import (
	"database/sql"
	"slices"

	_ "github.com/mattn/go-sqlite3"
)

type Recipe struct {
	ID           int64        `json:"id"`
	Name         string       `json:"name"`
	CookingTime  string       `json:"cookingTime"`
	Description  string       `json:"description"`
	Instructions string       `json:"instructions"`
	Ingredients  []Ingredient `json:"ingredients"`
	Tags         []string     `json:"tags"`
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

		// Get all tags for this recipe
		tags, err := findTagsByRecipe(recipe.ID)
		if err != nil {
			return []Recipe{}, err
		}

		// Add the tag name to this recipe
		for _, tag := range tags {
			recipe.Tags = append(recipe.Tags, tag.Name)
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

	tags, err := findTagsByRecipe(recipe.ID)
	if err != nil {
		return Recipe{}, err
	}

	var tagStrs []string
	for _, tag := range tags {
		tagStrs = append(tagStrs, tag.Name)
	}

	// Add ingredients and tags to recipe object
	recipe.Ingredients = ingredients
	recipe.Tags = tagStrs

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

	// Create ingredients
	for i := 0; i < len(recipe.Ingredients); i++ {
		ingredient := recipe.Ingredients[i]
		ingredient.RecipeID = id

		_, err = createIngredient(ingredient)
		if err != nil {
			return -1, err
		}
	}

	// Create tags
	for _, tag := range recipe.Tags {
		_, err := createTag(id, tag)
		if err != nil {
			return -1, err
		}
	}

	return id, nil
}

func UpdateRecipe(id int64, recipe Recipe) (int64, error) {
	query := `UPDATE recipes SET name = ?, cooking_time = ?, description = ?, instructions = ? WHERE id = ?`

	// Send query
	_, err := db.Exec(query, recipe.Name, recipe.CookingTime, recipe.Description, recipe.Instructions, id)
	if err != nil {
		return -1, err
	}

	var keptIngredientIds []int64
	for i := 0; i < len(recipe.Ingredients); i++ {
		ingredient := recipe.Ingredients[i]
		ingredient.RecipeID = id

		// Check if ingredient was previously in recipe
		prevIngredient, err := FindIngredient(ingredient.Name, id)
		if err != nil && err == sql.ErrNoRows {
			// A previous version of this recipe's ingredient does not exist, so
			// create it
			_, err = createIngredient(ingredient)
			if err != nil {
				return -1, err
			}
		} else if err != nil {
			return -1, err
		} else {
			// Update previous version of ingredient
			query = `UPDATE ingredients SET name = ?, quantity = ?, unit = ? WHERE id = ?`

			_, err = db.Exec(query, ingredient.Name, ingredient.Quantity, ingredient.Unit, prevIngredient.ID)
			if err != nil {
				return -1, err
			}

			keptIngredientIds = append(keptIngredientIds, prevIngredient.ID)
		}
	}

	// Remove ingredients that are no longer used
	deleteQuery := `DELETE FROM ingredients WHERE id = ?`

	prevAndCurrIngredients, err := ListIngredientsByRecipe(id)
	if err != nil {
		return -1, err
	}

	for _, ingredient := range prevAndCurrIngredients {
		if !slices.Contains(keptIngredientIds, ingredient.ID) {
			db.Exec(deleteQuery, ingredient.ID)
		}
	}

	err = updateRecipeTags(id, recipe)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func DeleteRecipe(id int64) error {
	query := `DELETE FROM recipes WHERE id = ?`

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

// Helper Functions

// Takes a recipe ID and an updated recipe object. Updates the recipe in the
// datase to reflect the new list of tags.
func updateRecipeTags(recipeId int64, recipe Recipe) error {
	tagsToDeleteSlice, err := findTagsByRecipe(recipeId)
	if err != nil {
		return err
	}

	tagsToDelete := map[int64]bool{}
	for _, tag := range tagsToDeleteSlice {
		tagsToDelete[tag.ID] = true
	}

	for _, tag := range recipe.Tags {
		// Check if recipe previously included tag
		prevTag, err := FindTag(recipeId, tag)
		if err != nil && err == sql.ErrNoRows {
			// A previous version of this tag does not exist, so create it
			_, err = createTag(recipeId, tag)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		} else {
			tagsToDelete[prevTag.ID] = false
		}
	}

	deleteQuery := `DELETE FROM recipe_tags WHERE id = ?`

	for id, delete := range tagsToDelete {
		if delete {
			db.Exec(deleteQuery, id)
		}
	}

	return nil
}
