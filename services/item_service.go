package services

import (
	"errors"
	"project-app-inventaris-cli-zahra/models"
	"project-app-inventaris-cli-zahra/repository"
	"strings"
	"time"
)

func ListItems() ([]models.Item, error) {
	return repository.GetAllItems()
}

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

func GetItemDetail(id int) (models.Item, error) {
	return repository.GetItemByID(id)
}

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

func RemoveItem(id int) error {
	return repository.DeleteItem(id)
}
