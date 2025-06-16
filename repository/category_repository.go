package repository

import (
	"database/sql"
	"errors"
	"project-app-inventaris-cli-zahra/database"
	"project-app-inventaris-cli-zahra/models"
)

func GetAllCategories() ([]models.Category, error) {
	query := `SELECT id, name, description FROM categories ORDER BY id`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func CreateCategory(category models.Category) error {
	query := `INSERT INTO categories (name, description) VALUES ($1, $2)`
	_, err := database.DB.Exec(query, category.Name, category.Description)
	return err
}

func GetCategoryByID(id int) (models.Category, error) {
	query := `SELECT id, name, description FROM categories WHERE id = $1`
	row := database.DB.QueryRow(query, id)

	var c models.Category
	err := row.Scan(&c.ID, &c.Name, &c.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return c, errors.New("kategori tidak ditemukan")
		}
		return c, err
	}
	return c, nil
}

func UpdateCategory(id int, category models.Category) error {
	query := `UPDATE categories SET name = $1, description = $2 WHERE id = $3`
	result, err := database.DB.Exec(query, category.Name, category.Description, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("kategori tidak ditemukan")
	}
	return nil
}

func DeleteCategory(id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	result, err := database.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("kategori tidak ditemukan")
	}
	return nil
}
