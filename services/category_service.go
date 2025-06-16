package services

import (
	"errors"
	"project-app-inventaris-cli-zahra/models"
	"project-app-inventaris-cli-zahra/repository"
	"strings"
)

// using the GetAllCategories function
func ListCategories() ([]models.Category, error) {
	return repository.GetAllCategories()
}

// using the CreateCategory function
func AddCategory(name, description string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("nama kategori tidak boleh kosong")
	}
	c := models.Category{
		Name:        name,
		Description: description,
	}
	return repository.CreateCategory(c)
}

// using the GetCategoryByID function
func GetCategoryDetail(id int) (models.Category, error) {
	return repository.GetCategoryByID(id)
}

// using the UpdateCategory function
func EditCategory(id int, name, description string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("nama kategori tidak boleh kosong")
	}
	c := models.Category{
		Name:        name,
		Description: description,
	}
	return repository.UpdateCategory(id, c)
}

// using the DeleteCategory function
func RemoveCategory(id int) error {
	return repository.DeleteCategory(id)
}
