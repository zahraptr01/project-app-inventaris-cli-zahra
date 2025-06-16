package services

import (
	"project-app-inventaris-cli-zahra/database"
	"project-app-inventaris-cli-zahra/models"
	"project-app-inventaris-cli-zahra/repository"
	"testing"
	"time"
)

func TestAddAndGetItem(t *testing.T) {
	err := database.InitDB()
	if err != nil {
		t.Fatalf("Gagal koneksi ke database: %v", err)
	}

	// Add a test category first
	testCategory := models.Category{
		Name:        "Kategori Test Item",
		Description: "Kategori khusus untuk pengujian item",
	}

	err = repository.CreateCategory(testCategory)
	if err != nil {
		t.Fatalf("Gagal membuat kategori untuk item test: %v", err)
	}

	// Retrieve the category (use a name)
	categories, err := repository.GetAllCategories()
	if err != nil {
		t.Fatal("Gagal ambil kategori")
	}

	var catID int
	for _, cat := range categories {
		if cat.Name == "Kategori Test Item" {
			catID = cat.ID
			break
		}
	}

	if catID == 0 {
		t.Fatal("Kategori yang dibuat tidak ditemukan")
	}

	// Add test items
	testItem := models.Item{
		Name:         "Laptop Test",
		Price:        15000000,
		PurchaseDate: time.Now().AddDate(-1, 0, 0),
		CategoryID:   catID,
	}

	err = repository.CreateItem(testItem)
	if err != nil {
		t.Fatalf("Gagal create item: %v", err)
	}

	// Clear test data when done
	defer func() {
		_ = repository.DeleteItemByName("Laptop Test")
		_ = repository.DeleteCategoryByName("Kategori Test Item")
	}()

	// Check if items were successfully entered
	items, err := repository.GetAllItems()
	if err != nil {
		t.Fatalf("Gagal ambil item: %v", err)
	}

	found := false
	for _, item := range items {
		if item.Name == "Laptop Test" {
			found = true
			break
		}
	}

	if !found {
		t.Error("Item yang ditambahkan tidak ditemukan")
	}
}
