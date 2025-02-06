package models

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Ingredient struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Queries the database for all ingredients.
func ListIngredients() ([]Ingredient, error) {
	query := `SELECT id, name FROM ingredients`

	// Sends query
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error encountered")
		return []Ingredient{}, err
	}
	defer rows.Close()

	// Maps database response into ingredients slice
	var ingredients []Ingredient
	for rows.Next() {
		var ingredient Ingredient

		err = rows.Scan(&ingredient.ID, &ingredient.Name)
		if err != nil {
			return nil, err
		}

		ingredients = append(ingredients, ingredient)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ingredients, nil
}

// Queries the database for an ingredient that has the given id.
func FindIngredient(id int) (Ingredient, error) {
	query := `SELECT id, name FROM ingredients WHERE id = ?`

	// Query the database
	result := db.QueryRow(query, id)

	// Scan database result into ingredient object
	var ingredient Ingredient
	err := result.Scan(&ingredient.ID, &ingredient.Name)
	if err != nil {
		return ingredient, err
	}

	return ingredient, nil

}

// Queries the database to create an ingredient.
func CreateIngredient(name string) (int, error) {
	query := `INSERT INTO ingredients (name) VALUES (?)`

	// Send query
	result, err := db.Exec(query, name)
	if err != nil {
		return -1, err
	}

	// Get id of created ingredient
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func UpdateIngredient(id int, ingredient Ingredient) error {
	query := `UPDATE ingredients SET name = ? WHERE id = ?`

	// Send query
	_, err := db.Exec(query, ingredient.Name, id)
	if err != nil {
		return err
	}

	return nil
}
