package services

import (
	"errors"
	"project-app-inventaris-cli-zahra/models"
	"project-app-inventaris-cli-zahra/repository"
	"strings"
	"time"
)

// using the GetAllItems function
func ListItems() ([]models.Item, error) {
	return repository.GetAllItems()
}

// menggunakan fungsi CreateItem
func AddItem(name string, categoryID int, price float64, purchaseDate time.Time) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("nama barang tidak boleh kosong")
	}
	item := models.Item{
		Name:         name,
		CategoryID:   categoryID,
		Price:        price,
		PurchaseDate: purchaseDate,
	}
	return repository.CreateItem(item)
}

// menggunakan fungsi GetItemByID
func GetItemDetail(id int) (models.Item, error) {
	return repository.GetItemByID(id)
}

// menggunakan fungsi UpdateItem
func EditItem(id int, name string, categoryID int, price float64, purchaseDate time.Time) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("nama barang tidak boleh kosong")
	}
	item := models.Item{
		Name:         name,
		CategoryID:   categoryID,
		Price:        price,
		PurchaseDate: purchaseDate,
	}
	return repository.UpdateItem(id, item)
}

// menggunakan fungsi DeleteItem
func RemoveItem(id int) error {
	return repository.DeleteItem(id)
}

// menggunakan fungsi GetAllItems for search item
func SearchItemsByName(keyword string) ([]models.Item, error) {
	items, err := repository.GetAllItems()
	if err != nil {
		return nil, err
	}

	var result []models.Item
	for _, item := range items {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(keyword)) {
			result = append(result, item)
		}
	}
	return result, nil
}

// func SearchItemsByName(keyword string) ([]models.Item, error) {
// 	return repositories.SearchItemsByName(keyword)
// }
