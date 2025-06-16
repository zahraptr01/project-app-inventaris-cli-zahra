package utils

import (
	"math"
	"time"
)

const DepreciationRate = 0.20 // 20% per tahun

// Displays the total investment value of all items after calculating depreciation at 20% per year
func CalculateDepreciation(price float64, purchaseDate time.Time) float64 {
	years := time.Since(purchaseDate).Hours() / 24 / 365
	rate := 1 - DepreciationRate
	value := price * math.Pow(rate, years)
	return value
}
