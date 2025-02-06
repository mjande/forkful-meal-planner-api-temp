package models

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Ingredient struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ListIngredients() ([]Ingredient, error) {
	query := `SELECT id, name FROM ingredients`

	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error encountered")
		return []Ingredient{}, err
	}
	defer rows.Close()

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

func FindIngredient(id int) (Ingredient, error) {
	query := `SELECT id, name FROM ingredients WHERE id = ?`

	result := db.QueryRow(query, id)

	var ingredient Ingredient

	err := result.Scan(&ingredient.ID, &ingredient.Name)
	if err != nil {
		return ingredient, err
	}

	return ingredient, nil

}

func CreateIngredient(name string) (int, error) {
	query := `INSERT INTO ingredients (name) VALUES (?)`

	result, err := db.Exec(query, name)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}
