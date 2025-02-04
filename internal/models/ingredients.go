package models

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Ingredient struct {
	ID   int
	Name string
}

func ListIngredients() ([]Ingredient, error) {
	log.Println(db)

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
