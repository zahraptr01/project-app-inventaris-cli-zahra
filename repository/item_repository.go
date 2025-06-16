package repository

import (
	"database/sql"
	"errors"
	"project-app-inventaris-cli-zahra/database"
	"project-app-inventaris-cli-zahra/models"
	"time"
)

// view all inventory data
func GetAllItems() ([]models.Item, error) {
	query := `SELECT id, name, category_id, price, purchase_date FROM items ORDER BY id`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var i models.Item
		var purchaseDate time.Time
		if err := rows.Scan(&i.ID, &i.Name, &i.CategoryID, &i.Price, &purchaseDate); err != nil {
			return nil, err
		}
		i.PurchaseDate = purchaseDate
		items = append(items, i)
	}
	return items, nil
}

// create new inventory item
func CreateItem(item models.Item) error {
	query := `INSERT INTO items (name, category_id, price, purchase_date) VALUES ($1, $2, $3, $4)`
	_, err := database.DB.Exec(query, item.Name, item.CategoryID, item.Price, item.PurchaseDate)
	return err
}

// view inventory item data by id
func GetItemByID(id int) (models.Item, error) {
	query := `SELECT id, name, category_id, price, purchase_date FROM items WHERE id = $1`
	row := database.DB.QueryRow(query, id)

	var item models.Item
	err := row.Scan(&item.ID, &item.Name, &item.CategoryID, &item.Price, &item.PurchaseDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, errors.New("barang tidak ditemukan")
		}
		return item, err
	}
	return item, nil
}

// update inventory data
func UpdateItem(id int, item models.Item) error {
	query := `UPDATE items SET name = $1, category_id = $2, price = $3, purchase_date = $4 WHERE id = $5`
	result, err := database.DB.Exec(query, item.Name, item.CategoryID, item.Price, item.PurchaseDate, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("barang tidak ditemukan")
	}
	return nil
}

// delete inventory data
func DeleteItem(id int) error {
	query := `DELETE FROM items WHERE id = $1`
	result, err := database.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("barang tidak ditemukan")
	}
	return nil
}

// Fungsi pencarian barang berdasarkan nama (dengan keyword LIKE)
func SearchItemsByName(keyword string) ([]models.Item, error) {
	query := `
		SELECT i.id, i.name, i.price, i.purchase_date, i.category_id,
		       c.name, c.description
		FROM items i
		JOIN categories c ON i.category_id = c.id
		WHERE LOWER(i.name) LIKE LOWER($1)
	`

	rows, err := database.DB.Query(query, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		var category models.Category

		err := rows.Scan(
			&item.ID, &item.Name, &item.Price, &item.PurchaseDate, &item.CategoryID,
			&category.Name, &category.Description,
		)
		if err != nil {
			return nil, err
		}
		item.Category = category
		items = append(items, item)
	}

	return items, nil
}

// delete item by name (for testing)
func DeleteItemByName(name string) error {
	query := `DELETE FROM items WHERE name = $1`
	_, err := database.DB.Exec(query, name)
	return err
}
