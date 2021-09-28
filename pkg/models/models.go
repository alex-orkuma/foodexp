package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Products struct {
	ID        int
	FoodID    string
	FoodName  string
	ShelfLife string
}

type User struct {
	ID             int
	FirstName      string
	LastName       string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Status         bool
	Role           string
}
