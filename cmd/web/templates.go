package main

import "github.com/alex-orkuma/foodexp/pkg/models"

type templateData struct {
	Product  *models.Products
	Products []*models.Products
}
