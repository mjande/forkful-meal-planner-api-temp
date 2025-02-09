package models

import "database/sql"

type Meal struct {
	ID       int64  `json:"id"`
	RecipeId int64  `json:"recipeId"`
	Date     string `json:"date"`
}

func ListMealsByDate(start string, end string) ([]Meal, error) {
	query := `SELECT id, recipe_id, date FROM meals where date BETWEEN ? AND ?`

	rows, err := db.Query(query, start, end)
	if err != nil && err == sql.ErrNoRows {
		return []Meal{}, nil
	} else if err != nil {
		return []Meal{}, err
	}

	var meals []Meal
	for rows.Next() {
		var meal Meal

		err = rows.Scan(&meal.ID, &meal.RecipeId, &meal.Date)
		if err != nil {
			return []Meal{}, err
		}

		meals = append(meals, meal)
	}

	return meals, nil
}

func FindMeal(id int64) (Meal, error) {
	query := `SELECT id, recipe_id, date FROM meals WHERE id = ?`

	result := db.QueryRow(query, id)

	var meal Meal
	err := result.Scan(&meal.ID, &meal.RecipeId, &meal.Date)
	if err != nil {
		return Meal{}, err
	}

	return meal, nil
}

func CreateMeal(meal Meal) (int64, error) {
	query := `INSERT INTO meals (recipe_id, date) VALUES (?, ?)`

	result, err := db.Exec(query, meal.RecipeId, meal.Date)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func DeleteMeal(id int64) error {
	query := `DELETE FROM meals WHERE id = ?`

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
