package models

import "database/sql"

type Tag struct {
	ID       int64
	RecipeId int64
	Name     string
}

// Get all tags for a given recipe
func findTagsByRecipe(recipeId int64) ([]Tag, error) {
	query := `SELECT id, recipe_id, name FROM recipe_tags WHERE recipe_id = ?`

	rows, err := db.Query(query, recipeId)
	if err != nil && err == sql.ErrNoRows {
		return []Tag{}, nil
	} else if err != nil {
		return []Tag{}, err
	}
	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		var tag Tag

		err = rows.Scan(&tag.ID, &tag.RecipeId, &tag.Name)
		if err != nil {
			return []Tag{}, err
		}

		tags = append(tags, tag)
	}

	if err = rows.Err(); err != nil {
		return []Tag{}, err
	}

	return tags, nil
}

func FindTag(recipeId int64, name string) (Tag, error) {
	query := `SELECT id, recipe_id, name FROM recipe_tags WHERE recipe_id = ? AND name = ?`

	result := db.QueryRow(query, recipeId, name)

	var tag Tag
	err := result.Scan(&tag.ID, &tag.RecipeId, &tag.Name)
	if err != nil {
		return Tag{}, err
	}

	return tag, nil
}

func createTag(recipeId int64, tag string) (int64, error) {
	query := `INSERT INTO recipe_tags (recipe_id, name) VALUES (?, ?)`

	result, err := db.Exec(query, recipeId, tag)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}
