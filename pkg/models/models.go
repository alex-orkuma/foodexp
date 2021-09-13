package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Products struct {
	ID        int
	FoodID    string
	FoodName  string
	ShelfLife string
}
