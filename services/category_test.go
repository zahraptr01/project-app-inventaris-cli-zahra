package services

import (
	"project-app-inventaris-cli-zahra/database"
	"project-app-inventaris-cli-zahra/models"
	"project-app-inventaris-cli-zahra/repository"
	"testing"
)

func TestAddAndGetCategory(t *testing.T) {
	// Initialise DB connection
	err := database.InitDB()
	if err != nil {
		t.Fatalf("Gagal koneksi ke database: %v", err)
	}

	// Test category data
	testCategory := models.Category{
		Name:        "Test Kategori",
		Description: "Kategori untuk pengujian",
	}

	// Add category
	err = repository.CreateCategory(testCategory)
	if err != nil {
		t.Fatalf("Gagal tambah kategori: %v", err)
	}

	// Remove category after test
	defer func() {
		err := repository.DeleteCategoryByName("Test Kategori")
		if err != nil {
			t.Logf("Peringatan: Gagal hapus kategori setelah test: %v", err)
		}
	}()

	// Retrieve all categories
	categories, err := repository.GetAllCategories()
	if err != nil {
		t.Fatalf("Gagal ambil kategori: %v", err)
	}

	// Verify category found
	found := false
	for _, cat := range categories {
		if cat.Name == "Test Kategori" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Kategori yang ditambahkan tidak ditemukan")
	}
}
