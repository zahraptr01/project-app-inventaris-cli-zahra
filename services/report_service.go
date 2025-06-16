package services

import (
	"project-app-inventaris-cli-zahra/models"
	"project-app-inventaris-cli-zahra/repository"
	"project-app-inventaris-cli-zahra/utils"
	"time"
)

func GetItemsOver100Days() ([]models.Item, error) {
	allItems, err := repository.GetAllItems()
	if err != nil {
		return nil, err
	}

	var result []models.Item
	for _, item := range allItems {
		days := int(time.Since(item.PurchaseDate).Hours() / 24)
		if days > 100 {
			result = append(result, item)
		}
	}
	return result, nil
}

func GetTotalInvestmentValue() (float64, error) {
	allItems, err := repository.GetAllItems()
	if err != nil {
		return 0, err
	}
	var total float64
	for _, item := range allItems {
		value := utils.CalculateDepreciation(item.Price, item.PurchaseDate)
		total += value
	}
	return total, nil
}

func GetItemInvestmentValueByID(id int) (float64, float64, error) {
	item, err := repository.GetItemByID(id)
	if err != nil {
		return 0, 0, err
	}
	depreciated := utils.CalculateDepreciation(item.Price, item.PurchaseDate)
	return item.Price, depreciated, nil
}
