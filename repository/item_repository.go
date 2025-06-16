package repository

import (
	"database/sql"
	"errors"
	"project-app-inventaris-cli-zahra/database"
	"project-app-inventaris-cli-zahra/models"
	"time"
)

func CreateItem(item models.Item) error {
	query := `INSERT INTO items (name, category_id, price, purchase_date) VALUES ($1, $2, $3, $4)`
	_, err := database.DB.Exec(query, item.Name, item.CategoryID, item.Price, item.PurchaseDate)
	return err
}

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
