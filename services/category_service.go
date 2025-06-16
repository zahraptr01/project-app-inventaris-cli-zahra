package services

import (
	"errors"
	"project-app-inventaris-cli-zahra/models"
	"project-app-inventaris-cli-zahra/repository"
	"strings"
)

func ListCategories() ([]models.Category, error) {
	return repository.GetAllCategories()
}

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

func GetCategoryDetail(id int) (models.Category, error) {
	return repository.GetCategoryByID(id)
}

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

func RemoveCategory(id int) error {
	return repository.DeleteCategory(id)
}
