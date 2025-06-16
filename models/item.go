package models

import "time"

type Item struct {
	ID           int
	Name         string
	Price        float64
	PurchaseDate time.Time
	UsageDays    int
	CategoryID   int
}
